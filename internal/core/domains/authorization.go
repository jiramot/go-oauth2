package domains

type Authorization struct {
    LoginChallenge string `json:"login_challenge"`
    LoginEndpointUrl string `json:"login_endpoint_url"`
}
