package domains

import (
    "errors"
)

type Clients []Client

type Client struct {
    ClientId     string `json:"client_id" mapstructure:"client_id"`
    ClientSecret string `json:"client_secret" mapstructure:"client_secret"`
    Scope        string `json:"scope" mapstructure:"scope"`
    RedirectUrl  string `json:"redirect_url" mapstructure:"redirect_url"`
    GrantType    string `json:"grant_type" mapstructure:"grant_type"`
}

func (clients *Clients) FindClientByClientId(clientId string) (*Client, error) {
    for _, c := range *clients {
        if c.ClientId == clientId {
            return &c, nil
        }
    }
    return nil, errors.New("")
}
