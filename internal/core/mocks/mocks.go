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
//var LoginChallengeCode = "123456"
var LoginEndpointUrl = "http://localhost:3000"
var AuthorizationCode = "123456789"

//type loginRequest struct {
//    State               string
//    CodeChallenge       string
//    CodeChallengeMethod string
//    LoginChallengeCode  string
//}
//
//var LoginRequest = &loginRequest{}
//
//func NewLoginRequest(
//    state string,
//    codeChallenge string,
//    codeChallengeMethod string,
//    loginChallengeCode string,
//) *loginRequest {
//    LoginRequest.State = state
//    LoginRequest.CodeChallenge = codeChallenge
//    LoginRequest.CodeChallengeMethod = codeChallengeMethod
//    LoginRequest.LoginChallengeCode = loginChallengeCode
//    return LoginRequest
//}
