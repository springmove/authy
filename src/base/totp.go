package base

const (
	ServiceTotp = "totp"
)

type IServiceTotp interface {
	Gererate(endpoint string, account string) (string, string, error)
	Validate(endpoint string, code string, key string) (bool, error)
}
