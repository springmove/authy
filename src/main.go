package main

import (
	"flag"
	"github.com/linshenqi/authy/src/services/auth"
	"github.com/linshenqi/authy/src/services/jwt"
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

	authService := &auth.Service{}
	authService.SetupProviders(map[string]auth.IAuthProvider{
		auth.WeChatOAuth:       &wechatService.OAuth,
		auth.WeChatMiniProgram: &wechatService.MiniProgram,
	})

	services := sptty.Services{
		wechatService,
		authService,
		&jwt.Service{},
		&totp.Service{},
	}

	configs := sptty.Configs{
		&wechat.Config{},
		&jwt.Config{},
		&totp.Config{},
	}

	app.AddServices(services)
	app.AddConfigs(configs)

	app.Sptting()
}
