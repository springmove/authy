package oauth

import (
	"errors"
	"github.com/linshenqi/sptty"
)

const (
	ServiceName = "oauth"
)

type Service struct {
	oauthProviders map[string]IOAuthProvider
}

func (s *Service) Init(app sptty.Sptty) error {
	app.AddRoute("POST", "/oauth", s.postAuth)
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

	provider, exist := s.oauthProviders[req.Type]
	if !exist {
		return resp, errors.New("Provider Not Found ")
	}

	respData, err := provider.OAuth(&req)
	if err != nil {
		return resp, err
	}

	respData.Type = req.Type
	return resp, nil
}

func (s *Service) SetupProviders(providers map[string]IOAuthProvider) {
	s.oauthProviders = providers
}
