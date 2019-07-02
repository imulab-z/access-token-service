package svc

import (
	"context"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/pkg"
)

func (s *service) Peek(ctx context.Context, accessToken string) (*exported.Session, error) {
	if s.isRevoked(ctx, accessToken) {
		return nil, pkg.ErrInvalidToken(accessToken)
	}
	return s.strategy.DeTokenize(ctx, accessToken)
}
