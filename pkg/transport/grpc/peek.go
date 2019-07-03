package grpc

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/pb"
	"github.com/imulab-z/access-token-service/pkg"
)

func decodePeekRequest(_ context.Context, r interface{}) (interface{}, error) {
	return r.(*pb.PeekRequest).AccessToken, nil
}

func encodePeekResponse(_ context.Context, response interface{}) (interface{}, error) {
	switch response.(type) {
	case *pb.PeekResponse:
		return response, nil
	case *exported.Session:
		var accessClaimsJson string
		{
			accessClaimsJson = ""
			raw, err := json.Marshal(response.(*exported.Session).AccessClaims)
			if err != nil {
				return nil, pkg.ErrServer("failed to encode response")
			}
			accessClaimsJson = string(raw)
		}

		return &pb.PeekResponse{
			Success: true,
			Session: &pb.Session{
				RequestId:        response.(*exported.Session).RequestId,
				ClientId:         response.(*exported.Session).ClientId,
				RedirectUri:      response.(*exported.Session).RedirectUri,
				Subject:          response.(*exported.Session).Subject,
				GrantedScopes:    response.(*exported.Session).GrantedScopes,
				AccessClaimsJson: accessClaimsJson,
			},
		}, nil
	default:
		return nil, pkg.ErrServer("unknown response type")
	}
}

func makePeekEndpoint(svc exported.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		session, err := svc.Peek(ctx, request.(string))
		if err != nil {
			var se *pkg.ServiceError
			switch err.(type) {
			case *pkg.ServiceError:
				se = err.(*pkg.ServiceError)
			default:
				se = pkg.ErrServer(err.Error()).(*pkg.ServiceError)
			}

			return &pb.PeekResponse{
				Success:          false,
				ErrorCode:        se.Err,
				ErrorDescription: se.Description,
			}, nil
		}
		return session, nil
	}
}
