package exported

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/imulab-z/access-token-service/pb"
	"google.golang.org/grpc"
	"strings"
)

func NewGrpcClient(conn *grpc.ClientConn, logger log.Logger, errFunc func(code, description string) error) Service {
	var options []gt.ClientOption
	{
		options = make([]gt.ClientOption, 0)
	}

	var c *grpcClient
	{
		c = &grpcClient{}
		c.logger = logger
		c.errFunc = errFunc

		c.issueEndpoint = gt.NewClient(
			conn,
			"pb.AccessTokenService",
			"Issue",
			c.encodeIssueRequest,
			c.decodeIssueResponse,
			&pb.IssueResponse{},
			options...
		).Endpoint()

		c.peekEndpoint = gt.NewClient(
			conn,
			"pb.AccessTokenService",
			"Peek",
			c.encodePeekRequest,
			c.decodePeekResponse,
			&pb.PeekResponse{},
			options...
		).Endpoint()

		c.revokeEndpoint = gt.NewClient(
			conn,
			"pb.AccessTokenService",
			"Revoke",
			c.encodeRevokeRequest,
			c.decodeRevokeResponse,
			&pb.RevokeResponse{},
			options...
		).Endpoint()
	}

	return c
}

type grpcClient struct {
	errFunc        func(code, description string) error
	issueEndpoint  endpoint.Endpoint
	peekEndpoint   endpoint.Endpoint
	revokeEndpoint endpoint.Endpoint
	logger         log.Logger
}

func (c *grpcClient) Issue(ctx context.Context, session *Session) (string, error) {
	resp, err := c.issueEndpoint(ctx, session)
	if err != nil {
		c.logger.Log("error", err)
		return "", err
	}
	return resp.(string), nil
}

func (c *grpcClient) Peek(ctx context.Context, refreshToken string) (*Session, error) {
	resp, err := c.peekEndpoint(ctx, refreshToken)
	if err != nil {
		c.logger.Log("error", err)
		return nil, err
	}
	return resp.(*Session), nil
}

func (c *grpcClient) Revoke(ctx context.Context, refreshToken string) error {
	_, err := c.revokeEndpoint(ctx, refreshToken)
	if err != nil {
		c.logger.Log("error", err)
		return err
	}
	return nil
}

func (c *grpcClient) encodeIssueRequest(_ context.Context, request interface{}) (interface{}, error) {
	r := request.(*Session)

	var accessClaimsJson string
	{
		accessClaimsJson = "{}"
		if len(r.AccessClaims) > 0 {
			raw, err := json.Marshal(r.AccessClaims)
			if err != nil {
				return nil, err
			}
			accessClaimsJson = string(raw)
		}
	}

	return &pb.IssueRequest{
		Session: &pb.Session{
			RequestId:        r.RequestId,
			ClientId:         r.ClientId,
			RedirectUri:      r.RedirectUri,
			Subject:          r.Subject,
			GrantedScopes:    r.GrantedScopes,
			AccessClaimsJson: accessClaimsJson,
		},
	}, nil
}

func (c *grpcClient) decodeIssueResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	r := grpcReply.(*pb.IssueResponse)
	if r.Success {
		return r.AccessToken, nil
	} else {
		return "", c.errFunc(r.ErrorCode, r.ErrorDescription)
	}
}

func (c *grpcClient) encodePeekRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &pb.PeekRequest{AccessToken: request.(string)}, nil
}

func (c *grpcClient) decodePeekResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	r := grpcReply.(*pb.PeekResponse)
	if r.Success {
		var accessClaims map[string]interface{}
		{
			accessClaims = make(map[string]interface{})
			if len(r.Session.AccessClaimsJson) > 0 {
				if err := json.NewDecoder(strings.NewReader(r.Session.AccessClaimsJson)).Decode(&accessClaims); err != nil {
					return nil, err
				}
			}
		}

		return &Session{
			RequestId:     r.Session.RequestId,
			ClientId:      r.Session.ClientId,
			RedirectUri:   r.Session.RedirectUri,
			Subject:       r.Session.Subject,
			GrantedScopes: r.Session.GrantedScopes,
			AccessClaims:  accessClaims,
		}, nil
	} else {
		return nil, c.errFunc(r.ErrorCode, r.ErrorDescription)
	}
}

func (c *grpcClient) encodeRevokeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &pb.RevokeRequest{AccessToken: request.(string)}, nil
}

func (c *grpcClient) decodeRevokeResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	r := grpcReply.(*pb.RevokeResponse)
	if r.Success {
		return struct{}{}, nil
	} else {
		return struct{}{}, c.errFunc(r.ErrorCode, r.ErrorDescription)
	}
}
