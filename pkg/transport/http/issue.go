package http

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/pkg"
	"net/http"
)

type tokenPayload struct {
	AccessToken string `json:"access_token"`
}

func decodeIssueRequest(_ context.Context, r *http.Request) (interface{}, error) {
	payload := &exported.Session{}
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		return nil, pkg.ErrInvalidRequest("unable to read request body")
	}
	return payload, nil
}

func encodeIssueResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(&tokenPayload{
		AccessToken: response.(string),
	})
}

func makeIssueEndpoint(svc exported.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Issue(ctx, request.(*exported.Session))
	}
}
