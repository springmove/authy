package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
)

type MiniProgram struct {
	http      *resty.Client
	endpoints map[string]Endpoint
}

func (s *MiniProgram) Auth(req interface{}) (interface{}, error) {
	authData, endpoint, err := preAuth(req, s.endpoints)
	if err != nil {
		return nil, err
	}

	resp, err := s.miniprogramAuth(endpoint.AppID, endpoint.AppSecret, authData.Code)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *MiniProgram) init(endpoints map[string]Endpoint) {
	s.endpoints = endpoints
	s.http = sptty.CreateHttpClient(sptty.DefaultHttpClientConfig())
}

func (s *MiniProgram) miniprogramAuth(appID string, secret string, code string) (MiniProgramAuthResponse, error) {
	rt := MiniProgramAuthResponse{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appID=%s&secret=%s&js_code=%s&grant_type=authorization_code",
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
