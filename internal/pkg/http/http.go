package http

import (
	"context"
	"net/http"

	"github.com/tommyatchiron/togolist/internal/pkg/config"
	"github.com/tommyatchiron/togolist/internal/pkg/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(newHttp),
	fx.Invoke(runHttpServer),
)

func newHttp(config *config.Config, router *router.Router) *http.Server {
	return &http.Server{
		Addr:    config.Http.ListenAddr,
		Handler: router.GetHttpRouter(),
	}
}

func runHttpServer(lifecycle fx.Lifecycle, httpServer *http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go httpServer.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})
}
