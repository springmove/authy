package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/core/errors"
	"github.com/springmove/authy/src/base"
	jwt2 "github.com/springmove/authy/src/services/jwt"
	"github.com/springmove/authy/src/services/totp"
	"github.com/springmove/sptty"
	"gopkg.in/resty.v1"
)

type Config struct {
	Url string `yaml:"url"`
}

type Authy struct {
	cfg  *Config
	http *resty.Client
}

func (s *Authy) Init(cfg *Config) {
	s.cfg = cfg
	s.http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())
}

func (s *Authy) OAuth(req base.Request) (base.Response, error) {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/oauth", s.cfg.Url)
	resp, err := r.Post(url)

	ab := base.Response{}

	if err != nil {
		return ab, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return ab, errors.New(string(resp.Body()))
		} else {
			_ = json.Unmarshal(resp.Body(), &ab)
			return ab, nil
		}
	}
}

func (s *Authy) OAuthEndpoint(oauthType string, endpoint string) (*base.Endpoint, error) {
	r := s.http.R().SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/oauth?type=%s&endpoint=%s", s.cfg.Url, oauthType, endpoint)
	resp, err := r.Get(url)

	ab := base.Endpoint{}

	if err != nil {
		return nil, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return nil, errors.New(string(resp.Body()))
		} else {
			_ = json.Unmarshal(resp.Body(), &ab)
			return &ab, nil
		}
	}
}

func (s *Authy) JwtSigner(claims jwt.MapClaims) (string, error) {
	r := s.http.R().SetBody(claims).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/jwt-signer", s.cfg.Url)
	resp, err := r.Put(url)

	jr := jwt2.Response{}

	if err != nil {
		return jr.Token, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return jr.Token, errors.New(string(resp.Body()))
		} else {
			_ = json.Unmarshal(resp.Body(), &jr)
			return jr.Token, nil
		}
	}
}

func (s *Authy) JwtAuther(req jwt2.Request) (jwt.MapClaims, error) {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/jwt-auther", s.cfg.Url)
	resp, err := r.Put(url)

	if err != nil {
		return nil, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return nil, errors.New(string(resp.Body()))
		} else {
			claims := jwt.MapClaims{}
			_ = json.Unmarshal(resp.Body(), &claims)
			return claims, nil
		}
	}
}

func (s *Authy) JwtParser(token string) (jwt.MapClaims, error) {
	req := jwt2.Request{
		Token: token,
	}

	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/jwt-parser", s.cfg.Url)
	resp, err := r.Put(url)

	if err != nil {
		return nil, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return nil, errors.New(string(resp.Body()))
		} else {
			claims := jwt.MapClaims{}
			_ = json.Unmarshal(resp.Body(), &claims)
			return claims, nil
		}
	}
}

func (s *Authy) TotpGenerate(endpoint string) (totp.GenerateBody, error) {
	req := totp.RequestEndpoint{
		Endpoint: endpoint,
	}

	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/totp-generate", s.cfg.Url)
	resp, err := r.Put(url)

	body := totp.GenerateBody{}
	if err != nil {
		return body, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return body, errors.New(string(resp.Body()))
		} else {
			_ = json.Unmarshal(resp.Body(), &body)
			return body, nil
		}
	}
}

func (s *Authy) TotpValidate(req totp.ValidateRequest) (bool, error) {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/totp-validate", s.cfg.Url)
	resp, err := r.Put(url)

	if err != nil {
		return false, err
	} else {
		if resp.StatusCode() == http.StatusOK {
			return true, nil
		} else if resp.StatusCode() == http.StatusConflict {
			return false, nil
		} else {
			return false, errors.New(string(resp.Body()))
		}
	}
}
