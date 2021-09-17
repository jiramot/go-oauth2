package domains

type AccessToken struct {
    AccessToken string `json:"access_token"`
    ExpireAt    int64  `json:"expires_at"`
    TokenType   string `json:"token_type"`
}
