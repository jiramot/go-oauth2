package usecases

import "github.com/jiramot/go-oauth2/internal/core/domains"

type TokenUseCase interface {
    CreateTokenForAuthorizationCodeGrantType(token domains.Token) (*domains.AccessToken, error)
    CreateTokenForImplicitGrantType(request ImplicitGrantRequest) (*domains.AccessToken, error)
    IntrospectToken(token string) (*domains.TokenPayload, error)
}

type ImplicitGrantRequest struct {
    Cif      string
    ClientId string
    Amr      string
}
