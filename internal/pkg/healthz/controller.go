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

func registerHealthzRoutes(hc *HealthzController, r *router.Router) {
	r.RegisterApiRoutes("/healthz", func(rg *gin.RouterGroup) {
		rg.GET("", hc.healthCheck)
	})
}

// healthCheck godoc
// @Summary      Health Checking
// @Description  Health Checking for API services
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.HealthzResult
// @Router       /healthz [get]
func (hc *HealthzController) healthCheck(c *gin.Context) {
	result, err := hc.healthzService.healthCheck()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	if result.Status == "ok" {
		c.JSON(200, result)
	} else {
		c.JSON(503, result)
	}
}
