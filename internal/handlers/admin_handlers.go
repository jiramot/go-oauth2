package handlers

import (
    "fmt"
    "github.com/jiramot/go-oauth2/internal/core/mocks"
    usecases2 "github.com/jiramot/go-oauth2/internal/core/usecases"
    "github.com/jiramot/go-oauth2/internal/pkg"
    util "github.com/jiramot/go-oauth2/internal/pkg"
    "github.com/labstack/echo/v4"
    "net/http"
)

type AdminHttpHandler struct {
    adminUseCase usecases2.AdminAcceptLoginUseCase
    tokenUseCase usecases2.TokenUseCase
}

func NewAdminHttpHandler(adminUseCase usecases2.AdminAcceptLoginUseCase, tokenUseCase usecases2.TokenUseCase) *AdminHttpHandler {
    return &AdminHttpHandler{
        adminUseCase: adminUseCase,
        tokenUseCase: tokenUseCase,
    }
}

func (hdl *AdminHttpHandler) AcceptLoginChallenge(ctx echo.Context) error {
    loginChallengeCode := ctx.QueryParam("login_challenge")
    request := new(acceptLoginChallengeRequest)
    if err := pkg.BindAndValidateRequest(ctx, request); err != nil {
        return ctx.String(http.StatusBadRequest, "Bad request")
    }
    if authCode, err := hdl.adminUseCase.AcceptLogin(loginChallengeCode, request.Cif); err == nil {
        client, _ := mocks.NewClientDb().FindClientByClientId(authCode.ClientId)
        redirectUrl := fmt.Sprintf("%s?code=%s&state=%s", client.RedirectUrl, authCode.Code, authCode.State)
        return ctx.JSON(http.StatusOK, acceptLoginChallengeResponse{RedirectTo: redirectUrl})
    } else {
        return ctx.String(http.StatusBadRequest, "Bad request")
    }
}

func (hdl *AdminHttpHandler) IntrospectToken(ctx echo.Context) error {
    request := new(TokenIntrospectRequest)
    if err := util.BindAndValidateRequest(ctx, request); err != nil {
        return ctx.String(http.StatusBadRequest, "")
    }
    payload, err := hdl.tokenUseCase.IntrospectToken(request.Token)
    if err != nil {
        return ctx.String(http.StatusBadRequest, "Bad request")
    }
    return ctx.JSON(http.StatusOK, payload)
}

type (
    acceptLoginChallengeResponse struct {
        RedirectTo string `json:"redirect_url"`
    }
    acceptLoginChallengeRequest struct {
        Cif string `json:"cif" validate:"required"`
    }
    TokenIntrospectRequest struct {
        Token string `form:"token" validate:"required"`
    }
)
