package dto

import "github.com/tommyatchiron/togolist/internal/pkg/list/entities"

type CreateListInput struct {
	Title       *string `json:"title" validate:"required"`
	Description *string `json:"description" validate:"required"`
	Priority    *int    `json:"priority"`
}

func (c *CreateListInput) ToEntity() *entities.List {
	return &entities.List{
		Title:       c.Title,
		Description: c.Description,
		Priority:    c.Priority,
	}
}
