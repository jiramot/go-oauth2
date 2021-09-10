package domains

import "time"

type TokenPayload struct {
    Issue     string    `json:"iss"`
    Subject   string    `json:"sub"`
    Scope     string    `json:"scope"`
    Amr       [1]string `json:"amr"`
    IssuedAt  int64     `json:"iat"`
    ExpiredAt int64     `json:"exp"`
    ClientId  string    `json:"aud"`
}

const TokenTtl = time.Minute * 15

func NewTokenPayload(clientId string, subject string, scope string, amr string, duration time.Duration) *TokenPayload {
    now := time.Now()
    return &TokenPayload{
        Issue:     "http://github.com/jiramot/go-oauth2",
        ClientId:  clientId,
        Subject:   subject,
        Scope:     scope,
        Amr:       [1]string{amr},
        IssuedAt:  now.Unix(),
        ExpiredAt: now.Add(duration).Unix(),
    }
}
func (c TokenPayload) Valid() error {
    return nil
}
