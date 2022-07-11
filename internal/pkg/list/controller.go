package list

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tommyatchiron/togolist/internal/pkg/list/dto"
	"github.com/tommyatchiron/togolist/internal/pkg/router"
)

type ListController struct {
	listService *ListService
}

func NewListController(listService *ListService) *ListController {
	return &ListController{listService: listService}
}

func registerListRoutes(lc *ListController, r *router.Router) {
	r.RegisterApiRoutes("/lists", func(rg *gin.RouterGroup) {
		rg.POST("", lc.create)
	})
}

// create godoc
// @Summary      Create List
// @Description  Create a single List
// @Accept       json
// @Produce      json
// @Param        list  body      dto.CreateListInput  true  "Create List"
// @Success      201   {object}  dto.List
// @Router       /lists [post]
func (lc *ListController) create(c *gin.Context) {
	var createListInput dto.CreateListInput
	if err := c.ShouldBindJSON(&createListInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, err := lc.listService.Create(c, &createListInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, list)
}
