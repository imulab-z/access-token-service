package http

import (
	"github.com/go-kit/kit/log"
	ht "github.com/go-kit/kit/transport/http"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/common/errors"
	"net/http"
)

func NewHTTPServer(redis *redis.Client, svc exported.Service, logger log.Logger) *mux.Router {

	var (
		issueHandler  *ht.Server
		peekHandler   *ht.Server
		revokeHandler *ht.Server
		healthHandler *ht.Server
	)
	{
		options := []ht.ServerOption{
			ht.ServerErrorEncoder(errors.EncodeErrorToBody),
		}

		issueHandler = ht.NewServer(
			makeIssueEndpoint(svc),
			decodeIssueRequest,
			encodeIssueResponse,
			options...
		)

		peekHandler = ht.NewServer(
			makePeekEndpoint(svc),
			decodePeekRequest,
			encodePeekResponse,
			options...
		)

		revokeHandler = ht.NewServer(
			makeRevokeEndpoint(svc),
			decodeRevokeRequest,
			encodeRevokeResponse,
			options...
		)

		healthHandler = ht.NewServer(
			makeHealthEndpoint(redis, logger),
			decodeHealthRequest,
			encodeHealthResponse,
			options...
		)
	}

	var r *mux.Router
	{
		r = mux.NewRouter()
		r.Methods(http.MethodPost).Path("/token").Handler(issueHandler)
		r.Methods(http.MethodGet).Path("/session").Handler(peekHandler)
		r.Methods(http.MethodDelete).Path("/token").Handler(revokeHandler)
		r.Methods(http.MethodGet).Path("/health").Handler(healthHandler)
	}

	return r
}
