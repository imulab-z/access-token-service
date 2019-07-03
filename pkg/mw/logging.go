package mw

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/imulab-z/access-token-service/exported"
	"time"
)

func NewLoggingMiddleware(next exported.Service, logger log.Logger) exported.Service {
	return &loggingMiddleware{logger:logger, next:next}
}

type loggingMiddleware struct {
	logger 	log.Logger
	next 	exported.Service
}

func (mw loggingMiddleware) Issue(ctx context.Context, session *exported.Session) (*exported.AccessToken, error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Issue",
			"took", time.Since(begin))
	}(time.Now())
	return mw.next.Issue(ctx, session)
}

func (mw loggingMiddleware) Peek(ctx context.Context, accessToken string) (*exported.Session, error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Peek",
			"took", time.Since(begin))
	}(time.Now())
	return mw.next.Peek(ctx, accessToken)
}

func (mw loggingMiddleware) Revoke(ctx context.Context, accessToken string) error {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Revoke",
			"took", time.Since(begin))
	}(time.Now())
	return mw.next.Revoke(ctx, accessToken)
}

