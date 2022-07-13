package dto

import "github.com/tommyatchiron/togolist/internal/pkg/list/entities"

type CreateListInput struct {
	Title       *string `json:"title" validate:"required" example:"My List"`
	Description *string `json:"description" validate:"required" example:"This is my list"`
	Priority    *int64  `json:"priority" example:"0"`
}

func (c *CreateListInput) ToEntity() *entities.List {
	return &entities.List{
		Title:       c.Title,
		Description: c.Description,
		Priority:    c.Priority,
	}
}
