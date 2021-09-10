package domains

type Token struct {
    GrantType    string `json:"grant_type" validate:"required"`
    ClientId     string `json:"client_id" validate:"required"`
    ClientSecret string `json:"client_secret" validate:"required"`
    Code         string `json:"code"`
}
