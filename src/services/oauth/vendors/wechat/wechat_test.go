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
