package oauth

import (
	"github.com/linshenqi/authy/src/base"
	"github.com/linshenqi/sptty"
)

type Config struct {
	sptty.BaseConfig

	Endpoints map[string]base.Endpoint `yaml:"endpoints"`
}

func (s *Config) ConfigName() string {
	return base.ServiceOAuth
}
