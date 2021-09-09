package mocks

import "github.com/jiramot/go-oauth2/internal/core/domains"

var Clients = []domains.Client{
    domains.Client{
        ClientId:     "1234",
        ClientSecret: "secret",
        Scope:        "openid profile",
        RedirectUrl:  "http://partner.com",
        GrantType:    "authorization_code refresh_token",
    },
}
