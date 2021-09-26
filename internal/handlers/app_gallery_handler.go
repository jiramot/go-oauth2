package handlers

import (
    "github.com/jiramot/go-oauth2/internal/core/usecases"
    "github.com/labstack/echo/v4"
    "net/http"
)

type AppGalleryHandler struct {
    openAppGalleryUseCase usecases.OpenAppGalleryUseCase
}

func NewAppGalleryHandler(
    openAppGalleryUseCase usecases.OpenAppGalleryUseCase,
) *AppGalleryHandler {
    return &AppGalleryHandler{
        openAppGalleryUseCase: openAppGalleryUseCase,
    }
}

func (hdl *AppGalleryHandler) OpenAppGallery(ctx echo.Context) error {
    appId := ctx.Param("id")
    redirectUrl, err := hdl.openAppGalleryUseCase.OpenFormAppId(appId)
    if err != nil {
        return ctx.String(http.StatusBadRequest, "Bad request")
    }
    return ctx.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}
