package repositories

import (
    "errors"
    "github.com/golang-jwt/jwt"
    "github.com/google/uuid"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "time"
)

type tokenizeRepository struct {
}

func NewTokenizeRepository() *tokenizeRepository {
    return &tokenizeRepository{}
}

const (
    privateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEICcsu6xbgzl3/XUzWhAHml/IEtDg1FsLjtpr1r7NsEPioAoGCCqGSM49
AwEHoUQDQgAEH8PhRqrSbgLvZ1tXv/XbIEToWVRuekJP5z+YDsTaNoDbAJCbLj/s
8qxA2sKeBGxSZ62RKSgsCVD9TQCcvGY/Hw==
-----END EC PRIVATE KEY-----`
    publicKey = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEH8PhRqrSbgLvZ1tXv/XbIEToWVRu
ekJP5z+YDsTaNoDbAJCbLj/s8qxA2sKeBGxSZ62RKSgsCVD9TQCcvGY/Hw==
-----END PUBLIC KEY-----`
)

func (repo *tokenizeRepository) CreateToken(payload *domains.TokenPayload) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
    ecPrivateKey, _ := jwt.ParseECPrivateKeyFromPEM([]byte(privateKey))
    token.Header["kid"] = uuid.New().String()
    signedToken, err := token.SignedString(ecPrivateKey)
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

func (repo *tokenizeRepository) ValidateToken(token string) (*domains.TokenPayload, error) {
    now := time.Now()
    keyFunc := func(token *jwt.Token) (interface{}, error) {
        _, ok := token.Method.(*jwt.SigningMethodECDSA)
        if !ok {
            return nil, errors.New("No key")
        }
        ecPublicKey, _ := jwt.ParseECPublicKeyFromPEM([]byte(publicKey))
        return ecPublicKey, nil
    }
    var payload = &domains.TokenPayload{}
    jwt.ParseWithClaims(token, payload, keyFunc)

    if payload.ExpiredAt <= now.Unix() {
        return nil, errors.New("Token expired")
    }
    return payload, nil
}
