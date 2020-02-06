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
	app.AddRoute("GET", "/oauth-endpoint", s.postAuth)
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

func (s *Service) doOAuth(req Request) (Response, error) {
	resp := Response{
		Type: req.Type,
	}

	provider, err := s.getProvider(req.Type)
	if err != nil {
		return resp, err
	}

	respData, err := provider.OAuth(&req)
	if err != nil {
		return resp, err
	}

	respData.Type = req.Type
	return resp, nil
}

func (s *Service) getProvider(oauthType string) (IOAuthProvider, error) {
	provider, exist := s.oauthProviders[oauthType]
	if !exist {
		return nil, errors.New("Provider Not Found ")
	}

	return provider, nil
}

func (s *Service) findEndpoint(oauthType string, endpoint string) (*Endpoint, error) {
	provider, err := s.getProvider(oauthType)
	if err != nil {
		return nil, err
	}

	return provider.GetEndpoint(endpoint)
}

func (s *Service) SetupProviders(providers map[string]IOAuthProvider) {
	s.oauthProviders = providers
}
