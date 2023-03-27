package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/core/errors"
	"github.com/springmove/authy/src/base"
	"github.com/springmove/sptty"
)

type Service struct {
	sptty.BaseService

	cfg Config
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	app.AddRoute("PUT", "/jwt-signer", s.Signer)
	app.AddRoute("PUT", "/jwt-auther", s.Auther)
	app.AddRoute("PUT", "/jwt-parser", s.Parser)

	return nil
}

func (s *Service) ServiceName() string {
	return base.ServiceJwt
}

func (s *Service) Sign(claims jwt.MapClaims) (string, error) {
	claims[base.TimeStamp] = time.Now().Format(time.RFC3339)
	tokenString := ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(base.Secret))
	return tokenString, err
}

func (s *Service) Validate(myClaims jwt.MapClaims, tokenStr string) (jwt.MapClaims, error) {

	claims, err := s.Parse(tokenStr)
	if err != nil {
		return nil, err
	}

	for k := range claims {
		if k == base.TimeStamp {
			if s.cfg.Expiry == (0 * time.Second) {
				continue
			}

			ts, _ := time.Parse(time.RFC3339, claims[k].(string))
			if time.Now().After(ts.Add(s.cfg.Expiry)) {
				return nil, fmt.Errorf(base.CodeExpiry)
			}

		} else {
			if claims[k] != myClaims[k] {
				return nil, fmt.Errorf(base.CodeFailed)
			}
		}
	}

	return claims, nil
}

func (s *Service) Parse(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected Signing Method: %v", token.Header["alg"]))
		}

		return []byte(base.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Token Invalid")
	}
}

func (s *Service) Refresh(tokenStr string) (string, error) {
	claims, err := s.Parse(tokenStr)
	if err != nil {
		return "", err
	}

	return s.Sign(claims)
}

func (s *Service) SetSecret(secret string) {
	base.Secret = secret
}
