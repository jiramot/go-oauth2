package usecases

import "github.com/jiramot/go-oauth2/internal/core/domains"

type AuthorizationUseCase interface {
    RequestAuthorizationCode(
        amr string,
        clientId string,
        redirectUrl string,
        scope string, state string,
        codeChallenge string,
        codeChallengeMethod string,
        nonce string,
    ) (*domains.Authorization, error)
}
