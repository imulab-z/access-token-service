package http

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/pkg"
	"net/http"
)

func decodePeekRequest(_ context.Context, r *http.Request) (interface{}, error) {
	refreshToken := r.URL.Query().Get("access_token")
	if len(refreshToken) == 0 {
		return nil, pkg.ErrInvalidRequest("missing required parameter 'access_token'.")
	}
	return refreshToken, nil
}

func encodePeekResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func makePeekEndpoint(svc exported.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Peek(ctx, request.(string))
	}
}
