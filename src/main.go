package main

import (
	"flag"
	"github.com/linshenqi/authy/src/core"
	"github.com/linshenqi/authy/src/services/auth"
	"github.com/linshenqi/authy/src/services/wechat"
)

func main() {
	cfg := flag.String("config", "./config.yml", "--config")
	flag.Parse()

	app := core.GetApp()
	app.SetConf(*cfg)
	app.RegistService("wechat", &wechat.WeChatService{})
	app.RegistService("auth", &auth.AuthService{})

	app.Init()
	app.Release()
}
