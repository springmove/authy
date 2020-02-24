package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwt2 "github.com/linshenqi/authy/src/services/jwt"
	"github.com/linshenqi/authy/src/services/oauth"
	"github.com/linshenqi/authy/src/services/totp"
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

	_, err := authy.OAuth(oauth.Request{
		Provider: oauth.WeChat,
		Endpoint: "ashibro_dev",
		Code:     "1234",
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

	token, err := authy.JwtSigner(claims)
	if err == nil {
		fmt.Println(token)
	}

	//time.Sleep(2 * time.Second)
	//claims["id"] = "wef"
	_, err = authy.JwtAuther(jwt2.Request{
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

func TestTotp(t *testing.T) {
	authy := getApi()

	body, err := authy.TotpGenerate("sms")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(body.Code)
	fmt.Println(body.Key)

	success, err := authy.TotpValidate(totp.ValidateRequest{
		RequestEndpoint: totp.RequestEndpoint{Endpoint: "sms"},
		GenerateBody: totp.GenerateBody{
			Code: body.Code,
			Key:  body.Key,
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if success {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
}
