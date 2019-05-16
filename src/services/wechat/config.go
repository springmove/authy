package wechat

import (
	"github.com/linshenqi/authy/src/base"
)

type WeChatConfig struct {
	Enable       bool              `yaml:"enable"`
	AppID        string            `yaml:"appid"`
	AppSecret    string            `yaml:"appsecret"`
	Timeout      base.Duration     `yaml:"timeout"`
	Headers      map[string]string `yaml:"headers"`
	PushInterval base.Duration     `yaml:"push_interval"`
	MaxRetry     int               `yaml:"max_retry"`
}
