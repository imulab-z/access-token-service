package grpc

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/atpb"
	"github.com/imulab-z/access-token-service/pkg"
	"strings"
)

func decodeIssueRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*atpb.IssueRequest)

	var accessClaims map[string]interface{}
	{
		accessClaims = make(map[string]interface{})
		if len(req.Session.AccessClaimsJson) > 0 {
			if err := json.NewDecoder(strings.NewReader(req.Session.AccessClaimsJson)).Decode(&accessClaims); err != nil {
				return nil, err
			}
		}
	}

	return &exported.Session{
		RequestId:     req.Session.RequestId,
		ClientId:      req.Session.ClientId,
		RedirectUri:   req.Session.RedirectUri,
		Subject:       req.Session.Subject,
		GrantedScopes: req.Session.GrantedScopes,
		AccessClaims:  accessClaims,
	}, nil
}

func encodeIssueResponse(_ context.Context, response interface{}) (interface{}, error) {
	switch response.(type) {
	case *atpb.IssueResponse:
		return response.(*atpb.IssueResponse), nil
	case *exported.AccessToken:
		return &atpb.IssueResponse{
			Success:          true,
			ErrorCode:        "",
			ErrorDescription: "",
			AccessToken:      response.(*exported.AccessToken).Token,
			TokenType:        response.(*exported.AccessToken).TokenType,
			ExpiresIn:        response.(*exported.AccessToken).ExpiresIn,
		}, nil
	default:
		return nil, pkg.ErrServer("unknown response type")
	}
}

func makeIssueEndpoint(svc exported.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		token, err := svc.Issue(ctx, request.(*exported.Session))
		if err != nil {
			var se *pkg.ServiceError
			switch err.(type) {
			case *pkg.ServiceError:
				se = err.(*pkg.ServiceError)
			default:
				se = pkg.ErrServer(err.Error()).(*pkg.ServiceError)
			}

			return &atpb.IssueResponse{
				Success:          false,
				ErrorCode:        se.Err,
				ErrorDescription: se.Description,
			}, nil
		}
		return token, nil
	}
}
