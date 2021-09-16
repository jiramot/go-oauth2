package main

import (
    "net/http"

    "github.com/go-playground/validator"
    "github.com/go-redis/redis/v8"
    "github.com/jiramot/go-oauth2/internal/core/services"
    "github.com/jiramot/go-oauth2/internal/handlers"
    "github.com/jiramot/go-oauth2/internal/repositories"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    loginChallengeRepository := repositories.NewLoginChallengeRepository(rdb)
    authorizationService := services.NewAuthorizationService(loginChallengeRepository)
    tokenizeRepository := repositories.NewTokenizeRepository()
    tokenService := services.NewTokenService(tokenizeRepository)
    hdl := handlers.NewPublicHandler(authorizationService, tokenService)

    e := echo.New()
    e.Use(middleware.Logger())
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
