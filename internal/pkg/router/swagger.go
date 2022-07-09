package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tommyatchiron/togolist/docs"
)

// gin-swagger middleware
// swagger embed files

// @title        ToGo List API
// @version      1.0
// @description  This is an API for ToGo List application

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

type SwaggerHandler struct {
	handler gin.HandlerFunc
}

func NewSwaggerHandler() *SwaggerHandler {
	return &SwaggerHandler{ginSwagger.WrapHandler(swaggerFiles.Handler)}
}

func registerSwaggerRoutes(sh *SwaggerHandler, r *Router) {
	r.RegisterApiRoutes("/docs", func(rg *gin.RouterGroup) {
		rg.GET("/*any", sh.handler)
	})
}
