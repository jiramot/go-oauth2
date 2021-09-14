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

func (svc *authorizationService) AuthorizationCode(
    amr string,
    clientId string,
    redirectUrl string,
    scope string,
    state string,
    codeChallenge string,
    codeChallengeMethod string,
    nonce string,
) (domains.Authorization, error) {
    //Check AMR Table
    //Check Client ID & redirect url & scope are match
    //Generate login challenge then SAVE challenge to authentication_core_request,
    if clientId == mocks.Client.ClientId {
        loginChallengeCode := mocks.LoginChallengeCode
        loginEndpoint := mocks.LoginEndpointUrl

        if amr == "next" {
            loginEndpoint = "next://login"
        }
        //Save requestId, codeChallenge, codeChallengeMethod
        mocks.NewLoginRequest(state, codeChallenge, codeChallengeMethod, loginChallengeCode)
        
        return domains.Authorization{
            LoginChallenge:   loginChallengeCode,
            LoginEndpointUrl: fmt.Sprintf("%s?login_challenge=%s", loginEndpoint, loginChallengeCode),
        }, nil
    }

    return domains.Authorization{}, errors.New("No client found")

}
