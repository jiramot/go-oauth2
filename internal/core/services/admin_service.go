package services

import (
    "errors"
    "github.com/jiramot/go-oauth2/internal/core/domains"
)

type adminService struct {
}

func NewAdminService() *adminService {
    return &adminService{}
}

func (svc *adminService) AcceptLogin(loginChallengeCode string, cif string) (domains.AuthorizationCode, error) {
    //TODO
    //Check login challenge
    //Generate authorization code for client_id ... and save

    if loginChallengeCode == "12345" {
        return domains.AuthorizationCode{
            Code: "123456789",
        }, nil
    } else {
        return domains.AuthorizationCode{}, errors.New("Invalid request")
    }
}

func (svc *adminService) RejectLogin(loginChallengeCode string, cif string) error {
    return nil
}
