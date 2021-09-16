package repositories

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-redis/redis/v8"
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "time"
)

type authorizationCodeRepository struct {
    client *redis.Client
}

func NewAuthorizationCodeRepository(client *redis.Client) *authorizationCodeRepository {
    return &authorizationCodeRepository{
        client: client,
    }
}

func (repo *authorizationCodeRepository) SaveAuthorizationCode(authorizationCode *domains.AuthorizationCode) error {
    ctx := context.Background()
    json, err := json.Marshal(authorizationCode)
    if err != nil {
        return err
    }
    errs := repo.client.Set(ctx, repo.getKey(authorizationCode.Code), json, 60*time.Second).Err()
    return errs
}

func (repo *authorizationCodeRepository) GetAuthorizationCodeFromCode(code string) (*domains.AuthorizationCode, error) {
    ctx := context.Background()
    val, err := repo.client.Get(ctx, repo.getKey(code)).Result()
    if err != nil {
        return nil, err
    }
    b := []byte(val)
    authorizationCode := &domains.AuthorizationCode{}
    err = json.Unmarshal(b, authorizationCode)
    if err != nil {
        return nil, err
    }
    return authorizationCode, nil
}

func (repo *authorizationCodeRepository) RemoveAuthorizationCodeFromCode(code string) error {
    ctx := context.Background()
    err := repo.client.Del(ctx, repo.getKey(code)).Err()
    return err
}

func (repo *authorizationCodeRepository) getKey(code string) string {
    return fmt.Sprintf("authorization_code:%s", code)
}
