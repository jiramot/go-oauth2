package repositories

import (
    "github.com/golang-jwt/jwt"
    "github.com/google/uuid"
    "github.com/jiramot/go-oauth2/internal/core/domains"
)

type tokenizeRepository struct {
}

func NewTokenizeRepository() *tokenizeRepository {
    return &tokenizeRepository{}
}

const (
    secret = "secret"
)

func (repo *tokenizeRepository) CreateToken(payload *domains.TokenPayload) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
    token.Header["kid"] = uuid.New().String()
    signedToken, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

func (repo *tokenizeRepository) ValidateToken(token string) (*domains.TokenPayload, error) {
    return nil, nil
}
