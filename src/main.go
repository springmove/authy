package main

import (
	"flag"
	"github.com/linshenqi/authy/src/services/auth"
	"github.com/linshenqi/authy/src/services/jwt"
	"github.com/linshenqi/authy/src/services/wechat"
	"github.com/linshenqi/sptty"
)

func main() {
	cfg := flag.String("config", "./config.yml", "--config")
	flag.Parse()

	app := sptty.GetApp()
	app.ConfFromFile(*cfg)

	services := map[string]sptty.Service{
		"wechat": &wechat.WeChatService{},
		"auth":   &auth.AuthService{},
		"jwt":    &jwt.JwtService{},
	}

	configs := sptty.SpttyConfig{
		"http":   sptty.HttpConfig{},
		"model":  sptty.ModelConfig{},
		"wechat": wechat.WeChatConfig{},
		"jwt":    jwt.JwtConfig{},
	}

	app.AddServices(services)
	app.AddConfigs(configs)

	app.Sptting()
}
