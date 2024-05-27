package common

type Module struct {
	Controllers []ControllerBase
}

type ModuleConfig struct {
	Controllers []ControllerBase
}

func NewlokiModule(config ModuleConfig) *Module {
	return &Module{
		Controllers: config.Controllers,
	}
}

type Router interface {
	RegisterRoutes(
		routes []Route,
	)
}

type ControllerBase interface {
	Routes() []Route
}
