package main

import (
	"flag"

	"github.com/linshenqi/authy/src/services/jwt"
	"github.com/linshenqi/authy/src/services/oauth"
	"github.com/linshenqi/authy/src/services/totp"
	"github.com/linshenqi/sptty"
)

func main() {
	cfg := flag.String("config", "./config.yml", "--config")
	flag.Parse()

	app := sptty.GetApp()
	app.ConfFromFile(*cfg)

	services := sptty.Services{
		&oauth.Service{},
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
