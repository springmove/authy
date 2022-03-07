package totp

const (
	RequestFailed = "TotpRequestFailed"
)

type GenerateBody struct {
	Code string `json:"code"`
	Key  string `json:"key"`
}

type RequestEndpoint struct {
	Endpoint string `json:"endpoint"`
	Account  string `json:"account"`
}

type ValidateRequest struct {
	RequestEndpoint
	GenerateBody
}
