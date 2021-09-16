package services

import (
    "errors"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/mocks"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    "github.com/jiramot/go-oauth2/internal/pkg/pkce"
)

type tokenService struct {
    tokenizePort          ports.TokenizePort
    authorizationCodePort ports.AuthorizationCodePort
}

func NewTokenService(
    tokenizePort ports.TokenizePort,
    authorizationCodePort ports.AuthorizationCodePort,
) *tokenService {
    return &tokenService{
        tokenizePort:          tokenizePort,
        authorizationCodePort: authorizationCodePort,
    }
}

func (svc *tokenService) CreateTokenForAuthorizationCodeGrantType(token domains.Token) (string, error) {
    authorizationCode, _ := svc.authorizationCodePort.GetAuthorizationCodeFromCode(token.Code)
    svc.authorizationCodePort.RemoveAuthorizationCodeFromCode(token.Code)
    if authorizationCode == nil {
        return "", errors.New("")
    }
    isValidClientSecret := false
    if token.ClientSecret != "" {
        client, err := mocks.NewClientDb().FindClientByClientId(token.ClientId)
        if err != nil {
            return "", errors.New("invalid request")
        }
        if token.ClientSecret == client.ClientSecret {
            isValidClientSecret = true
        }
    }
    isValidCodeVerifier, _ := pkce.VerifyCodeChallenge(
        token.CodeVerifier,
        authorizationCode.CodeChallenge,
        authorizationCode.CodeChallengeMethod,
    )

    if token.ClientId == authorizationCode.ClientId && (isValidCodeVerifier || isValidClientSecret) {
        payload := domains.NewTokenPayload(
            authorizationCode.ClientId,
            authorizationCode.Cif,
            authorizationCode.Scope,
            authorizationCode.Amr,
            domains.TokenTtl,
        )
        tokenString, _ := svc.tokenizePort.CreateToken(payload)
        return tokenString, nil
    } else {
        return "", errors.New("invalid request")
    }
}

func (svc *tokenService) IntrospectToken(token string) (*domains.TokenPayload, error) {
    payload, err := svc.tokenizePort.ValidateToken(token)
    if err != nil {
        return &domains.TokenPayload{}, err
    }
    return payload, nil
}
