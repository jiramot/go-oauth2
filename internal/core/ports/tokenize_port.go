package ports

import "github.com/jiramot/go-oauth2/internal/core/domains"

type TokenizePort interface {
    CreateToken(*domains.TokenPayload) (string, error)
    ValidateToken(token string) (*domains.TokenPayload, error)
}
