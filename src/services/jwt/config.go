package jwt

import (
	"time"

	"github.com/springmove/authy/src/base"
	"github.com/springmove/sptty"
)

type Config struct {
	sptty.BaseConfig

	Expiry time.Duration `yaml:"expiry"`
}

func (s *Config) ConfigName() string {
	return base.ServiceJwt
}

func (s *Config) Default() interface{} {
	return &Config{
		Expiry: 24 * time.Hour,
	}
}
