package api

import (
	"github.com/linshenqi/authy/src/services/auth"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthWechat(t *testing.T) {
	cfg := AuthyConfig{
		Url:          "http://127.0.0.1:9090",
		Timeout:      3,
		Headers:      map[string]string{"Content-Type": "application/json"},
		PushInterval: 1,
		MaxRetry:     1,
	}

	authy := Authy{}
	authy.InitService(&cfg)

	_, err := authy.Auth(&auth.AuthBody{
		Type: auth.AUTH_TYPE_WECHAT,
		Data: auth.AuthWeChat{
			Code: "1234",
		},
	})

	assert.Nil(t, err)

	//resp.Data.(wechat.WXAuthResponse)
	//fmt.Printf("%s\n", wx)
}
