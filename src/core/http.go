package core

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/linshenqi/authy/src/base"
	"github.com/mitchellh/mapstructure"
)

const (
	BASE_API_ROUTE = "/api/v1"
)

type HttpConfig struct {
	Addr string
}

type Route struct {
	RouteType   string
	Method      string
	Pattern     string
	HandlerFunc context.Handler
}

type CorsConfig struct {
	AllowedOrigins   []string `yaml:"allowed-origins"`
	AllowCredentials bool     `yaml:"allow-credentials"`
	AllowedMethods   []string `yaml:"allowed-methods"`
}

type HttpService struct {
	app   *iris.Application
	party iris.Party
	base.Service
	//cors CorsConfig
	//routes []Route
}

func (http *HttpService) SetOptions() {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})

	http.party = http.app.Party(BASE_API_ROUTE, crs).AllowMethods(iris.MethodOptions)
}

func (http *HttpService) Init(service base.BaseService) error {
	cfg := HttpConfig{}
	mapstructure.Decode(service.Config("http"), &cfg)

	http.app.Run(iris.Addr(cfg.Addr), iris.WithoutInterruptHandler)

	return nil
}

func (http *HttpService) Release() {

}

func (http *HttpService) AddRoute(method string, route string, handler context.Handler) {
	http.party.Handle(method, route, handler)
}
