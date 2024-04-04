package internal

import (
	"github.com/ankorstore/yokai/config"
	"github.com/labstack/echo/v4"
	middleware "github.com/oapi-codegen/echo-middleware"
	"go.uber.org/fx"

	"github.com/ndk/test123/internal/handler"
	"github.com/ndk/test123/spec"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		fx.Invoke(func(e *echo.Echo) {
			swagger, err := spec.GetSwagger()
			if err != nil {
				//No proper error handling yet
				panic(err)
			}

			e.Use(middleware.OapiRequestValidator(swagger))

			a := handler.NewAPI(&config.Config{})
			strictApiHandler := spec.NewStrictHandler(a, nil)
			spec.RegisterHandlers(e, strictApiHandler)
		}),
	)
}
