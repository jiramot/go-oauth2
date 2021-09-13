package services

import (
    "errors"
    "fmt"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/mocks"
)

type authorizationService struct {
}

func NewAuthorizationService() *authorizationService {
    return &authorizationService{}
}

func (svc *authorizationService) AuthorizationCode(amr string, clientId string, redirectUrl string, scope string) (domains.Authorization, error) {
    //Check AMR Table
    //Check Client ID & redirect url & scope are match
    //Generate login challenge then SAVE challenge to authentication_core_request,
    if clientId == mocks.Client.ClientId {
        loginChallengeCode := mocks.LoginChallengeCode
        loginEndpoint := mocks.LoginEndpointUrl
        if amr == "next" {
            loginEndpoint = "next://login"
        }

        return domains.Authorization{
            LoginChallenge:   loginChallengeCode,
            LoginEndpointUrl: fmt.Sprintf("%s?challenge_code=%s", loginEndpoint, loginChallengeCode),
        }, nil
    }

    return domains.Authorization{}, errors.New("No client found")

}
