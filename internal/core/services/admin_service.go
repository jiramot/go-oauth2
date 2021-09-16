package services

import (
    "errors"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/mocks"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    "log"
)

type adminService struct {
    loginChallengePort ports.LoginChallengePort
}

func NewAdminService(loginChallengePort ports.LoginChallengePort) *adminService {
    return &adminService{
        loginChallengePort: loginChallengePort,
    }
}

func (svc *adminService) AcceptLogin(loginChallengeCode string, cif string) (domains.AuthorizationCode, error) {
    //TODO
    loginChallenge, _ := svc.loginChallengePort.GetLoginChallenge(loginChallengeCode)
    //Generate authorization code for client_id ... and save
    code := mocks.AuthorizationCode
    if loginChallenge != nil {
        log.Println(loginChallenge)
        return domains.AuthorizationCode{
            Code: code,
        }, nil
    } else {
        return domains.AuthorizationCode{}, errors.New("Invalid request")
    }
}

func (svc *adminService) RejectLogin(loginChallengeCode string, cif string) error {
    return nil
}
