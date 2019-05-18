package api

import (
	"encoding/json"
	"fmt"
	"github.com/linshenqi/authy/src/services/auth"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
	"net/http"
)

type AuthyConfig struct {
	Url          string            `yaml:"timeout"`
	Timeout      int               `yaml:"timeout"`
	Headers      map[string]string `yaml:"headers"`
	PushInterval int               `yaml:"push_interval"`
	MaxRetry     int               `yaml:"max_retry"`
}

type Authy struct {
	cfg  *AuthyConfig
	http *resty.Client
}

func (s *Authy) InitService(cfg *AuthyConfig) error {

	s.cfg = cfg

	clientCfg := sptty.HttpClientConfig{
		Timeout:      s.cfg.Timeout,
		Headers:      s.cfg.Headers,
		PushInterval: s.cfg.PushInterval,
		MaxRetry:     s.cfg.MaxRetry,
	}

	s.http = sptty.CreateHttpClient(&clientCfg)

	return nil
}

func (s *Authy) Auth(req *auth.AuthBody) (auth.AuthBody, error) {
	r := s.http.R().SetBody(req).SetHeader("content-type", "application/json")
	url := fmt.Sprintf("%s/api/v1/auth", s.cfg.Url)
	resp, err := r.Post(url)

	ab := auth.AuthBody{}

	if err != nil {
		return ab, err
	} else {
		if resp.StatusCode() != http.StatusOK {
			return ab, err
		} else {
			json.Unmarshal(resp.Body(), &ab)
			return ab, nil
		}
	}
}
