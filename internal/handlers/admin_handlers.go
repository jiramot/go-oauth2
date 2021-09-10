package handlers

import (
    "fmt"
    "github.com/jiramot/go-oauth2/internal/core/mocks"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    "github.com/jiramot/go-oauth2/internal/pkg"
    util "github.com/jiramot/go-oauth2/internal/pkg"
    "github.com/labstack/echo/v4"
    "net/http"
)

type AdminHttpHandler struct {
    adminUseCase    ports.AdminAcceptLoginUseCase
    tokenUseCase ports.TokenUseCase
}

func NewAdminHttpHandler(adminUseCase ports.AdminAcceptLoginUseCase, tokenUseCase ports.TokenUseCase) *AdminHttpHandler {
    return &AdminHttpHandler{
        adminUseCase:    adminUseCase,
        tokenUseCase: tokenUseCase,
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

func (hdl *AdminHttpHandler) IntrospectToken(ctx echo.Context) error {
    request := new(TokenIntrospectRequest)
    if err := util.BindAndValidateRequest(ctx, request); err != nil {
        return ctx.String(http.StatusBadRequest, "")
    }
    payload, err := hdl.tokenUseCase.IntrospectToken(request.Token)
    if err != nil {
        return ctx.String(http.StatusBadRequest, "Bad json")
    }
    return ctx.JSON(http.StatusOK, payload)
}

type (
    acceptLoginChallengeResponse struct {
        RedirectTo string `json:"redirect_to"`
    }
    acceptLoginChallengeRequest struct {
        Cif string `json:"cif" validate:"required"`
    }
    TokenIntrospectRequest struct {
        Token string `form:"token" validate:"required"`
    }
)
