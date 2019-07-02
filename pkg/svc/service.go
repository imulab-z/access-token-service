package svc

import (
	"github.com/go-redis/redis"
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
