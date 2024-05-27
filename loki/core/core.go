package core

import (
	"loki/common"
)

type loki struct {
	module *common.Module
	mux    Mux
}

type Mux interface {
	HandleFunc(route *common.Route)
	ListenAndServe() error
	AddRouters(routers ...common.ControllerBase)
}

type lokiFactory struct {
	loki *loki
}

func NewlokiFactory(mux Mux) *lokiFactory {
	return &lokiFactory{
		loki: &loki{
			mux: mux,
		},
	}
}

func (n *lokiFactory) Create(module *common.Module) *loki {
	n.loki.module = module
	return n.loki
}

func (n *loki) Listen(port string) {
	if len(n.module.Controllers) > 0 {

		n.mux.AddRouters(n.module.Controllers...)
		n.mux.ListenAndServe()

	}
}
