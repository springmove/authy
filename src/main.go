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

	oauthService := &oauth.Service{}
	oauthService.SetupProviders(map[string]oauth.IOAuthProvider{
		oauth.WeChat:            &wechat.OAuth{},
		oauth.WeChatMiniProgram: &wechat.MiniProgram{},
		oauth.AliPay:            &alipay.OAuth{},
	})

	services := sptty.Services{
		oauthService,
		&jwt.Service{},
		&totp.Service{},
	}

	configs := sptty.Configs{
		&oauth.Config{},
		&jwt.Config{},
		&totp.Config{},
	}

	app.AddServices(services)
	app.AddConfigs(configs)
	
	app.Sptting()
}
