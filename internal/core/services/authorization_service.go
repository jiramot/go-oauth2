package services

import (
    "errors"
    "fmt"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/mocks"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    "github.com/segmentio/ksuid"
)

type authorizationService struct {
    loginChallengePort ports.LoginChallengePort
}

func NewAuthorizationService(loginChallengePort ports.LoginChallengePort) *authorizationService {
    return &authorizationService{
        loginChallengePort: loginChallengePort,
    }
}

func (svc *authorizationService) RequestAuthorizationCode(
    amr string,
    clientId string,
    redirectUrl string,
    scope string,
    state string,
    codeChallenge string,
    codeChallengeMethod string,
    nonce string,
) (*domains.Authorization, error) {
    db := mocks.NewClientDb()
    client, _ := db.FindClientByClientId(clientId)
    if client != nil && clientId == client.ClientId {
        loginChallengeCode := ksuid.New().String()
        loginEndpoint := mocks.LoginEndpointUrl

        if amr == "next" {
            loginEndpoint = "next://login"
        }

        if scope != "" && scope != client.Scope {
            return nil, errors.New("invalid scope")
        }

        if redirectUrl != "" && redirectUrl != client.RedirectUrl {
            return nil, errors.New("invalid redirect url")
        }

        loginChallenge := &domains.LoginChallenge{
            LoginChallengeCode:  loginChallengeCode,
            CodeChallengeMethod: codeChallengeMethod,
            CodeChallenge:       codeChallenge,
            State:               state,
            ClientId:            clientId,
            RedirectUrl:         redirectUrl,
            Scope:               client.Scope,
            Amr:                 amr,
        }
        svc.loginChallengePort.SaveLoginChallenge(loginChallenge)

        return &domains.Authorization{
            LoginChallenge:   loginChallengeCode,
            LoginEndpointUrl: fmt.Sprintf("%s?login_challenge=%s", loginEndpoint, loginChallengeCode),
        }, nil
    }

    return nil, errors.New("No client found")

}
