package pkce

import (
    "crypto/sha256"
    "encoding/base64"
    "strings"
)

func VerifyCodeChallenge(codeVerifier string, codeChallenge string, codeChallengeMethod string) (bool, error) {
    verifyCodeChallenge, _ := CreateCodeChallenge(codeVerifier, codeChallengeMethod)
    if verifyCodeChallenge == codeChallenge {
        return true, nil
    } else {
        return false, nil
    }
}

func CreateCodeChallenge(codeVerifier string, codeChallengeMethod string) (string, error) {
    h := sha256.New()
    h.Write([]byte(codeVerifier))
    return base64Encode(h.Sum(nil)), nil
}

func base64Encode(msg []byte) string {
    encoded := base64.StdEncoding.EncodeToString(msg)
    encoded = strings.Replace(encoded, "+", "-", -1)
    encoded = strings.Replace(encoded, "/", "_", -1)
    encoded = strings.Replace(encoded, "=", "", -1)
    return encoded
}
