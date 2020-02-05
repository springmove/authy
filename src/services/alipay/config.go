package alipay

import "github.com/linshenqi/authy/src/services/oauth"

type Config struct {
	Endpoints map[string]oauth.Endpoint `yaml:"endpoints"`
}

func (s *Config) ConfigName() string {
	return ServiceName
}

func (s *Config) Validate() error {
	return nil
}

func (s *Config) Default() interface{} {
	return &Config{
		Endpoints: map[string]oauth.Endpoint{
			"default": {
				AppID:     "AppID",
				AppSecret: "AppSecret",
				PublicKey: "PublicKey",
			},
		},
	}
}
