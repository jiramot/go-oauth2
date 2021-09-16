package usecases

import "github.com/jiramot/go-oauth2/internal/core/domains"

type TokenUseCase interface {
    CreateTokenForAuthorizationCodeGrantType(token domains.Token) (string, error)
    IntrospectToken(token string) (*domains.TokenPayload, error)
}
