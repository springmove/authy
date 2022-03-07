package totp

import (
	"fmt"
	"testing"
	"time"
)

const (
	EndpointSMS = "sms"
)

func getService() *Service {
	srv := Service{cfg: Config{Endpoints: map[string]Endpoint{
		EndpointSMS: {
			Issuer:  EndpointSMS,
			CodeLen: 6,
			Expiry:  10 * time.Second,
		},
	}}}

	return &srv
}

func TestTotp(t *testing.T) {
	srv := getService()

	code, key, err := srv.Gererate(EndpointSMS, "1")
	if err != nil {
		return
	}

	fmt.Printf("%s::%s\n", code, key)

	code2, key2, err := srv.Gererate(EndpointSMS, "1")
	if err != nil {
		return
	}

	fmt.Printf("%s::%s\n", code2, key2)

	rt, _ := srv.Validate(EndpointSMS, code, key)
	fmt.Println(rt)
}
