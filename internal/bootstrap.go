package internal

import (
	"context"
	"fmt"
	"testing"

	"github.com/ankorstore/yokai/fxcore"
	"github.com/ankorstore/yokai/fxhttpserver"
	"go.uber.org/fx"
)

// Bootstrapper can be used to load modules, options, dependencies, routing and bootstraps the application.
func newBootstrapper() *fxcore.Bootstrapper {
	return fxcore.NewBootstrapper().WithOptions(
		// modules registration
		fxhttpserver.FxHttpServerModule,
		// dependencies registration
		Register(),
		// routing registration
		Router(),
	)
}

// Run starts the application, with a provided [context.Context].
func Run(ctx context.Context) {
	newBootstrapper().WithContext(ctx).RunApp()
}

// RunTest starts the application in test mode, with an optional list of [fx.Option].
func RunTest(tb testing.TB, options ...fx.Option) {
	tb.Helper()

	// RootDir is the application root directory.
	RootDir := fxcore.RootDir(1)

	tb.Setenv("APP_CONFIG_PATH", fmt.Sprintf("%s/configs", RootDir))

	newBootstrapper().RunTestApp(tb, fx.Options(options...))
}
