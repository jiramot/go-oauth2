package domains

type AuthorizationCode struct {
    Code  string `json:"code"`
    State string `json:"state"`
}
