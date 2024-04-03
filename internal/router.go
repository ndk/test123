package internal

import (
	"net/http"

	"github.com/ankorstore/yokai/config"
	"github.com/ankorstore/yokai/fxhttpserver"
	"github.com/labstack/echo/v4"
	middleware "github.com/oapi-codegen/echo-middleware"
	"go.uber.org/fx"

	"github.com/ndk/test123/internal/handler"
	"github.com/ndk/test123/spec"
)

type echoOptions []fx.Option

func (e *echoOptions) CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	*e = append(*e, fxhttpserver.AsHandler(http.MethodConnect, path, h))
	return nil
}

func (e *echoOptions) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	*e = append(*e, fxhttpserver.AsHandler(http.MethodDelete, path, h))
	return nil
}

func (e *echoOptions) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	*e = append(*e, fxhttpserver.AsHandler(http.MethodGet, path, h))
	return nil
}

func (e *echoOptions) HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	*e = append(*e, fxhttpserver.AsHandler(http.MethodHead, path, h))
	return nil
}

func (e *echoOptions) OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	*e = append(*e, fxhttpserver.AsHandler(http.MethodOptions, path, h))
	return nil
}

func (e *echoOptions) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	*e = append(*e, fxhttpserver.AsHandler(http.MethodPatch, path, h))
	return nil
}

func (e *echoOptions) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	*e = append(*e, fxhttpserver.AsHandler(http.MethodPost, path, h))
	return nil
}

func (e *echoOptions) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	*e = append(*e, fxhttpserver.AsHandler(http.MethodPut, path, h))
	return nil
}

func (e *echoOptions) TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	*e = append(*e, fxhttpserver.AsHandler(http.MethodTrace, path, h))
	return nil
}

// Router is used to register the application HTTP middlewares and handlers.
func Router() fx.Option {
	a := handler.NewAPI(&config.Config{})
	strictApiHandler := spec.NewStrictHandler(a, nil)
	swagger, err := spec.GetSwagger()
	if err != nil {
		//No proper error handling yet
		panic(err)
	}
	var opts echoOptions = []fx.Option{fxhttpserver.AsMiddleware(middleware.OapiRequestValidator(swagger), fxhttpserver.GlobalUse)}
	spec.RegisterHandlers(&opts, strictApiHandler)

	return fx.Options(opts...)
}
