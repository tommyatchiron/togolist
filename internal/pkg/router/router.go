package router

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/tommyatchiron/togolist/internal/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.Provide(NewSwaggerHandler),
	fx.Invoke(registerSwaggerRoutes),
)

type Router struct {
	http         *gin.Engine
	routerGroups struct {
		api *gin.RouterGroup
	}
	logger *zap.SugaredLogger
}

func New(config *config.Config, logger *zap.SugaredLogger) *Router {
	var r Router
	gin.SetMode(gin.ReleaseMode)
	r.http = gin.New()
	r.http.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
	r.routerGroups.api = r.http.Group("/v1")
	r.logger = logger
	return &r
}

func (r *Router) RegisterApiRoutes(path string, f func(*gin.RouterGroup)) {
	r.logger.Infof("Registered routes %s", path)
	f(r.routerGroups.api.Group(path))
}

func (r *Router) GetHttpRouter() *gin.Engine {
	return r.http
}
