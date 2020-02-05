package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/core/errors"
	"github.com/linshenqi/sptty"
)

const (
	ServiceName = "jwt"
	Secret      = "LijwefL(*IJWE)J@309j@#)(I#$@)(*"
)

type Service struct {
	cfg Config
}

func (s *Service) Init(app sptty.Sptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	app.AddRoute("PUT", "/jwt-signer", s.Signer)
	app.AddRoute("PUT", "/jwt-auther", s.Auther)
	app.AddRoute("PUT", "/jwt-parser", s.Parser)

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

func (s *Service) Sign(claims jwt.MapClaims) (string, error) {
	tokenString := ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(Secret))
	return tokenString, err
}

func (s *Service) Validate(myClaims jwt.MapClaims, tokenStr string) (interface{}, error) {

	claims, err := s.Parse(tokenStr)
	if err != nil {
		return nil, err
	}

	for k, _ := range claims.(jwt.MapClaims) {
		if claims.(jwt.MapClaims)[k] != myClaims[k] {
			return nil, errors.New("token validate failed")
		}
	}

	return claims, nil
}

func (s *Service) Parse(tokenStr string) (interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token validate failed")
	}
}
