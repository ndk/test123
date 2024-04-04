package internal

import (
	"github.com/ankorstore/yokai/fxhttpserver"
	middleware "github.com/oapi-codegen/echo-middleware"
	"go.uber.org/fx"

	"github.com/ndk/test123/spec"
)

// Router is used to register the application HTTP middlewares and handlers.
func Router() fx.Option {
	swagger, err := spec.GetSwagger()
	if err != nil {
		//No proper error handling yet
		panic(err)
	}

	return fx.Options(fxhttpserver.AsMiddleware(middleware.OapiRequestValidator(swagger), fxhttpserver.GlobalUse))
}
