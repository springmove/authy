package auth

import (
	"github.com/linshenqi/authy/src/base"
)

type AuthService struct {
	core        base.BaseService
	controllers *AuthControllers
}

func (s *AuthService) Init(service base.BaseService) error {
	s.core = service
	s.controllers = &AuthControllers{
		service: s,
	}

	s.core.AddRoute("POST", "/auth", s.controllers.Auth)

	return nil
}

func (s *AuthService) Release() {

}
