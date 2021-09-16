package ports

import "github.com/jiramot/go-oauth2/internal/core/domains"

type AdminAcceptLoginUseCase interface {
    AcceptLogin(loginChallengeCode string, cif string) (*domains.AuthorizationCode, error)
    RejectLogin(loginChallengeCode string, cif string) error
}
