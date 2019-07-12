package svc

import (
	"context"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/imulab-z/access-token-service/pkg"
	"github.com/imulab-z/common/errors"
)

func (s *service) Peek(ctx context.Context, accessToken string) (*exported.Session, error) {
	if s.isRevoked(ctx, accessToken) {
		return nil, pkg.ErrInvalidToken()
	}

	if session, err := s.strategy.DeTokenize(ctx, accessToken); err != nil {
		return nil, errors.EnsureConnectError(err)
	} else {
		return session, nil
	}
}
