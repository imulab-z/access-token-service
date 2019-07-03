package grpc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/pb"
	"github.com/imulab-z/access-token-service/pkg"
)

func decodeRevokeRequest(_ context.Context, r interface{}) (interface{}, error) {
	return r.(*pb.RevokeRequest).AccessToken, nil
}

func encodeRevokeResponse(_ context.Context, response interface{}) (interface{}, error) {
	switch response.(type) {
	case *pb.RevokeResponse:
		return response, nil
	default:
		return &pb.RevokeResponse{Success: true}, nil
	}
}

func makeRevokeEndpoint(svc exported.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		err = svc.Revoke(ctx, request.(string))
		if err != nil {
			var se *pkg.ServiceError
			switch err.(type) {
			case *pkg.ServiceError:
				se = err.(*pkg.ServiceError)
			default:
				se = pkg.ErrServer(err.Error()).(*pkg.ServiceError)
			}

			return &pb.RevokeResponse{
				Success:          false,
				ErrorCode:        se.Err,
				ErrorDescription: se.Description,
			}, nil
		}
		return struct{}{}, nil
	}
}
