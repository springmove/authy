package wechat

const (
	WX_OK               = 0
	WX_ERR_SYS_BUSY     = -1
	WX_ERR_INVALID_CODE = 40029
	WX_ERR_OVERLOAD     = 45011
)

type WXAuthResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
