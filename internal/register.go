package internal

import (
	"github.com/ankorstore/yokai/config"
	"github.com/labstack/echo/v4"
	"github.com/ndk/test123/internal/handler"
	"github.com/ndk/test123/spec"
	"go.uber.org/fx"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		fx.Invoke(func(e *echo.Echo) {
			a := handler.NewAPI(&config.Config{})
			strictApiHandler := spec.NewStrictHandler(a, nil)
			spec.RegisterHandlers(e, strictApiHandler)
		}),
	)
}
