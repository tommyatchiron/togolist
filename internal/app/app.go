package app

import (
	"github.com/tommyatchiron/togolist/internal/pkg/config"
	"github.com/tommyatchiron/togolist/internal/pkg/db"
	"github.com/tommyatchiron/togolist/internal/pkg/healthz"
	"github.com/tommyatchiron/togolist/internal/pkg/http"
	"github.com/tommyatchiron/togolist/internal/pkg/logger"
	"github.com/tommyatchiron/togolist/internal/pkg/router"
	"go.uber.org/fx"
)

func New() *fx.App {
	app := fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(router.New),
		fx.Provide(db.New),
		logger.Module,
		http.Module,
		healthz.Module,
	)
	return app
}
