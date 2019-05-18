package wechat

type WeChatConfig struct {
	Enable       bool              `yaml:"enable"`
	AppID        string            `yaml:"appid"`
	AppSecret    string            `yaml:"appsecret"`
	Timeout      int               `yaml:"timeout"`
	Headers      map[string]string `yaml:"headers"`
	PushInterval int               `yaml:"push_interval"`
	MaxRetry     int               `yaml:"max_retry"`
}
