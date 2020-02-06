package oauth

const (
	ErrAuthFailed       = "ErrAuthFailed"
	ErrEndpointNotFound = "EndpointNotFound"
)

const (
	WeChatOAuth       = "wechat_oauth"
	WeChatMiniProgram = "wechat_miniprogram"
	AliPayOAuth       = "alipay_oauth"
)

type Endpoint struct {
	AppID     string `yaml:"app_id" json:"app_id"`
	AppSecret string `yaml:"app_secret" json:"app_secret"`
	PublicKey string `yaml:"public_key" json:"public_key"`
}

type Request struct {
	Type string `json:"type"`

	Endpoint string `json:"Endpoint"`
	Code     string `json:"Code"`
}

type Response struct {
	Type string `json:"type"`

	OpenID  string `json:"open_id"`
	UnionID string `json:"union_id"`
	Name    string `json:"name"`

	// 0:未知 1:男 2:女
	Gender int `json:"gender"`

	Province string `json:"province"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Avatar   string `json:"avatar"`
}
