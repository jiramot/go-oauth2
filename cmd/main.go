package main

import (
    "github.com/go-playground/validator"
    "github.com/jiramot/go-oauth2/internal/core/services"
    "github.com/jiramot/go-oauth2/internal/handlers"
    "github.com/jiramot/go-oauth2/internal/repositories"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "net/http"
)

func main() {
    adminSvc := services.NewAdminService()
    adminHdl := handlers.NewAdminHttpHandler(adminSvc)

    authorizationService := services.NewAuthorizationService()
    tokenizeRepository := repositories.NewTokenizeRepository()
    tokenService := services.NewTokenService(tokenizeRepository)
    hdl := handlers.NewPublicHandler(authorizationService, tokenService)

    e := echo.New()
    e.Use(middleware.Logger())
    e.Validator = &CustomValidator{validator: validator.New()}
    e.GET("/oauth2/auth", hdl.Authorization)
    e.POST("/oauth2/token", hdl.Token)
    //Admin
    e.POST("oauth2/auth/request/login/accept", adminHdl.AcceptLoginChallenge)
    e.Logger.Fatal(e.Start(":8080"))
}

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}
