package exported

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/imulab-z/common/errors"
	"net/rpc"
)

const RpcServiceName = "AccessTokenService"

type RpcService interface {
	Issue(request *RpcIssueRequest, response *RpcIssueResponse) error
	Peek(request *RpcPeekRequest, response *RpcPeekResponse) error
	Revoke(request *RpcRevokeRequest, response *RpcRevokeResponse) error
}

type RpcIssueRequest struct {
	Session *Session
}

type RpcIssueResponse struct {
	AccessToken *AccessToken
	Err         *errors.ServiceError
}

type RpcPeekRequest struct {
	AccessToken string
}

type RpcPeekResponse struct {
	Session *Session
	Err     *errors.ServiceError
}

type RpcRevokeRequest struct {
	AccessToken string
}

type RpcRevokeResponse struct {
	Err *errors.ServiceError
}

func NewRpcClient(host string, port int, logger log.Logger) Service {
	return &rpcClient{
		host:   host,
		port:   port,
		logger: logger,
	}
}

type rpcClient struct {
	host   string
	port   int
	logger log.Logger
}

func (s *rpcClient) connected(ctx context.Context, methodName string, args interface{}, reply interface{}) error {
	var (
		errChan  chan error
		doneChan chan struct{}
	)
	{
		errChan = make(chan error, 1)
		doneChan = make(chan struct{}, 1)

		defer close(errChan)
		defer close(doneChan)
	}

	go func() {
		client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
		if err != nil {
			errChan <- err
			return
		}

		call := <-client.Go(
			fmt.Sprintf("%s.%s", RpcServiceName, methodName),
			args,
			reply,
			make(chan *rpc.Call, 1),
		).Done

		if err := call.Error; err != nil {
			errChan <- err
			return
		}

		doneChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	case <-doneChan:
		return nil
	}
}

func (s *rpcClient) Issue(ctx context.Context, session *Session) (*AccessToken, error) {
	var (
		request  = &RpcIssueRequest{Session: session}
		response = &RpcIssueResponse{}
	)

	if err := s.connected(ctx, "Issue", request, response); err != nil {
		s.logger.Log("error", err)
		return nil, errors.TemporarilyUnavailable()
	}

	if response.Err != nil {
		s.logger.Log("error", response.Err)
		return nil, response.Err
	}

	return response.AccessToken, nil
}

func (s *rpcClient) Peek(ctx context.Context, accessToken string) (*Session, error) {
	var (
		request  = &RpcPeekRequest{AccessToken: accessToken}
		response = &RpcPeekResponse{}
	)

	if err := s.connected(ctx, "Peek", request, response); err != nil {
		s.logger.Log("error", err)
		return nil, errors.TemporarilyUnavailable()
	}

	if response.Err != nil {
		s.logger.Log("error", response.Err)
		return nil, response.Err
	}

	return response.Session, nil
}

func (s *rpcClient) Revoke(ctx context.Context, accessToken string) error {
	var (
		request  = &RpcRevokeRequest{AccessToken: accessToken}
		response = &RpcRevokeResponse{}
	)

	if err := s.connected(ctx, "Revoke", request, response); err != nil {
		s.logger.Log("error", err)
		return errors.TemporarilyUnavailable()
	}

	if response.Err != nil {
		s.logger.Log("error", response.Err)
		return response.Err
	}

	return nil
}
