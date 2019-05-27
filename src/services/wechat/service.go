package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/linshenqi/sptty"
	"gopkg.in/resty.v1"
)

type WeChatService struct {
	app sptty.Sptty
	cfg WeChatConfig

	http *resty.Client
}

func (s *WeChatService) Init(service sptty.Sptty) error {
	s.app = service
	s.app.GetConfig("wechat", &s.cfg)

	clientCfg := sptty.HttpClientConfig{
		Timeout:      s.cfg.Timeout,
		Headers:      s.cfg.Headers,
		PushInterval: s.cfg.PushInterval,
		MaxRetry:     s.cfg.MaxRetry,
	}

	s.http = sptty.CreateHttpClient(&clientCfg)

	return nil
}

func (s *WeChatService) Release() {

}

func (s *WeChatService) Enable() bool {
	return true
}

func (s *WeChatService) Auth(code string) (WXAuthResponse, error) {
	war := WXAuthResponse{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		s.cfg.AppID,
		s.cfg.AppSecret,
		code)

	resp, err := s.http.R().Get(url)

	if err != nil {
		return war, err
	} else {
		err = json.Unmarshal(resp.Body(), &war)
		if err != nil {
			return war, err
		} else {
			return war, nil
		}
	}
}
