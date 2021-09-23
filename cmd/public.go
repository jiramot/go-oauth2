package main

import (
    "fmt"
    "github.com/jiramot/go-oauth2/internal/config"
    "net/http"
    "os"

    "github.com/go-playground/validator"
    "github.com/go-redis/redis/v8"
    "github.com/jiramot/go-oauth2/internal/core/services"
    "github.com/jiramot/go-oauth2/internal/handlers"
    "github.com/jiramot/go-oauth2/internal/repositories"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    config, err := config.Load()
    if err != nil {
        return
    }

    redisHost := os.Getenv("REDIS_HOST")
    if redisHost == "" {
        redisHost = "localhost"
    }
    redisPort := os.Getenv("REDIS_PORT")
    if redisPort == "" {
        redisPort = "6379"
    }
    rdb := redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
        Password: "",
        DB:       0,
    })

    loginChallengeRepository := repositories.NewLoginChallengeRepository(rdb)
    authorizationService := services.NewAuthorizationService(loginChallengeRepository, config.Client, config.LoginEndpoint)

    tokenizeRepository := repositories.NewTokenizeRepository(rdb, config.AccessTokenDuration)
    authorizationCodeRepository := repositories.NewAuthorizationCodeRepository(rdb)
    tokenService := services.NewTokenService(tokenizeRepository, authorizationCodeRepository, config.Client, config.AccessTokenDuration)

    hdl := handlers.NewPublicHandler(authorizationService, tokenService)

    e := echo.New()
    if config.AccessLogEnabled {
        e.Use(middleware.Logger())
    }
    e.Use(middleware.CORS())
    e.Validator = &PublicValidator{validator: validator.New()}
    e.GET("/oauth2/auth", hdl.RequestAuthorization)
    e.POST("/oauth2/token", hdl.Token)
    e.Logger.Fatal(e.Start(":8080"))
}

type PublicValidator struct {
    validator *validator.Validate
}

func (cv *PublicValidator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}
