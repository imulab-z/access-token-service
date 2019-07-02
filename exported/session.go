package exported

type Session struct {
	RequestId     string                 `json:"request_id"`
	ClientId      string                 `json:"client_id"`
	RedirectUri   string                 `json:"redirect_uri"`
	Subject       string                 `json:"subject"`
	GrantedScopes []string               `json:"granted_scopes"`
	AccessClaims  map[string]interface{} `json:"access_claims"`
}
