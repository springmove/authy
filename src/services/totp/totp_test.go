package totp

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"testing"
	"time"
)

func TestTotp(t *testing.T) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Ashibro",
		AccountName: "Ashibro",
		Period:      9999,
	})

	if err != nil {
		return
	}

	opts := totp.ValidateOpts{
		Period: 2,
		Digits: 6,
	}
	fmt.Println(key.Secret())
	code, _ := totp.GenerateCodeCustom(key.Secret(), time.Now(), opts)
	fmt.Println(code)

	time.Sleep(3 * time.Second)
	fmt.Println(totp.ValidateCustom(code, key.Secret(), time.Now(), opts))
}
