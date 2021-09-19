package mocks

import (
    "errors"
    "github.com/jiramot/go-oauth2/internal/core/domains"
)

type clientDb struct {
}

var clients = []domains.Client{
    domains.Client{
        ClientId:     "1234",
        ClientSecret: "12345",
        RedirectUrl:  "http://localhost:9100",
        Scope:        "openid profile",
    },
    domains.Client{
        ClientId:     "6789",
        ClientSecret: "12345",
        RedirectUrl:  "http://localhost:9100",
        Scope:        "openid profile",
    },
}

func NewClientDb() *clientDb {
    return &clientDb{}
}

func (client *clientDb) FindClientByClientId(clientId string) (*domains.Client, error) {
    for _, c := range clients {
        if c.ClientId == clientId {
            return &c, nil
        }
    }
    return nil, errors.New("")
}

var LoginEndpointUrl = "http://localhost:3000/login"
