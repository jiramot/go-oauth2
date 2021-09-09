package services

import (
    "errors"
    "github.com/jiramot/go-oauth2/internal/core/domains"
)

type tokenService struct {
}

func NewTokenService() *tokenService {
    return &tokenService{
    }
}

func (svc *tokenService) GenerateToken(token domains.Token) (string, error) {
    if token.ClientId == "1234" && token.ClientSecret == "12345" {
        return "jwt", nil
    } else {
        return "", errors.New("")
    }

}
