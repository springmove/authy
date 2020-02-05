package auth

const (
	AuthFailed = "AuthFailed"
)

const (
	WeChatOAuth       = "wechat_oauth"
	WeChatMiniProgram = "wechat_miniprogram"
	AliPayOAuth       = "alipay_oauth"
)

type Request struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Response struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
