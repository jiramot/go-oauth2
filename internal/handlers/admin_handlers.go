package handlers

import (
    "fmt"
    "github.com/jiramot/go-oauth2/internal/core/usecases"
    "github.com/jiramot/go-oauth2/internal/pkg"
    util "github.com/jiramot/go-oauth2/internal/pkg"
    "github.com/labstack/echo/v4"
    "net/http"
)

type AdminHttpHandler struct {
    adminUseCase usecases.AdminAcceptLoginUseCase
    tokenUseCase usecases.TokenUseCase
}

func NewAdminHttpHandler(adminUseCase usecases.AdminAcceptLoginUseCase, tokenUseCase usecases.TokenUseCase) *AdminHttpHandler {
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
        redirectUrl := fmt.Sprintf("%s?code=%s&state=%s", authCode.RedirectUrl, authCode.Code, authCode.State)
        response := acceptLoginChallengeResponse{
            RedirectTo: redirectUrl,
            Code:       authCode.Code,
            State:      authCode.State,
        }
        return ctx.JSON(http.StatusOK, response)
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

type acceptLoginChallengeResponse struct {
    RedirectTo string `json:"redirect_url"`
    Code       string `json:"code"`
    State      string `json:"state"`
}

type acceptLoginChallengeRequest struct {
    Cif string `json:"cif" validate:"required"`
}

type TokenIntrospectRequest struct {
    Token string `form:"token" validate:"required"`
}
