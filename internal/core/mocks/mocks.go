package mocks

type client struct {
    ClientId        string
    ClientSecret    string
    PartnerEndpoint string
}

var Client = client{
    ClientId:        "1234",
    ClientSecret:    "12345",
    PartnerEndpoint: "https://partner.com",
}
var LoginChallengeCode = "123456"
var LoginEndpointUrl = "http://localhost:3000"
var AuthorizationCode = "123456789"
