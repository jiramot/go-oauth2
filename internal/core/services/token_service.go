package services

import (
    "errors"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/mocks"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    "github.com/jiramot/go-oauth2/internal/pkg/pkce"
)

type tokenService struct {
    tokenizePort ports.TokenizePort
}

func NewTokenService(tokenizePort ports.TokenizePort) *tokenService {
    return &tokenService{
        tokenizePort: tokenizePort,
    }
}

func (svc *tokenService) GenerateToken(token domains.Token) (string, error) {
    //check client secret or code verifier

    if isValidCodeVerifier, _ := pkce.VerifyCodeChallenge(token.CodeVerifier, mocks.LoginRequest.CodeChallenge, mocks.LoginRequest.CodeChallengeMethod);
        token.ClientId == mocks.Client.ClientId &&

            token.Code == mocks.AuthorizationCode &&
            token.GrantType == "authorization_code" &&
            isValidCodeVerifier {

        payload := domains.NewTokenPayload(token.ClientId, "cif", "openid profile", "next", domains.TokenTtl)
        tokenString, _ := svc.tokenizePort.CreateToken(payload)
        return tokenString, nil
    } else {
        return "", errors.New("")
    }

}

func (svc *tokenService) IntrospectToken(token string) (*domains.TokenPayload, error) {
    payload, err := svc.tokenizePort.ValidateToken(token)
    if err != nil {
        return &domains.TokenPayload{}, err
    }
    return payload, nil
}
