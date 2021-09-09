package handlers

import (
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    "github.com/labstack/echo/v4"
    "net/http"
)

type PublicHttpHandler struct {
    authorizationUseCase ports.AuthorizationUseCase
    tokenUseCase         ports.TokenUseCase
}

func NewPublicHandler(
    authorizationUseCase ports.AuthorizationUseCase,
    tokenUseCase ports.TokenUseCase,
) *PublicHttpHandler {

    return &PublicHttpHandler{
        authorizationUseCase: authorizationUseCase,
        tokenUseCase:         tokenUseCase,
    }
}

func (hdl *PublicHttpHandler) Authorization(ctx echo.Context) error {
    request := new(AuthorizationRequest)
    if err := ctx.Bind(request); err != nil {
        return ctx.JSON(http.StatusBadRequest, nil)
    }
    if err := ctx.Validate(request); err != nil {
        return err
    }
    response, err := hdl.authorizationUseCase.AuthorizationCode(request.Amr, request.ClientId, request.RedirectUrl, request.Scope)

    if err != nil {
        return ctx.String(http.StatusBadRequest, "Bad request")
    }
    return ctx.JSON(http.StatusOK, response)
}

func (hdl *PublicHttpHandler) Token(ctx echo.Context) error {
    request := new(TokenRequest)
    if err := ctx.Bind(request); err != nil {
        return ctx.JSON(http.StatusBadRequest, nil)
    }
    if err := ctx.Validate(request); err != nil {
        return err
    }
    token := domains.Token{
        GrantType:    request.GrantType,
        ClientId:     request.ClientId,
        ClientSecret: request.ClientSecret,
    }
    accessToken, _ := hdl.tokenUseCase.GenerateToken(token)

    return ctx.JSON(http.StatusOK, accessToken)
}

type AuthorizationRequest struct {
    ResponseType string `query:"response_type" validate:"required"`
    Amr          string `query:"Amr" default:"sso"`
    ClientId     string `query:"client_id" validate:"required"`
    RedirectUrl  string `query:"redirect_url"`
    Scope        string `query:"Scope"`
}

type TokenRequest struct {
    GrantType    string `json:"grant_type" form:"grant_type" validate:"required"`
    ClientId     string `json:"client_id" form:"client_id"validate:"required"`
    ClientSecret string `json:"client_secret" form:"client_secret" validate:"required"`
}
