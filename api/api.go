package api

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/core/errors"
	"github.com/linshenqi/authy/src/services/auth"
	jwt2 "github.com/linshenqi/authy/src/services/jwt"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
	"net/http"
)

type Config struct {
	Url string `yaml:"url"`
}

type Authy struct {
	cfg  *Config
	http *resty.Client
}

func (s *Authy) Init(cfg *Config) error {
	s.cfg = cfg
	s.http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())

	return nil
}

func (s *Authy) Auth(req auth.Request) (auth.Response, error) {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/auth", s.cfg.Url)
	resp, err := r.Post(url)

	ab := auth.Response{}

	if err != nil {
		return ab, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return ab, errors.New(resp.Status())
		} else {
			_ = json.Unmarshal(resp.Body(), &ab)
			return ab, nil
		}
	}
}

func (s *Authy) JwtSigner(claims *jwt.MapClaims) (string, error) {
	r := s.http.R().SetBody(claims).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/jwt-signer", s.cfg.Url)
	resp, err := r.Put(url)

	jr := jwt2.JWTResponse{}

	if err != nil {
		return jr.Token, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return jr.Token, errors.New(resp.Status())
		} else {
			json.Unmarshal(resp.Body(), &jr)
			return jr.Token, nil
		}
	}
}

func (s *Authy) JwtAuther(req *jwt2.JwtAutherRequest) (jwt.MapClaims, error) {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/jwt-auther", s.cfg.Url)
	resp, err := r.Put(url)

	if err != nil {
		return nil, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return nil, errors.New(resp.Status())
		} else {
			claims := jwt.MapClaims{}
			_ = json.Unmarshal(resp.Body(), &claims)
			return claims, nil
		}
	}
}

func (s *Authy) JwtParser(req *jwt2.JwtAutherRequest) (jwt.MapClaims, error) {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/jwt-parser", s.cfg.Url)
	resp, err := r.Put(url)

	if err != nil {
		return nil, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return nil, errors.New(resp.Status())
		} else {
			claims := jwt.MapClaims{}
			_ = json.Unmarshal(resp.Body(), &claims)
			return claims, nil
		}
	}
}
