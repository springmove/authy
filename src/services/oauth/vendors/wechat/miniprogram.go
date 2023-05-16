package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/springmove/authy/src/base"
)

type MiniProgram struct {
	base.BaseOAuth
}

func (s *MiniProgram) OAuth(req *base.Request) (*base.Response, error) {
	endpoint, err := s.PreAuth(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.miniprogramAuth(endpoint.AppID, endpoint.AppSecret, req.Code)
	if err != nil {
		return nil, err
	}

	respOAuth := resp.toAuthResponseData()
	mobile, err := s.getMobileByAuthCode(endpoint.AppID, endpoint.AppSecret, req.AuthCodeMobile)
	if err != nil {
		return nil, err
	}

	respOAuth.Mobile = mobile

	return respOAuth, nil
}

func (s *MiniProgram) miniprogramAuth(appID string, secret string, code string) (MiniProgramAuthResponse, error) {
	rt := MiniProgramAuthResponse{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appID=%s&secret=%s&js_code=%s&grant_type=authorization_code",
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

func (s *MiniProgram) getMobileByAuthCode(appID string, secret string, authCodeMobile string) (string, error) {
	if authCodeMobile == "" {
		return "", nil
	}

	accessToken, err := getAccessToken(&ReqAccessToken{
		http:   s.Http,
		AppID:  appID,
		Secret: secret,
	})

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", accessToken.Token)
	resp, err := s.Http.R().SetBody(&ReqMiniProgramAuthMobile{
		Code: authCodeMobile,
	}).Post(url)
	if err != nil {
		return "", err
	}

	respBody := RespMiniProgramAuthMobile{}
	if err = json.Unmarshal(resp.Body(), &respBody); err != nil {
		return "", err
	}

	if respBody.ErrCode != WxOK {
		return "", fmt.Errorf("ErrCode: %d, ErrMsg: %s", respBody.ErrCode, respBody.ErrMsg)
	}

	return respBody.Mobile.ToValidMobile(), nil
}
