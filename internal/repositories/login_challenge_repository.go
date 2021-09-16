package repositories

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-redis/redis/v8"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "time"
)

type loginChallengeRepository struct {
    client *redis.Client
}

func NewLoginChallengeRepository(client *redis.Client) *loginChallengeRepository {
    return &loginChallengeRepository{
        client: client,
    }
}

func (repo *loginChallengeRepository) SaveLoginChallenge(challenge *domains.LoginChallenge) error {
    ctx := context.Background()
    json, err := json.Marshal(challenge)
    if err != nil {
        return err
    }
    errs := repo.client.Set(ctx, fmt.Sprintf("code_challenge:%s", challenge.LoginChallengeCode), json, 5*time.Minute).Err()
    return errs
}

func (repo *loginChallengeRepository) GetLoginChallenge(loginChallengeCode string) (*domains.LoginChallenge, error) {
    ctx := context.Background()
    val, err := repo.client.Get(ctx, fmt.Sprintf("code_challenge:%s", loginChallengeCode)).Result()
    if err != nil {
        return nil, err
    }
    b := []byte(val)
    challenge := &domains.LoginChallenge{}
    err = json.Unmarshal(b, challenge)
    if err != nil {
        return nil, err
    }
    return challenge, nil
}
