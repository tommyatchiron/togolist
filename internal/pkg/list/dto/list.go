package dto

import "github.com/tommyatchiron/togolist/internal/pkg/list/entities"

type List struct {
	ID          uint    `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Priority    *int    `json:"priority"`
}

func FromEntity(list *entities.List) *List {
	return &List{
		ID:          list.ID,
		Title:       list.Title,
		Description: list.Description,
		Priority:    list.Priority,
	}
}
