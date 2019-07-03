package exported

import "context"

type Service interface {
	// Issue a new access token
	Issue(ctx context.Context, session *Session) (*AccessToken, error)
	// Peek inside an access token. The returned session may not contain ALL the information
	// used to create it.
	Peek(ctx context.Context, accessToken string) (*Session, error)
	// Revoke an access token
	Revoke(ctx context.Context, accessToken string) error
}

type AccessToken struct {
	Token		string	`json:"token"`
	TokenType	string	`json:"token_type"`
	ExpiresIn	int64	`json:"expires_in"`
}