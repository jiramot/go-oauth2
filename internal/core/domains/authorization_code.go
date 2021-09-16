package domains

type AuthorizationCode struct {
    LoginChallengeCode  string `json:"login_challenge_code"`
    ClientId            string `json:"client_id"`
    CodeChallenge       string `json:"code_challenge"`
    CodeChallengeMethod string `json:"code_challenge_method"`
    State               string `json:"state"`
    RedirectUrl         string `json:"redirect_url"`
    Scope               string `json:"scope"`
    Amr                 string `json:"amr"`
    Code                string `json:"code"`
    Cif                 string `json:"cif"`
}
