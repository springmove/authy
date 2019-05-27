package auth

import (
	"github.com/linshenqi/sptty"
)

type AuthService struct {
	app         sptty.Sptty
	controllers *AuthControllers
}

func (s *AuthService) Init(app sptty.Sptty) error {
	s.app = app
	s.controllers = &AuthControllers{
		service: s,
	}

	s.app.AddRoute("POST", "/auth", s.controllers.Auth)

	return nil
}

func (s *AuthService) Release() {

}

func (s *AuthService) Enable() bool {
	return true
}
