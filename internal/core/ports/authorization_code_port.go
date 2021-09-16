package ports

import "github.com/jiramot/go-oauth2/internal/core/domains"

type AuthorizationCodePort interface {
    SaveAuthorizationCode(code *domains.AuthorizationCode) error
    GetAuthorizationCodeFromCode(code string) (*domains.AuthorizationCode, error)
    RemoveAuthorizationCodeFromCode(code string) error
}
