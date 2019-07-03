package svc

import (
	"context"
	"github.com/imulab-z/access-token-service/exported"
	"time"
)

func (s *service) Issue(ctx context.Context, session *exported.Session) (*exported.AccessToken, error) {
	var (
		token string
		err   error
	)

	token, err = s.strategy.New(ctx, session)
	if err != nil {
		return nil, err
	}

	return &exported.AccessToken{
		Token: token,
		TokenType: s.strategy.TokenType(),
		ExpiresIn: int64(s.defaultLife / time.Second),
	}, nil
}
