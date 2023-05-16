package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/springmove/authy/src/base"
)

type OAuth struct {
	base.BaseOAuth
}

func (s *OAuth) OAuth(req *base.Request) (*base.Response, error) {
	endpoint, err := s.PreAuth(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.doAuth(endpoint.AppID, endpoint.AppSecret, req.Code)
	if err != nil {
		return nil, err
	}

	user, err := s.getUserInfo(resp.AccessToken, resp.OpenID)
	if err != nil {
		return nil, err
	}

	return user.toAuthResponseData(), nil
}

func (s *OAuth) doAuth(appID string, secret string, code string) (OAuthResponse, error) {
	rt := OAuthResponse{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		appID,
		secret,
		code)

	resp, err := s.Http.R().Get(url)

	if err != nil {
		return rt, err
	} else {
		err = json.Unmarshal(resp.Body(), &rt)
		if err != nil {
			return rt, err
		} else {
			if rt.ErrCode != WxOK {
				return rt, fmt.Errorf("ErrCode: %d, ErrMsg: %s", rt.ErrCode, rt.ErrMsg)
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

	resp, err := s.Http.R().Get(url)

	if err != nil {
		return rt, err
	} else {
		err = json.Unmarshal(resp.Body(), &rt)
		if err != nil {
			return rt, err
		} else {
			if rt.ErrCode != WxOK {
				return rt, fmt.Errorf("ErrCode: %d, ErrMsg: %s", rt.ErrCode, rt.ErrMsg)
			}

			return rt, nil
		}
	}
}

func getAccessToken(req *ReqAccessToken) (*RespAccessToken, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		req.AppID,
		req.Secret)

	resp, err := req.http.R().Get(url)
	if err != nil {
		return nil, err
	}

	rt := RespAccessToken{}
	if err = json.Unmarshal(resp.Body(), &rt); err != nil {
		return nil, err
	}

	if rt.ErrCode != WxOK {
		return nil, fmt.Errorf("ErrCode: %d, ErrMsg: %s", rt.ErrCode, rt.ErrMsg)
	}

	return &rt, nil
}
