package base

import (
	"github.com/kataras/iris/context"
)

type BaseService interface {
	Init()
	Release()
	Config(name string) interface{}
	GetService(name string) Service
	AddRoute(method string, route string, handler context.Handler)
}

type Service interface {
	Init(service BaseService) error
	Release()
}
