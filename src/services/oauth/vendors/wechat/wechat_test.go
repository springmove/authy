package wechat

import (
	"fmt"
	"testing"
)

func getMiniProgram() *MiniProgram {
	mp := MiniProgram{}
	mp.Init()

	return &mp
}

func TestMiniProgram(t *testing.T) {
	mp := getMiniProgram()
	resp, err := mp.miniprogramAuth("", "", "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rt := resp.toAuthResponseData()
	fmt.Println(rt)
}

func TestAccessToken(t *testing.T) {
	mp := getMiniProgram()
	resp, err := getAccessToken(&ReqAccessToken{
		http:   mp.Http,
		AppID:  "",
		Secret: "",
	})

	if err != nil {
		t.Logf(err.Error())
		return
	}

	t.Log(resp)
}

func TestAuthMobile(t *testing.T) {
	mp := getMiniProgram()
	resp, err := mp.getMobileByAuthCode("", "", "")
	if err != nil {
		t.Logf(err.Error())
		return
	}

	t.Log(resp)
}
