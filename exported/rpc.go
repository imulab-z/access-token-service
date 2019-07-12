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

func NewRpcClient(host string, port int, logger log.Logger) (service Service, closer func(), err error) {
	conn, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logger.Log("error", err)
		return nil, func() {}, err
	}
	return &rpcClient{c: conn, logger: logger}, func() {
		_ = conn.Close()
	}, nil
}

type rpcClient struct {
	c      *rpc.Client
	logger log.Logger
}

func (s *rpcClient) Issue(ctx context.Context, session *Session) (*AccessToken, error) {
	var (
		request  = &RpcIssueRequest{Session: session}
		response = &RpcIssueResponse{}
	)
	if err := s.c.Call(RpcServiceName+".Issue", request, response); err != nil {
		s.logger.Log("error", err)
		return nil, errors.TemporarilyUnavailable()
	}

	if response.Err != nil {
		s.logger.Log("error", response.Err)
		return nil, response.Err
	} else {
		return response.AccessToken, nil
	}
}

func (s *rpcClient) Peek(ctx context.Context, accessToken string) (*Session, error) {
	var (
		request  = &RpcPeekRequest{AccessToken: accessToken}
		response = &RpcPeekResponse{}
	)
	if err := s.c.Call(RpcServiceName+".Peek", request, response); err != nil {
		s.logger.Log("error", err)
		return nil, errors.TemporarilyUnavailable()
	}

	if response.Err != nil {
		s.logger.Log("error", response.Err)
		return nil, response.Err
	} else {
		return response.Session, nil
	}
}

func (s *rpcClient) Revoke(ctx context.Context, accessToken string) error {
	var (
		request  = &RpcRevokeRequest{AccessToken: accessToken}
		response = &RpcRevokeResponse{}
	)
	if err := s.c.Call(RpcServiceName+".Revoke", request, response); err != nil {
		s.logger.Log("error", err)
		return errors.TemporarilyUnavailable()
	}

	if response.Err != nil {
		s.logger.Log("error", response.Err)
		return response.Err
	} else {
		return nil
	}
}
