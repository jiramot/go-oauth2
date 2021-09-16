package ports

import "github.com/jiramot/go-oauth2/internal/core/domains"

type LoginChallengePort interface {
    SaveLoginChallenge(challenge *domains.LoginChallenge) error
    GetLoginChallenge(loginChallengeCode string) (*domains.LoginChallenge, error)
    RemoveLoginChallenge(loginChallengeCode string) error
}
