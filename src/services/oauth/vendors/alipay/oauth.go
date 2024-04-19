package alipay

import (
	"errors"
	"fmt"

	v3 "github.com/smartwalle/alipay/v3"
	"github.com/springmove/authy/src/base"
	"github.com/springmove/sptty"
)

type OAuth struct {
	base.BaseOAuth
	clients map[string]*v3.Client
}

func (s *OAuth) Init() {
	s.clients = map[string]*v3.Client{}

	for k, v := range s.Endpoints {
		client, err := v3.New(v.AppID, v.AppSecret, true)
		if err != nil {
			sptty.Log(sptty.ErrorLevel, fmt.Sprintf("Create Alipay Client Failed: %s", err.Error()), base.AliPay)
			continue
		}

		if err := client.LoadAliPayPublicKey(v.PublicKey); err != nil {
			sptty.Log(sptty.ErrorLevel, fmt.Sprintf("Load PublicKey Failed: %s", err.Error()), base.AliPay)
			continue
		}

		s.clients[k] = client
	}
}

func (s *OAuth) OAuth(req *base.Request) (*base.Response, error) {
	if req == nil {
		return nil, errors.New("Request Data Is Nil ")
	}

	client, exist := s.clients[req.Endpoint]
	if !exist {
		return nil, errors.New("Client Not Found ")
	}

	authResp, err := client.SystemOauthToken(v3.SystemOauthToken{
		Code:      req.Code,
		GrantType: "authorization_code",
	})

	if err != nil {
		return nil, err
	}

	user, err := client.UserInfoShare(v3.UserInfoShare{
		AppAuthToken: "",
		AuthToken:    authResp.AccessToken,
	})

	if err != nil {
		return nil, err
	}

	resp := UserInfoResponse{
		UserInfoShareRsp: *user,
	}

	return resp.toAuthResponseData(), nil
}
