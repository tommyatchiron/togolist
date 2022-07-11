package list

import (
	"net/http"
	"strconv"

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
		rg.GET("", lc.getAll)
		rg.GET("/:id", lc.getOne)
		rg.PUT("/:id", lc.update)
		rg.DELETE("/:id", lc.delete)
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

// getAll godoc
// @Summary      Get All Lists
// @Description  Get all Lists
// @Produce      json
// @Success      200  {array}  dto.List
// @Router       /lists [get]
func (lc *ListController) getAll(c *gin.Context) {
	lists, err := lc.listService.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lists)
}

// getOne godoc
// @Summary      Get One List
// @Description  Get a single List
// @Produce      json
// @Param        id   path      uint      true  "List ID"
// @Success      200  {object}  dto.List
// @Failure      404  {string}  string  "List not found"
// @Router       /lists/{id} [get]
func (lc *ListController) getOne(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, err := lc.listService.GetOne(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if list == nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, list)
}

// update godoc
// @Summary      Update List
// @Description  Update a single List
// @Accept       json
// @Produce      json
// @Param        id    path      uint                 true  "List ID"
// @Param        list  body      dto.UpdateListInput  true  "Update List"
// @Success      200   {object}  dto.List
// @Failure      404   {string}  string  "List not found"
// @Router       /lists/{id} [put]
func (lc *ListController) update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var partialUpdateListInput dto.PartialUpdateListInput
	if err := c.ShouldBindJSON(&partialUpdateListInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateListInput := dto.UpdateListInput{
		ID:                     uint(id),
		PartialUpdateListInput: partialUpdateListInput,
	}
	list, err := lc.listService.Update(c, &updateListInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if list == nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, list)
}

// delete godoc
// @Summary      Delete List
// @Description  Delete a single List
// @Produce      json
// @Param        id   path      uint  true  "List ID"
// @Success      200  {object}  dto.List  "List deleted"
// @Failure      404  {string}  string    "List not found"
// @Router       /lists/{id} [delete]
func (lc *ListController) delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, err := lc.listService.Delete(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if list == nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, list)
}
