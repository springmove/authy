package oauth

import (
	"github.com/springmove/authy/src/base"
	"github.com/springmove/sptty"
)

type Config struct {
	sptty.BaseConfig

	Endpoints map[string]base.Endpoint `yaml:"endpoints"`
}

func (s *Config) ConfigName() string {
	return base.ServiceOAuth
}
