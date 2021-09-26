package main

import (
    "github.com/go-playground/validator"
    "github.com/jiramot/go-oauth2/internal/config"
    "github.com/jiramot/go-oauth2/internal/core/services"
    "github.com/jiramot/go-oauth2/internal/handlers"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "net/http"
)

func main() {
    config, err := config.Load()
    if err != nil {
        return
    }
    e := echo.New()
    if config.AccessLogEnabled {
        e.Use(middleware.Logger())
    }

    opeAppGalleryUseCase := services.NewAppGalleryService(config.Client)
    handler := handlers.NewAppGalleryHandler(opeAppGalleryUseCase)
    e.Validator = &AppGalleryValidator{validator: validator.New()}
    e.GET("/:id", handler.OpenAppGallery)
    e.Logger.Fatal(e.Start(":8082"))
}

type AppGalleryValidator struct {
    validator *validator.Validate
}

func (cv *AppGalleryValidator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}
