package thor

import (
	"context"
	"expvar"
	"fmt"
	"log"
	"nest/common"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arl/statsviz"
	"github.com/gorilla/mux"
)

// Thor is a struct that holds the mux
type Thor struct {
	shutdown  chan os.Signal
	ctx       context.Context
	Logger    common.Logger
	mux       *mux.Router
	port      string
	debugPort string
	debug     bool
	mws       []common.MiddleWare
}

type ThorConfig struct {
	Logger    common.Logger
	Port      string
	DebugPort string
	Debug     bool
	Mws       []common.MiddleWare
}

func (t *Thor) SignalShutdown() {
	t.shutdown <- syscall.SIGTERM
}

func enableCORS(next http.Handler) http.Handler {
	log.Println("CORS MIDDLEWARE")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("CORS MIDDLEWARE")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}

func NewApp(ctx context.Context, config ThorConfig) *Thor {

	if config.Logger == nil {
		panic("Logger is required")
	}

	if config.Port == "" {
		config.Port = ":8080"
	}

	if config.DebugPort == "" {
		config.DebugPort = ":8081"
	}

	newMux := mux.NewRouter().StrictSlash(true)
	newMux.Use(enableCORS)

	shutdown := make(chan os.Signal, 1)

	return &Thor{
		ctx:       ctx,
		Logger:    config.Logger,
		shutdown:  shutdown,
		mux:       newMux,
		port:      config.Port,
		debugPort: config.DebugPort,
		debug:     config.Debug,
		mws:       config.Mws,
	}
}

type routerWithMiddleware interface {
	Middlewares() []common.MiddleWare
}

type routerWithPrefix interface {
	Prefix() string
}

type routerWithMiddlewareAndPrefix interface {
	Prefix() string
	Middlewares() []common.MiddleWare
}

func (t *Thor) AddRouters(routers ...common.ControllerBase) {
	for _, router := range routers {

		rmwp, isRouterWithMiddlewareAndPrefix := router.(routerWithMiddlewareAndPrefix)

		rmw, isRouterWithMiddleware := router.(routerWithMiddleware)
		rp, isRouterWithPrefix := router.(routerWithPrefix)

		if isRouterWithMiddlewareAndPrefix {

			newRoutes := []common.Route{}
			for _, route := range router.Routes() {
				newRoute := route.AppendMiddleware(rmwp.Middlewares()...)
				newRoutes = append(newRoutes, newRoute)
			}

			t.SubRouter(rp.Prefix()).AddRoutes(newRoutes)

			continue
		}

		if isRouterWithPrefix {
			subRouter := t.SubRouter(rp.Prefix())
			subRouter.AddRoutes(router.Routes())
			continue
		}

		if isRouterWithMiddleware {
			newRoutes := []common.Route{}

			for _, route := range router.Routes() {
				newRoute := route.AppendMiddleware(rmw.Middlewares()...)
				newRoutes = append(newRoutes, newRoute)
			}
			t.AddRoutes(newRoutes)

			continue
		}

		t.AddRoutes(router.Routes())
	}
}

func (t *Thor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	log.Println("CORS MIDDLEWARE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Max-Age", "86400")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	t.mux.ServeHTTP(w, r.WithContext(ctx))
}

func (t *Thor) AddRoutes(routes []common.Route) {
	for _, router := range routes {
		t.HandleFunc(&common.Route{
			Method:  router.Method,
			Path:    router.Path,
			Handler: router.Handler,
		})
	}
}

func (t *Thor) HandleFunc(
	route *common.Route,
) {

	route.AppendMiddleware(route.MiddleWares...)
	route.AppendMiddleware(t.mws...)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := newContext(w, r, t.Logger)
		err := route.Handler(ctx)

		if err != nil {
			if validateShutdownError(err) {
				t.SignalShutdown()
				return
			}
			t.Logger.Error(ctx.GetContext(), "error", err)
			ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		}
	}

	t.mux.HandleFunc(route.Path, h).Methods(route.Method)
}

func (t *Thor) Get(path string, handler common.RouteHandler, middlewares ...common.MiddleWare) {
	t.HandleFunc(&common.Route{
		Path:        path,
		MiddleWares: middlewares,
		Method:      http.MethodGet,
	})
}

func (t *Thor) Post(path string, handler common.RouteHandler, middlewares ...common.MiddleWare) {
	t.HandleFunc(&common.Route{
		Path:        path,
		MiddleWares: middlewares,
		Method:      http.MethodPost,
	})
}

func (t *Thor) Put(path string, handler common.RouteHandler, middlewares ...common.MiddleWare) {
	t.HandleFunc(&common.Route{
		Path:        path,
		MiddleWares: middlewares,
		Method:      http.MethodPut,
	})
}

func (t *Thor) Delete(path string, handler common.RouteHandler, middlewares ...common.MiddleWare) {
	t.HandleFunc(&common.Route{
		Path:        path,
		MiddleWares: middlewares,
		Method:      http.MethodDelete,
	})
}

func (t *Thor) Patch(path string, handler common.RouteHandler, middlewares ...common.MiddleWare) {
	t.HandleFunc(&common.Route{
		Path:        path,
		MiddleWares: middlewares,
		Method:      http.MethodPatch,
	})
}

func (t *Thor) SubRouter(path string, mws ...common.MiddleWare) *Thor {
	mws = append(t.mws, mws...)
	subMux := t.mux.PathPrefix(path).Subrouter()
	thor := &Thor{
		Logger: t.Logger,
		mux:    subMux,
		mws:    mws,
	}
	return thor
}

func (t *Thor) ListenAndServe() error {

	signal.Notify(t.shutdown, syscall.SIGINT, syscall.SIGTERM)
	serverErrors := make(chan error, 1)
	if t.debug {
		go func() {
			build := os.Getenv("ENVIRONMENT")

			if build == "" {
				build = "development"
			}
			mux := http.NewServeMux()

			expvar.NewString("build").Set("development")

			mux.HandleFunc("/debug/pprof/", pprof.Index)
			mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
			mux.Handle("/debug/vars", expvar.Handler())

			statsviz.Register(mux)

			t.Logger.Info(context.Background(), fmt.Sprintf("Debug server is available at http://localhost%s", t.debugPort))

			serverErrors <- http.ListenAndServe(t.debugPort, mux)
		}()
	}

	app := &http.Server{
		Addr:    t.port,
		Handler: t,
	}

	go func() {
		t.Logger.Info(t.ctx, "startup", "status", "starting server")
		t.Logger.Info(context.Background(), fmt.Sprintf("Listening on port %s", t.port))
		t.Logger.Info(context.Background(), fmt.Sprintf("App is available at http://localhost%s", t.port))
		serverErrors <- app.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-t.shutdown:
		t.Logger.Info(t.ctx, "shutdown", "status", "shutdown started", "signal", sig)
		defer t.Logger.Info(t.ctx, "shutdown", "status", "shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(t.ctx, time.Second*20)
		defer cancel()

		if err := app.Shutdown(ctx); err != nil {
			app.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
