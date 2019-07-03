package svc

import (
	"github.com/go-redis/redis"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/pkg"
	"time"
)
import "github.com/go-kit/kit/log"

type service struct {
	redis            *redis.Client
	strategy         pkg.TokenStrategy
	defaultLife      time.Duration
	validationLeeway time.Duration
	logger           log.Logger
}

func NewService(redis *redis.Client, strategy pkg.TokenStrategy, ttl int64, leeway int64, logger log.Logger) exported.Service {
	return &service{
		redis:            redis,
		strategy:         strategy,
		defaultLife:      time.Duration(ttl) * time.Second,
		validationLeeway: time.Duration(leeway) * time.Second,
		logger:           logger,
	}
}
