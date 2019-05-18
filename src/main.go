package main

import (
	"flag"
	"github.com/linshenqi/authy/src/services/auth"
	"github.com/linshenqi/authy/src/services/wechat"
	"github.com/linshenqi/sptty"
)

type ServiceConfig struct {
	Http   sptty.HttpConfig    `yaml:"http"`
	WeChat wechat.WeChatConfig `yaml:"wechat"`
}

func main() {
	cfg := flag.String("config", "./config.yml", "--config")
	flag.Parse()

	app := sptty.GetApp()
	app.SetConf(*cfg)

	services := map[string]sptty.Service{
		"wechat": &wechat.WeChatService{},
		"auth":   &auth.AuthService{},
	}

	configs := sptty.SpttyConfig{
		"wechat": wechat.WeChatConfig{},
		"http":   sptty.HttpConfig{},
	}

	app.AddServices(services)
	app.AddConfigs(configs)

	app.Sptting()
}
