package dto

import "github.com/tommyatchiron/togolist/internal/pkg/list/entities"

type List struct {
	ID          uint    `json:"id" example:"1"`
	Title       *string `json:"title" example:"My List"`
	Description *string `json:"description" example:"This is my list"`
	Priority    *int    `json:"priority" example:"0"`
}

func FromEntity(list *entities.List) *List {
	return &List{
		ID:          list.ID,
		Title:       list.Title,
		Description: list.Description,
		Priority:    list.Priority,
	}
}
