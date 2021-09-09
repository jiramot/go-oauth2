package ports

import "github.com/jiramot/go-oauth2/internal/core/domains"

type AuthorizationPort interface {
    Authorization(authorizationType string, amr string, clientId string, clientSecret string, scope string) (domains.Authorization, error)
}
