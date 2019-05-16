package auth

const (
	AUTHY_ERR_REQUEST_BODY = "0001"
	AUTHY_ERR_TYPE         = "0002"
	AUTHY_ERR_FAIL         = "0003"
)

const (
	AUTH_TYPE_WECHAT = "wechat"
)

type AuthBody struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type AuthWeChat struct {
	Code string `json:"code"`
}
