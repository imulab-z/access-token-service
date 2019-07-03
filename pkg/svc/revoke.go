package svc

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/imulab-z/access-token-service/pkg"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

func (s *service) Revoke(ctx context.Context, accessToken string) error {
	if s.isRevoked(ctx, accessToken) {
		return nil
	}

	tok, err := jose.ParseSigned(accessToken)
	if err != nil {
		s.logger.Log("error", err)
		return pkg.ErrInvalidToken(accessToken)
	}

	claims := jwt.Claims{}
	if err := json.NewDecoder(bytes.NewReader(tok.UnsafePayloadWithoutVerification())).Decode(&claims); err != nil {
		s.logger.Log("error", err)
		return pkg.ErrInvalidToken(accessToken)
	}

	ttl := claims.Expiry.Time().Sub(time.Now())
	if ttl < 0 {
		return pkg.ErrTokenExpired(accessToken)
	}

	cmd := s.redis.Set(tokenKey(accessToken), "0", ttl + s.validationLeeway)
	if err := cmd.Err(); err != nil {
		s.logger.Log("error", err)
		return pkg.ErrServer("failed to revoke token")
	}

	return nil
}

func (s *service) isRevoked(ctx context.Context, accessToken string) bool {
	return s.redis.Get(tokenKey(accessToken)).Err() == nil
}

func tokenKey(accessToken string) string {
	return fmt.Sprintf("access:%s", hash(accessToken))
}

func hash(accessToken string) string {
	h := sha256.New()
	return string(h.Sum([]byte(accessToken)))
}