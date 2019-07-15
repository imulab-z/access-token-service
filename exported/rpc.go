package exported

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/go-kit/kit/log"
	"github.com/imulab-z/common/errors"
	"net"
	"net/rpc"
	"sync"
	"time"
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
	client := &rpcClient{
		host:    host,
		port:    port,
		logger:  logger,
		errChan: make(chan error),
	}
	if err := client.connect(); err != nil {
		return nil, func() {}, err
	}

	go client.ensureConnection()

	return client, func() {
		close(client.errChan)
		_ = client.c.Close()
	}, nil
}

type rpcClient struct {
	sync.RWMutex
	host    string
	port    int
	errChan chan error
	c      *rpc.Client
	logger log.Logger
}

func (s *rpcClient) connect() error {
	s.Lock()
	defer s.Unlock()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", s.host, s.port), 10*time.Second)
	if err != nil {
		s.logger.Log("error", err)
		return err
	}

	s.c = rpc.NewClient(conn)
	return nil
}

func (s *rpcClient) ensureConnection() {
	for {
		_ = <-s.errChan
		s.logger.Log("access-token-service", "reconnecting")
		_ = s.c.Close()
		err := backoff.Retry(s.connect, backoff.NewExponentialBackOff())
		if err != nil {
			s.logger.Log("error", err)
		} else {
			s.logger.Log("access-token-service", "connected")
		}
		time.Sleep(time.Second)
	}
}

func (s *rpcClient) handledError(err error) error {
	s.logger.Log("error", err)
	select {
	case s.errChan <- err:
	default:
	}
	return errors.TemporarilyUnavailable()
}

func (s *rpcClient) Issue(ctx context.Context, session *Session) (*AccessToken, error) {
	var (
		request  = &RpcIssueRequest{Session: session}
		response = &RpcIssueResponse{}
	)

	s.RLock()
	defer s.RUnlock()

	if err := s.c.Call(RpcServiceName+".Issue", request, response); err != nil {
		return nil, s.handledError(err)
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

	s.RLock()
	defer s.RUnlock()

	if err := s.c.Call(RpcServiceName+".Peek", request, response); err != nil {
		return nil, s.handledError(err)
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

	s.RLock()
	defer s.RUnlock()

	if err := s.c.Call(RpcServiceName+".Revoke", request, response); err != nil {
		return s.handledError(err)
	}

	if response.Err != nil {
		s.logger.Log("error", response.Err)
		return response.Err
	} else {
		return nil
	}
}
