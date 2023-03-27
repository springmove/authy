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

	return resp.toAuthResponseData(), nil
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
