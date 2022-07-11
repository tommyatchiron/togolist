package app

import (
	"github.com/tommyatchiron/togolist/internal/pkg/config"
	"github.com/tommyatchiron/togolist/internal/pkg/db"
	"github.com/tommyatchiron/togolist/internal/pkg/healthz"
	"github.com/tommyatchiron/togolist/internal/pkg/http"
	"github.com/tommyatchiron/togolist/internal/pkg/list"
	"github.com/tommyatchiron/togolist/internal/pkg/logger"
	"github.com/tommyatchiron/togolist/internal/pkg/router"
	"go.uber.org/fx"
)

func New() *fx.App {
	app := fx.New(
		fx.Provide(config.NewConfig),
		fx.Provide(db.New),
		logger.Module,
		http.Module,
		router.Module,
		healthz.Module,
		list.Module,
	)
	return app
}
