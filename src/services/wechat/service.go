package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/linshenqi/authy/src/base"
	"gopkg.in/resty.v1"
	"time"
)

type WeCharService struct {
	core base.BaseService
	cfg  *WeChatConfig

	http *resty.Client
}

func (s *WeCharService) Init(service base.BaseService) error {
	s.core = service
	s.cfg = s.core.Config("wechat").(*WeChatConfig)

	s.initHttpClient()

	return nil
}

func (s *WeCharService) Release() {

}

func (s *WeCharService) Auth(code string) (WXAuthResponse, error) {
	war := WXAuthResponse{}
	resp, err := s.http.R().Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		s.cfg.AppID,
		s.cfg.AppSecret,
		code))

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

func (s *WeCharService) initHttpClient() {
	s.http = resty.New()
	s.http.SetRESTMode()
	s.http.SetTimeout(time.Duration(s.cfg.Timeout))
	s.http.SetContentLength(true)
	s.http.SetHeaders(s.cfg.Headers)
	s.http.
		SetRetryCount(s.cfg.MaxRetry).
		SetRetryWaitTime(time.Duration(s.cfg.PushInterval)).
		SetRetryMaxWaitTime(20 * time.Second)
}
