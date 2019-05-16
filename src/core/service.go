package core

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/linshenqi/authy/src/base"
)

var appService *AppService = nil

func GetApp() *AppService {
	if appService == nil {
		appService = &AppService{
			http: &HttpService{
				app: iris.New(),
			},
			config:   &Config{},
			services: map[string]base.Service{},
		}

		appService.Http().SetOptions()
	}

	return appService
}

type AppService struct {
	services map[string]base.Service
	http     base.Service
	config   base.Service
	base.BaseService
}

func (bs *AppService) Init() {
	if bs.config.Init(bs) != nil {
		return
	}

	for _, v := range bs.services {
		if v.Init(bs) != nil {
			return
		}
	}

	bs.http.Init(bs)
}

func (bs *AppService) Release() {
	for _, v := range bs.services {
		v.Release()
	}

	bs.http.Release()
}

func (bs *AppService) SetConf(conf string) {
	bs.config.(*Config).SetConf(conf)
}

func (bs *AppService) Config(name string) interface{} {
	config := bs.config.(*Config)
	switch name {
	case "http":
		return &config.cfg.Http
	case "wechat":
		return &config.cfg.WeChat
	}

	return bs.config.(*Config).Config()
}

func (bs *AppService) AddRoute(method string, route string, handler context.Handler) {
	http := bs.http.(*HttpService)
	http.AddRoute(method, route, handler)
}

func (bs *AppService) Http() *HttpService {
	return bs.http.(*HttpService)
}

func (bs *AppService) GetService(name string) base.Service {
	return bs.services[name]
}

func (bs *AppService) RegistService(name string, service base.Service) {
	bs.services[name] = service
}
