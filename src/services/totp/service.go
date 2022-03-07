package totp

import (
	"errors"
	"time"

	"github.com/linshenqi/authy/src/base"
	"github.com/linshenqi/sptty"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type Service struct {
	sptty.BaseService

	cfg Config
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(base.ServiceTotp, &s.cfg); err != nil {
		return err
	}

	app.AddRoute("PUT", "/totp-generate", s.putGenerate)
	app.AddRoute("PUT", "/totp-validate", s.putValidate)

	return nil
}

func (s *Service) ServiceName() string {
	return base.ServiceTotp
}

func (s *Service) Gererate(endpoint string, account string) (string, string, error) {
	ep, exist := s.cfg.Endpoints[endpoint]
	if !exist {
		return "", "", errors.New("Endpoint Not Found ")
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      ep.Issuer,
		AccountName: account,
		Period:      9999,
	})

	if err != nil {
		return "", "", err
	}

	code, err := totp.GenerateCodeCustom(key.Secret(), time.Now(), totp.ValidateOpts{
		Period: uint(ep.Expiry.Seconds()),
		Digits: otp.Digits(ep.CodeLen),
	})

	if err != nil {
		return "", "", err
	}

	return code, key.Secret(), nil
}

func (s *Service) Validate(endpoint string, code string, key string) (bool, error) {
	ep, exist := s.cfg.Endpoints[endpoint]
	if !exist {
		return false, errors.New("Endpoint Not Found ")
	}

	return totp.ValidateCustom(code, key, time.Now(), totp.ValidateOpts{
		Period: uint(ep.Expiry.Seconds()),
		Digits: otp.Digits(ep.CodeLen),
	})
}
