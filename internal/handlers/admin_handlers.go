package handlers

import (
    "fmt"
    "github.com/jiramot/go-oauth2/internal/core/mocks"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    "github.com/jiramot/go-oauth2/internal/pkg"
    "github.com/labstack/echo/v4"
    "net/http"
)

type AdminHttpHandler struct {
    adminUseCase ports.AdminAcceptLoginUseCase
}

func NewAdminHttpHandler(adminUseCase ports.AdminAcceptLoginUseCase) *AdminHttpHandler {
    return &AdminHttpHandler{
        adminUseCase: adminUseCase,
    }
}

func (hdl *AdminHttpHandler) AcceptLoginChallenge(ctx echo.Context) error {
    loginChallengeCodeQueryPrams := ctx.QueryParam("login_challenge")
    request := new(acceptLoginChallengeRequest)
    if err := pkg.BindAndValidateRequest(ctx, request); err != nil {
        return ctx.String(http.StatusBadRequest, "")
    }
    if authCode, err := hdl.adminUseCase.AcceptLogin(loginChallengeCodeQueryPrams, request.Cif); err == nil {
        redirectUrl := fmt.Sprintf("%s?code=%s", mocks.Client.PartnerEndpoint, authCode.Code)
        return ctx.JSON(http.StatusOK, acceptLoginChallengeResponse{RedirectTo: redirectUrl})
    } else {
        return ctx.JSON(http.StatusBadRequest, nil)
    }
}

type (
    acceptLoginChallengeResponse struct {
        RedirectTo string `json:"redirect_to"`
    }
    acceptLoginChallengeRequest struct {
        Cif string `json:"cif" validate:"required"`
    }
)
