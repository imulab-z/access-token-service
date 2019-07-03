package http

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-redis/redis"
	"net/http"
)

const (
	OK = "ok"
	NotOK = "ko"
)

func makeHealthEndpoint(redisClient *redis.Client, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var resp *healthResponse
		{
			resp = &healthResponse{}
			if err := redisClient.Ping().Err(); err != nil {
				logger.Log("error", err)
				resp.RedisDatabase = NotOK
			} else {
				resp.RedisDatabase = OK
			}
		}
		return resp, nil
	}
}

func decodeHealthRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeHealthResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if response.(*healthResponse).HasNotOkItem() {
		w.WriteHeader(503)
	} else {
		w.WriteHeader(200)
	}
	return json.NewEncoder(w).Encode(response)
}

type healthResponse struct {
	RedisDatabase	string	`json:"redis_database"`
}

func (r *healthResponse) HasNotOkItem() bool {
	return r.RedisDatabase != OK
}


