package ports

import "github.com/jiramot/go-oauth2/internal/core/domains"

type AuthorizationUseCase interface {
    AuthorizationCode(amr string, clientId string, redirectUrl string, scope string) (domains.Authorization, error)
}
