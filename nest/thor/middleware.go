package thor

import (
	"fmt"
	"log"
	"nest/common"
	"net/http"
	"runtime/debug"
)

func ErrorMiddleware(h common.RouteHandler) common.RouteHandler {

	return func(ctx common.HttpContext) error {
		if err := h(ctx); err != nil {
			log.Println("CORS MIDDLEWARE")

			var er ErrorResponse
			var status int
			switch {

			case IsFieldErrors(err):
				er = ErrorResponse{
					Error:  "Validation Errors",
					Errors: GetFieldErrors(err).Fields(),
				}
				status = http.StatusBadRequest

			case IsTrustedError(err):
				status = GetTrustedError(err).Status
				er = ErrorResponse{
					Error: GetTrustedError(err).Error(),
				}

			default:
				ctx.GetLogger().Error(ctx.GetContext(), "message", "msg", err)
				status = http.StatusInternalServerError
				if GetEnvBool("APP_DEBUG") {
					er = ErrorResponse{
						Error: err.Error(),
					}
				} else {
					er = ErrorResponse{
						Error: "Internal Server Error",
					}
				}
			}

			ctx.JSON(status, &er)

			if IsShutdownError(err) {
				return err
			}

		}
		return nil
	}
}

func LoggingMiddleware(h common.RouteHandler) common.RouteHandler {
	return func(ctx common.HttpContext) error {
		ctx.GetLogger().Info(
			ctx.GetContext(), "request started",
			"method", ctx.GetMethod(),
			"path", ctx.GetPath(),
			"remote_addr", ctx.GetRemoteAddr(),
		)
		err := h(ctx)

		ctx.GetLogger().Info(
			ctx.GetContext(), "request completed",
			"method", ctx.GetMethod(),
			"path", ctx.GetPath(),
			"status", ctx.GetStatusCode(),
		)
		return err
	}
}

func PanicMiddleWare(h common.RouteHandler) common.RouteHandler {
	return func(ctx common.HttpContext) (err error) {
		defer func() {

			if r := recover(); r != nil {
				// trace := debug.Stack()
				debug.PrintStack()
				err = fmt.Errorf("PANIC RECOVERED %v", r)

			} else {
			}
		}()

		return h(ctx)
	}
}

func CORSMiddleWare(handler common.RouteHandler) common.RouteHandler {
	return func(ctx common.HttpContext) error {

		err := handler(ctx)

		return err
	}
}
