package pkg

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/imulab-z/access-token-service/exported"
	"github.com/satori/go.uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"strings"
	"time"
)

const (
	TokenStrategyJwt = "jwt"
)

type TokenStrategy interface {
	// Return the type of token this strategy generates
	TokenType() string

	// Create a new access token
	New(ctx context.Context, session *exported.Session) (string, error)

	// Validate and decode the given access token and reveal its data
	DeTokenize(ctx context.Context, accessToken string) (*exported.Session, error)
}

func NewJwtTokenStrategy(
	keySet *jose.JSONWebKeySet,
	alg jose.SignatureAlgorithm,
	issuer string,
	ttl int64,
	leeway int64,
	logger log.Logger,
) TokenStrategy {
	publicKeySet := &jose.JSONWebKeySet{Keys: make([]jose.JSONWebKey, 0)}
	for _, k := range keySet.Keys {
		publicKeySet.Keys = append(publicKeySet.Keys, k.Public())
	}
	return &jwtTokenStrategy{
		fullKeySet:       keySet,
		publicKeySet:     publicKeySet,
		alg:              alg,
		issuer:           issuer,
		defaultLife:      time.Duration(ttl) * time.Second,
		validationLeeway: time.Duration(leeway) * time.Second,
		logger:           logger,
	}
}

type jwtTokenStrategy struct {
	fullKeySet       *jose.JSONWebKeySet
	publicKeySet     *jose.JSONWebKeySet
	alg              jose.SignatureAlgorithm
	issuer           string
	defaultLife      time.Duration
	validationLeeway time.Duration
	logger           log.Logger
}

func (s *jwtTokenStrategy) TokenType() string {
	return "Bearer"
}

func (s *jwtTokenStrategy) New(ctx context.Context, session *exported.Session) (string, error) {
	jwk := selectSignatureKey(s.fullKeySet, s.alg, session.ClientId)
	if jwk == nil {
		return "", ErrServer("failed to select signing key for access token.")
	}

	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: s.alg, Key: jwk}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		s.logger.Log("error", err)
		return "", ErrServer("failed to create signer for access token")
	}

	token, err := jwt.Signed(signer).Claims(s.createClaims(session)).CompactSerialize()
	if err != nil {
		s.logger.Log("error", err)
		return "", ErrServer("failed to generate access token")
	}

	return token, nil
}

func (s *jwtTokenStrategy) DeTokenize(ctx context.Context, accessToken string) (*exported.Session, error) {
	var err error

	var tok *jwt.JSONWebToken
	{
		tok, err = jwt.ParseSigned(accessToken)
		if err != nil {
			s.logger.Log("error", err)
			return nil, ErrInvalidToken(accessToken)
		}
	}

	var claims *claims
	{
		claims = newClaims()
		err = tok.Claims(s.publicKeySet, claims)
		if err != nil {
			s.logger.Log("error", err)
			return nil, ErrInvalidToken(accessToken)
		}
	}

	err = claims.Claims.ValidateWithLeeway(jwt.Expected{
		Issuer: s.issuer,
		Time:   time.Now(),
	}, s.validationLeeway)
	if err != nil {
		s.logger.Log("error", err)
		return nil, ErrInvalidToken(accessToken)
	}

	var session *exported.Session
	{
		session = &exported.Session{}
		session.ClientId = claims.Audience[0]
		session.Subject = claims.Subject
		session.GrantedScopes = strings.Split(claims.Scope, " ")
		session.RequestId = claims.RequestId
		session.AccessClaims = claims.Extra
	}

	return session, nil
}

func (s *jwtTokenStrategy) createClaims(session *exported.Session) *claims {
	return &claims{
		Claims: &jwt.Claims{
			ID:        uuid.NewV4().String(),
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   session.Subject,
			Expiry:    jwt.NewNumericDate(time.Now().Add(s.defaultLife)),
			Audience:  jwt.Audience([]string{session.ClientId}),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		Scope:     strings.Join(session.GrantedScopes, " "),
		RequestId: session.RequestId,
		Extra:     session.AccessClaims,
	}
}

type claims struct {
	*jwt.Claims
	Scope     string                 `json:"scope,omitempty"`
	RequestId string                 `json:"request_id,omitempty"`
	Extra     map[string]interface{} `json:"extra,omitempty"`
}

func newClaims() *claims {
	return &claims{
		Claims: &jwt.Claims{},
		Extra:  make(map[string]interface{}),
	}
}
