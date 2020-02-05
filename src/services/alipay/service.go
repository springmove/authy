package alipay

import "github.com/linshenqi/sptty"

const (
	ServiceName = "alipay"
)

type Service struct {
	cfg Config

	OAuth OAuth
}

func (s *Service) Init(app sptty.Sptty) error {
	if err := app.GetConfig(ServiceName, &s.cfg); err != nil {
		return err
	}

	s.OAuth = OAuth{}
	s.OAuth.Init(s.cfg.Endpoints)

	return nil
}

func (s *Service) Release() {

}

func (s *Service) Enable() bool {
	return true
}

func (s *Service) ServiceName() string {
	return ServiceName
}
