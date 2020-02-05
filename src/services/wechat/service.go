package wechat

import (
	"encoding/json"
	"errors"
	"github.com/linshenqi/sptty"
)

const (
	ServiceName = "wechat"
)

type Service struct {
	cfg Config

	OAuth       OAuth
	MiniProgram MiniProgram
}

func (s *Service) Init(app sptty.Sptty) error {
	if err := app.GetConfig(ServiceName, &s.cfg); err != nil {
		return err
	}

	s.OAuth = OAuth{}
	s.OAuth.init(s.cfg.Endpoints)

	s.MiniProgram = MiniProgram{}
	s.MiniProgram.init(s.cfg.Endpoints)

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

func preAuth(req interface{}, endpoints map[string]Endpoint) (*AuthData, *Endpoint, error) {
	if req == nil {
		return nil, nil, errors.New("Params Is Nil ")
	}

	body, _ := json.Marshal(req)
	authReq := AuthData{}
	if err := json.Unmarshal(body, &authReq); err != nil {
		return nil, nil, err
	}

	endpoint, exist := endpoints[authReq.Endpoint]
	if !exist {
		return nil, nil, errors.New("Endpoint Not Found ")
	}

	return &authReq, &endpoint, nil
}
