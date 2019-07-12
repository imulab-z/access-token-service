package svc

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/common/errors"
)

func NewRpcService(service exported.Service, logger log.Logger) exported.RpcService {
	return &rpcService{service: service, logger: logger}
}

type rpcService struct {
	service exported.Service
	logger  log.Logger
}

func (s *rpcService) Issue(request *exported.RpcIssueRequest, response *exported.RpcIssueResponse) error {
	tok, err := s.service.Issue(context.Background(), request.Session)
	if err != nil {
		s.logger.Log("error", err)
		response.Err = toServiceError(err)
	}
	response.AccessToken = tok
	return nil
}

func (s *rpcService) Peek(request *exported.RpcPeekRequest, response *exported.RpcPeekResponse) error {
	session, err := s.service.Peek(context.Background(), request.AccessToken)
	if err != nil {
		s.logger.Log("error", err)
		response.Err = toServiceError(err)
	}
	response.Session = session
	return nil
}

func (s *rpcService) Revoke(request *exported.RpcRevokeRequest, response *exported.RpcRevokeResponse) error {
	err := s.service.Revoke(context.Background(), request.AccessToken)
	if err != nil {
		s.logger.Log("error", err)
		response.Err = toServiceError(err)
	}
	return nil
}

func toServiceError(err error) *errors.ServiceError {
	e := errors.EnsureConnectError(err)
	return &errors.ServiceError{
		Status:    e.StatusCode(),
		ErrorCode: e.Error(),
		Reason:    e.Description(),
		Meta:      e.Metadata(),
	}
}
