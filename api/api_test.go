package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/linshenqi/authy/src/services/auth"
	jwt2 "github.com/linshenqi/authy/src/services/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthWechat(t *testing.T) {
	cfg := AuthyConfig{
		Url: "http://127.0.0.1:9090",
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

func TestJwt(t *testing.T) {
	cfg := AuthyConfig{
		Url: "http://127.0.0.1:10001",
	}

	authy := Authy{}
	authy.InitService(&cfg)

	claims := jwt.MapClaims{
		"id":  "wefwef234",
		"wef": "232r2r2",
	}

	token, err := authy.JwtSigner(&claims)
	if err == nil {
		fmt.Println(token)
	}

	//claims["id"] = "wef"
	_, err = authy.JwtAuther(&jwt2.JwtAutherRequest{
		Token:  token,
		Claims: claims,
	})

	if err == nil {
		fmt.Println("ok")
	} else {
		fmt.Println("fail")
	}

	c, _ := authy.JwtParser(&jwt2.JwtAutherRequest{
		Token: token,
	})
	fmt.Println(c)
}
