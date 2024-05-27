package common

type RouteHandler func(ctx HttpContext) error

type Route struct {
	Method      string
	Path        string
	Handler     RouteHandler
	MiddleWares []MiddleWare
}

func GET(path string, handler RouteHandler, middlewares ...MiddleWare) Route {
	return Route{
		Method:      "GET",
		Path:        path,
		Handler:     handler,
		MiddleWares: middlewares,
	}
}

func POST(path string, handler RouteHandler, middlewares ...MiddleWare) Route {
	return Route{
		Method:      "POST",
		Path:        path,
		Handler:     handler,
		MiddleWares: middlewares,
	}
}

func PUT(path string, handler RouteHandler, middlewares ...MiddleWare) Route {
	return Route{
		Method:      "PUT",
		Path:        path,
		Handler:     handler,
		MiddleWares: middlewares,
	}
}

func DELETE(path string, handler RouteHandler, middlewares ...MiddleWare) Route {
	return Route{
		Method:      "DELETE",
		Path:        path,
		Handler:     handler,
		MiddleWares: middlewares,
	}
}

func wrapMiddleware(handler RouteHandler, middleware []MiddleWare) RouteHandler {
	for i := len(middleware) - 1; i >= 0; i-- {
		middleware := middleware[i]
		if middleware != nil {
			handler = middleware(handler)
		}
	}
	return handler
}

func (r Route) AppendMiddleware(middlewares ...MiddleWare) Route {
	return Route{
		Method:  r.Method,
		Path:    r.Path,
		Handler: wrapMiddleware(r.Handler, middlewares),
	}
}

type Controller interface {
	Routes() []Route
}

type MiddleWare func(RouteHandler) RouteHandler
