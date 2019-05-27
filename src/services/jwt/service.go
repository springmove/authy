package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/core/errors"
	"github.com/linshenqi/sptty"
)

type JwtService struct {
	app         sptty.Sptty
	controllers *JwtControllers
	cfg         JwtConfig
}

func (s *JwtService) Init(app sptty.Sptty) error {
	s.app = app
	app.GetConfig("jwt", &s.cfg)

	s.controllers = &JwtControllers{
		service: s,
	}

	s.app.AddRoute("PUT", "/jwt-signer", s.controllers.Signer)
	s.app.AddRoute("PUT", "/jwt-auther", s.controllers.Auther)

	return nil
}

func (s *JwtService) Release() {

}

func (s *JwtService) Enable() bool {
	return true
}

func (s *JwtService) Sign(claims jwt.MapClaims) (string, error) {
	tokenString := ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.cfg.Secret))
	return tokenString, err
}

func (s *JwtService) Validate(myClaims jwt.MapClaims, tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.cfg.Secret), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for k, _ := range claims {
			if claims[k] != myClaims[k] {
				return errors.New("token validate failed")
			}
		}

		return nil
	} else {
		return errors.New("token validate failed")
	}
}
