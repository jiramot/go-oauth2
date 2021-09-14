package domains

type Token struct {
    GrantType    string
    ClientId     string
    ClientSecret string
    Code         string
    CodeVerifier string
}
