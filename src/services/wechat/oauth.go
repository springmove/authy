package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
)

type OAuth struct {
	http      *resty.Client
	endpoints map[string]Endpoint
}

func (s *OAuth) Auth(req interface{}) (interface{}, error) {
	authData, endpoint, err := preAuth(req, s.endpoints)
	if err != nil {
		return nil, err
	}

	resp, err := s.doAuth(endpoint.AppID, endpoint.AppSecret, authData.Code)
	if err != nil {
		return nil, err
	}

	user, err := s.getUserInfo(resp.AccessToken, resp.OpenID)
	if err != nil {
		return nil, err
	}

	return user.UserInfo, nil
}

func (s *OAuth) init(endpoints map[string]Endpoint) {
	s.endpoints = endpoints
	s.http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())
}

func (s *OAuth) doAuth(appID string, secret string, code string) (OAuthResponse, error) {
	rt := OAuthResponse{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		appID,
		secret,
		code)

	resp, err := s.http.R().Get(url)

	if err != nil {
		return rt, err
	} else {
		err = json.Unmarshal(resp.Body(), &rt)
		if err != nil {
			return rt, err
		} else {
			if rt.ErrCode != WxOK {
				return rt, errors.New(fmt.Sprintf("ErrCode: %d, ErrMsg: %s", rt.ErrCode, rt.ErrMsg))
			}

			return rt, nil
		}
	}
}

func (s *OAuth) getUserInfo(accessToken string, openID string) (UserInfoResponse, error) {
	rt := UserInfoResponse{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",
		accessToken,
		openID)

	resp, err := s.http.R().Get(url)

	if err != nil {
		return rt, err
	} else {
		err = json.Unmarshal(resp.Body(), &rt)
		if err != nil {
			return rt, err
		} else {
			if rt.ErrCode != WxOK {
				return rt, errors.New(fmt.Sprintf("ErrCode: %d, ErrMsg: %s", rt.ErrCode, rt.ErrMsg))
			}

			return rt, nil
		}
	}
}
