package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/pkg"
	"net/http"
)

func decodeRevokeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	refreshToken := r.URL.Query().Get("access_token")
	if len(refreshToken) == 0 {
		return nil, pkg.ErrParameterRequired("access_token")
	}
	return refreshToken, nil
}

func encodeRevokeResponse(_ context.Context, w http.ResponseWriter, _ interface{}) error {
	w.WriteHeader(204)
	return nil
}

func makeRevokeEndpoint(svc exported.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return struct{}{}, svc.Revoke(ctx, request.(string))
	}
}
