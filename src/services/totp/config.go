package totp

import (
	"errors"
	"time"

	"github.com/springmove/authy/src/base"
	"github.com/springmove/sptty"
)

type Config struct {
	sptty.BaseConfig

	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Issuer  string        `yaml:"issuer"`
	CodeLen int           `yaml:"len"`
	Expiry  time.Duration `yaml:"expiry"`
}

func (s *Config) ConfigName() string {
	return base.ServiceTotp
}

func (s *Config) Validate() error {

	for _, v := range s.Endpoints {
		if v.Issuer == "" {
			return errors.New("Issuer Is Required ")
		}

		if v.CodeLen <= 0 {
			return errors.New("Code Len Must Be Greater Than 0 ")
		}
	}

	return nil
}

func (s *Config) Default() interface{} {
	return &Config{
		Endpoints: map[string]Endpoint{
			"default": {
				Issuer:  "default",
				CodeLen: 4,
				Expiry:  60 * time.Second,
			},
		},
	}
}
