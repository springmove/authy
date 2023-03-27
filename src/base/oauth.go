package base

import (
	"errors"

	"github.com/springmove/sptty"
	"gopkg.in/resty.v1"
)

const (
	ServiceOAuth = "oauth"
)

type IServiceOAuth interface {
	OAuth(req Request) (Response, error)
}

type Endpoint struct {
	Provider  string `yaml:"type" json:"provider"`
	AppID     string `yaml:"appid" json:"appid"`
	AppSecret string `yaml:"secret" json:"secret"`
	PublicKey string `yaml:"publicKey" json:"publicKey"`
}

const (
	ErrAuthFailed       = "ErrAuthFailed"
	ErrEndpointNotFound = "EndpointNotFound"
)

const (
	WeChat            = "wechat"
	WeChatMiniProgram = "wechat_miniprogram"
	AliPay            = "alipay"
)

type Request struct {
	Provider string `json:"provider"`
	Endpoint string `json:"endpoint"`
	Code     string `json:"code"`
}

type Response struct {
	Type string `json:"type"`

	OpenID  string `json:"open_id"`
	UnionID string `json:"union_id"`
	Name    string `json:"name"`

	// 0:未知 1:男 2:女
	Gender int `json:"gender"`

	Province string `json:"province"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Avatar   string `json:"avatar"`
}

type IOAuthProvider interface {
	OAuth(req *Request) (*Response, error)
	Init()
	GetEndpoint(name string) (*Endpoint, error)
	AddEndpoint(name string, endpoint Endpoint)
}

type BaseOAuth struct {
	Http      *resty.Client
	Endpoints map[string]Endpoint
}

func (s *BaseOAuth) GetEndpoint(name string) (*Endpoint, error) {
	ep, exist := s.Endpoints[name]
	if !exist {
		return nil, errors.New("Endpoint Not Found ")
	}

	return &ep, nil
}

func (s *BaseOAuth) Init() {
	s.Http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())
}

func (s *BaseOAuth) AddEndpoint(name string, endpoint Endpoint) {
	if s.Endpoints == nil {
		s.Endpoints = map[string]Endpoint{}
	}
	s.Endpoints[name] = endpoint
}

func (s *BaseOAuth) PreAuth(req *Request) (*Endpoint, error) {
	if req == nil {
		return nil, errors.New("Request Data Is Nil ")
	}

	endpoint, err := s.GetEndpoint(req.Endpoint)
	if err != nil {
		return nil, err
	}

	return endpoint, nil
}
