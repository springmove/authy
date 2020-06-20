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
	resp, err := mp.miniprogramAuth("wx77d83a2aa6c324ab", "2c6ecb6fe8a0394715704149a6afc56b", "071AZ3px1Q7Eoc0Pomnx1iWUox1AZ3pC")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rt := resp.toAuthResponseData()
	fmt.Println(rt)
}
