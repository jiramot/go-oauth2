package handlers

import (
    "github.com/jiramot/go-oauth2/internal/core/domains"
    "github.com/jiramot/go-oauth2/internal/core/ports"
    util "github.com/jiramot/go-oauth2/internal/pkg"
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
    if err := util.BindAndValidateRequest(ctx, request); err != nil {
        return ctx.String(http.StatusBadRequest, "")
    }
    response, err := hdl.authorizationUseCase.AuthorizationCode(request.Amr, request.ClientId, request.RedirectUrl, request.Scope)

    if err != nil {
        return ctx.String(http.StatusBadRequest, "Bad request")
    }
    return ctx.JSON(http.StatusOK, response)
}

func (hdl *PublicHttpHandler) Token(ctx echo.Context) error {
    request := new(TokenRequest)
    if err := util.BindAndValidateRequest(ctx, request); err != nil {
        return ctx.String(http.StatusBadRequest, "")
    }
    token := domains.Token{
        GrantType:    request.GrantType,
        ClientId:     request.ClientId,
        ClientSecret: request.ClientSecret,
        Code:         request.Code,
    }
    accessToken, _ := hdl.tokenUseCase.GenerateToken(token)
    response := tokenResponse{
        TokenType:   "Bearer",
        AccessToken: accessToken,
    }
    return ctx.JSON(http.StatusOK, response)
}

func (hdl *PublicHttpHandler) IntrospectToken(ctx echo.Context) error {
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
    Code         string `json:"code"`
}

type tokenResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
}

type TokenIntrospectRequest struct {
    Token string `form:"token" validate:"required"`
}