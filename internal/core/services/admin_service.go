package services

import (
    "errors"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    "github.com/segmentio/ksuid"
)

type adminService struct {
    loginChallengePort    ports.LoginChallengePort
    authorizationCodePort ports.AuthorizationCodePort
}

func NewAdminService(
    loginChallengePort ports.LoginChallengePort,
    authorizationCodePort ports.AuthorizationCodePort,
) *adminService {
    return &adminService{
        loginChallengePort:    loginChallengePort,
        authorizationCodePort: authorizationCodePort,
    }
}

func (svc *adminService) AcceptLogin(loginChallengeCode string, cif string) (*domains.AuthorizationCode, error) {
    loginChallenge, _ := svc.loginChallengePort.GetLoginChallenge(loginChallengeCode)
    svc.loginChallengePort.RemoveLoginChallenge(loginChallengeCode)
    if loginChallenge != nil {
        authorizationCode := createAuthorizationCodeFromLoginChallenge(loginChallengeCode, loginChallenge, cif)
        err := svc.authorizationCodePort.SaveAuthorizationCode(authorizationCode)
        if err != nil {
            return nil, err
        } else {
            return authorizationCode, nil
        }
    } else {
        return nil, errors.New("invalid request")
    }
}

func (svc *adminService) RejectLogin(loginChallengeCode string, cif string) error {
    return nil
}

func createAuthorizationCodeFromLoginChallenge(
    loginChallengeCode string,
    loginChallenge *domains.LoginChallenge,
    cif string,
) *domains.AuthorizationCode {
    return &domains.AuthorizationCode{
        Code:                ksuid.New().String(),
        LoginChallengeCode:  loginChallengeCode,
        CodeChallengeMethod: loginChallenge.CodeChallengeMethod,
        CodeChallenge:       loginChallenge.CodeChallenge,
        State:               loginChallenge.State,
        ClientId:            loginChallenge.ClientId,
        RedirectUrl:         loginChallenge.RedirectUrl,
        Scope:               loginChallenge.Scope,
        Amr:                 loginChallenge.Amr,
        Cif:                 cif,
    }
}
