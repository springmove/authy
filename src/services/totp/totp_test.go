package totp

import (
	"fmt"
	"testing"
	"time"

	"github.com/pquerna/otp/totp"
)

func TestTotp(t *testing.T) {
	// fmt.Println(key.Secret())
	for i := 0; i < 100; i++ {

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
		code, _ := totp.GenerateCodeCustom(key.Secret(), time.Now(), opts)
		fmt.Println(code)
	}

	// time.Sleep(3 * time.Second)
	// fmt.Println(totp.ValidateCustom(code, key.Secret(), time.Now(), opts))
}
