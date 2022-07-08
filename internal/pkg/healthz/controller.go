package healthz

import (
	"github.com/gin-gonic/gin"
	"github.com/tommyatchiron/togolist/internal/pkg/router"
)

type HealthzController struct {
	healthzService *HealthzService
}

func NewHealthzController(healthzService *HealthzService) *HealthzController {
	return &HealthzController{healthzService}
}

func registerHealthzRoutes(hc *HealthzController, r *router.Router) *router.Router {
	r.RegisterApiRoutes("/healthz", func(rg *gin.RouterGroup) {
		rg.GET("", hc.healthCheck)
	})
	return r
}

func (hc *HealthzController) healthCheck(c *gin.Context) {
	result, err := hc.healthzService.healthCheck()
	if err != nil {

	}
	if result.Status == "OK" {
		c.JSON(200, result)
	} else {
		c.JSON(503, result)
	}
}
