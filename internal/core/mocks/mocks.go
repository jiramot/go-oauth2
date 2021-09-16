package mocks

type client struct {
    ClientId        string
    ClientSecret    string
    PartnerEndpoint string
}

var Client = client{
    ClientId:        "1234",
    ClientSecret:    "12345",
    PartnerEndpoint: "http://localhost:9100",
}
var LoginEndpointUrl = "http://localhost:3000"
