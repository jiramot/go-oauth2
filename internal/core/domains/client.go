package domains

type Client struct {
    ClientId string `json:"client_id"`
    ClientSecret string `json:"client_secret"`
    Scope string `json:"scope"`
    RedirectUrl string `json:"redirect_url"`
    GrantType string `json:"grant_type"`
}


