package services

import (
    "errors"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/ports"
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
    if token.ClientId == "1234" && token.ClientSecret == "12345" {
        payload := domains.NewTokenPayload(token.ClientId, "cif", "openid profile", "next", domains.TokenTtl)
        tokenString, _ := svc.tokenizePort.CreateToken(payload)
        return tokenString, nil
    } else {
        return "", errors.New("")
    }

}
