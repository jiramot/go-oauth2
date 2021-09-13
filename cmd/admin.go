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
    adminService := services.NewAdminService()

    tokenizeRepository := repositories.NewTokenizeRepository()
    tokenService := services.NewTokenService(tokenizeRepository)

    adminHdl := handlers.NewAdminHttpHandler(adminService, tokenService)

    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.CORS()) //FOR TESTING FRONTEND
    e.Validator = &AdminValidator{validator: validator.New()}
    e.POST("oauth2/auth/request/login/accept", adminHdl.AcceptLoginChallenge)
    e.POST("/oauth2/introspect", adminHdl.IntrospectToken)
    e.Logger.Fatal(e.Start(":8081"))
}

type AdminValidator struct {
    validator *validator.Validate
}

func (cv *AdminValidator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}

