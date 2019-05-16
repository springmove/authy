package api

import (
	"encoding/json"
	"fmt"
	"github.com/linshenqi/authy/src/base"
	"github.com/linshenqi/authy/src/services/auth"
	"gopkg.in/resty.v1"
	"net/http"
	"time"
)

type AuthyConfig struct {
	Url          string            `yaml:"timeout"`
	Timeout      base.Duration     `yaml:"timeout"`
	Headers      map[string]string `yaml:"headers"`
	PushInterval base.Duration     `yaml:"push_interval"`
	MaxRetry     int               `yaml:"max_retry"`
}

type Authy struct {
	cfg  *AuthyConfig
	http *resty.Client
}

func (s *Authy) InitService(cfg *AuthyConfig) error {

	s.cfg = cfg
	s.http = resty.New()
	s.http.SetRESTMode()
	s.http.SetTimeout(time.Duration(cfg.Timeout))
	s.http.SetContentLength(true)
	s.http.SetHeaders(cfg.Headers)
	s.http.
		SetRetryCount(cfg.MaxRetry).
		SetRetryWaitTime(time.Duration(cfg.PushInterval)).
		SetRetryMaxWaitTime(20 * time.Second)

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
