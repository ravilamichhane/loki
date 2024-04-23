package core

import (
	"nest/common"
)

type Nest struct {
	module *common.Module
	mux    Mux
}

type Mux interface {
	HandleFunc(route *common.Route)
	ListenAndServe() error
	AddRouters(routers ...common.ControllerBase)
}

type NestFactory struct {
	nest *Nest
}

func NewNestFactory(mux Mux) *NestFactory {
	return &NestFactory{
		nest: &Nest{
			mux: mux,
		},
	}
}

func (n *NestFactory) Create(module *common.Module) *Nest {
	n.nest.module = module
	return n.nest
}

func (n *Nest) Listen(port string) {
	if len(n.module.Controllers) > 0 {

		n.mux.AddRouters(n.module.Controllers...)
		n.mux.ListenAndServe()

	}
}
