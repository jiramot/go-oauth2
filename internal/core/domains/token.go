package domains

type Token struct {
    GrantType    string `json:"grant_type" form:"grant_type" validate:"required"`
    ClientId     string `json:"client_id" form:"client_id"validate:"required"`
    ClientSecret string `json:"client_secret" form:"client_secret" validate:"required"`
}
