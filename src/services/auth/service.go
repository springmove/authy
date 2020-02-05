package auth

import (
	"errors"
	"github.com/linshenqi/sptty"
)

const (
	ServiceName = "auth"
)

type Service struct {
	authProviders map[string]IAuthProvider
}

func (s *Service) Init(app sptty.Sptty) error {
	app.AddRoute("POST", "/auth", s.postAuth)
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

func (s *Service) doAuth(req Request) (Response, error) {
	resp := Response{
		Type: req.Type,
	}

	provider, exist := s.authProviders[req.Type]
	if !exist {
		return resp, errors.New("Provider Not Found ")
	}

	respData, err := provider.Auth(req.Data)
	if err != nil {
		return resp, err
	}

	resp.Data = respData
	return resp, nil
}

func (s *Service) SetupProviders(providers map[string]IAuthProvider) {
	s.authProviders = providers
}
