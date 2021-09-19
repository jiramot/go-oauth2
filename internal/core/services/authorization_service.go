package services

import (
    "errors"
    "fmt"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    "github.com/segmentio/ksuid"
)

type authorizationService struct {
    loginChallengePort ports.LoginChallengePort
    clients            domains.Clients
    loginEndpoint      string
}

func NewAuthorizationService(
    loginChallengePort ports.LoginChallengePort,
    clients domains.Clients,
    loginEndpoint string,
) *authorizationService {
    return &authorizationService{
        loginChallengePort: loginChallengePort,
        clients:            clients,
        loginEndpoint:      loginEndpoint,
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
    client, _ := svc.clients.FindClientByClientId(clientId)
    if client != nil && clientId == client.ClientId {
        loginChallengeCode := ksuid.New().String()

        if amr == "next" {
            svc.loginEndpoint = "next://login"
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
            RedirectUrl:         client.RedirectUrl,
            Scope:               client.Scope,
            Amr:                 amr,
        }
        err := svc.loginChallengePort.SaveLoginChallenge(loginChallenge)
        if err != nil {
            return nil, err
        }
        return &domains.Authorization{
            LoginChallenge:   loginChallengeCode,
            LoginEndpointUrl: fmt.Sprintf("%s?login_challenge=%s", svc.loginEndpoint, loginChallengeCode),
        }, nil
    }

    return nil, errors.New("No client found")

}
