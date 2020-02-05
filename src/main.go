package main

import (
	"flag"
	"github.com/linshenqi/authy/src/services/alipay"
	"github.com/linshenqi/authy/src/services/jwt"
	"github.com/linshenqi/authy/src/services/oauth"
	"github.com/linshenqi/authy/src/services/totp"
	"github.com/linshenqi/authy/src/services/wechat"
	"github.com/linshenqi/sptty"
)

func main() {
	cfg := flag.String("config", "./config.yml", "--config")
	flag.Parse()

	app := sptty.GetApp()
	app.ConfFromFile(*cfg)

	wechatService := &wechat.Service{}
	alipayService := &alipay.Service{}

	authService := &oauth.Service{}
	authService.SetupProviders(map[string]oauth.IOAuthProvider{
		oauth.WeChatOAuth:       &wechatService.OAuth,
		oauth.WeChatMiniProgram: &wechatService.MiniProgram,
		oauth.AliPayOAuth:       &alipayService.OAuth,
	})

	services := sptty.Services{
		wechatService,
		alipayService,
		authService,
		&jwt.Service{},
		&totp.Service{},
	}

	configs := sptty.Configs{
		&wechat.Config{},
		&alipay.Config{},
		&jwt.Config{},
		&totp.Config{},
	}

	app.AddServices(services)
	app.AddConfigs(configs)

	app.Sptting()
}
