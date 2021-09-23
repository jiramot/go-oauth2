package services

import (
    "errors"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    "github.com/jiramot/go-oauth2/internal/pkg/pkce"
    "time"
)

type tokenService struct {
    tokenizePort          ports.TokenizePort
    authorizationCodePort ports.AuthorizationCodePort
    clients               domains.Clients
    accessTokenDuration   time.Duration
}

func NewTokenService(
    tokenizePort ports.TokenizePort,
    authorizationCodePort ports.AuthorizationCodePort,
    clients domains.Clients,
    accessTokenDuration time.Duration,
) *tokenService {
    return &tokenService{
        tokenizePort:          tokenizePort,
        authorizationCodePort: authorizationCodePort,
        clients:               clients,
        accessTokenDuration:   accessTokenDuration,
    }
}

func (svc *tokenService) CreateTokenForAuthorizationCodeGrantType(token domains.Token) (*domains.AccessToken, error) {
    authorizationCode, _ := svc.authorizationCodePort.GetAuthorizationCodeFromCode(token.Code)
    svc.authorizationCodePort.RemoveAuthorizationCodeFromCode(token.Code)
    if authorizationCode == nil {
        return nil, errors.New("")
    }
    isValidClientSecret := false
    if token.ClientSecret != "" {
        client, err := svc.clients.FindClientByClientId(token.ClientId)
        if err != nil {
            return nil, errors.New("invalid request")
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
            svc.accessTokenDuration,
        )
        tokenString, _ := svc.tokenizePort.CreateToken(payload)
        accessToken := &domains.AccessToken{
            AccessToken: tokenString,
            ExpireAt:    payload.ExpiredAt,
            TokenType:   "Bearer",
        }
        return accessToken, nil
    } else {
        return nil, errors.New("invalid request")
    }
}

func (svc *tokenService) IntrospectToken(token string) (*domains.TokenPayload, error) {
    payload, err := svc.tokenizePort.ValidateToken(token)
    if err != nil {
        return &domains.TokenPayload{}, err
    }
    return payload, nil
}
