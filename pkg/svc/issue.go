package svc

import (
	"context"
	"github.com/imulab-z/access-token-service/exported"
)

func (s *service) Issue(ctx context.Context, session *exported.Session) (string, error) {
	return s.strategy.New(ctx, session)
}
