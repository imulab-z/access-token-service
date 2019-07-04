package grpc

import (
	"context"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/atpb"
)

func NewGrpcServer(service exported.Service, logger log.Logger) atpb.AccessTokenServiceServer {
	var server *grpcServer
	{
		options := []gt.ServerOption{
			gt.ServerErrorLogger(logger),
		}
		server = &grpcServer{
			issue: gt.NewServer(
				makeIssueEndpoint(service),
				decodeIssueRequest,
				encodeIssueResponse,
				options...
			),
			peek: gt.NewServer(
				makePeekEndpoint(service),
				decodePeekRequest,
				encodePeekResponse,
				options...
			),
			revoke: gt.NewServer(
				makeRevokeEndpoint(service),
				decodeRevokeRequest,
				encodeRevokeResponse,
				options...
			),
		}
	}

	return server
}

type grpcServer struct {
	issue  gt.Handler
	peek   gt.Handler
	revoke gt.Handler
}

func (s *grpcServer) Issue(ctx context.Context, req *atpb.IssueRequest) (*atpb.IssueResponse, error) {
	_, resp, err := s.issue.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*atpb.IssueResponse), nil
}

func (s *grpcServer) Peek(ctx context.Context, req *atpb.PeekRequest) (*atpb.PeekResponse, error) {
	_, resp, err := s.peek.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*atpb.PeekResponse), nil
}

func (s *grpcServer) Revoke(ctx context.Context, req *atpb.RevokeRequest) (*atpb.RevokeResponse, error) {
	_, resp, err := s.revoke.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*atpb.RevokeResponse), nil
}
