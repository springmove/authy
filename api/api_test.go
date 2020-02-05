package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/linshenqi/authy/src/services/auth"
	jwt2 "github.com/linshenqi/authy/src/services/jwt"
	"github.com/linshenqi/authy/src/services/wechat"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getApi() *Authy {
	cfg := Config{
		Url: "http://127.0.0.1:10001",
	}

	authy := Authy{}
	authy.Init(&cfg)

	return &authy
}

func TestAuthWechatOAuth(t *testing.T) {
	authy := getApi()

	_, err := authy.Auth(auth.Request{
		Type: auth.WeChatOAuth,
		Data: wechat.AuthData{
			Endpoint: "ashibro_dev",
			Code:     "1234",
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		assert.NotNil(t, err)
		return
	}

	assert.Nil(t, err)
}

func TestJwt(t *testing.T) {
	authy := getApi()

	claims := jwt.MapClaims{
		"id":  "wefwef234",
		"wef": "232r2r2",
	}

	token, err := authy.JwtSigner(&claims)
	if err == nil {
		fmt.Println(token)
	}

	//time.Sleep(2 * time.Second)
	//claims["id"] = "wef"
	_, err = authy.JwtAuther(&jwt2.Request{
		Token:  token,
		Claims: claims,
	})

	if err == nil {
		fmt.Println("ok")
	} else {
		fmt.Println(err.Error())
	}

	c, _ := authy.JwtParser(token)
	fmt.Println(c)
}
