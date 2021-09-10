package domains

import "time"

type TokenPayload struct {
    Subject   string    `json:"sub"`
    Scope     string    `json:"scope"`
    Amr       [1]string `json:"amr"`
    IssuedAt  time.Time `json:"iat"`
    ExpiredAt time.Time `json:"exp"`
    ClientId  string    `json:"aud"`
}

const TokenTtl = time.Minute * 15

func NewTokenPayload(clientId string, subject string, scope string, amr string, duration time.Duration) *TokenPayload {
    now := time.Now()
    return &TokenPayload{
        ClientId:  clientId,
        Subject:   subject,
        Scope:     scope,
        Amr:       [1]string{amr},
        IssuedAt:  now,
        ExpiredAt: now.Add(duration),
    }
}
func (c TokenPayload) Valid() error {
    return nil
}
