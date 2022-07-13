package dto

import (
	"github.com/tommyatchiron/togolist/internal/pkg/list/entities"
	"gorm.io/gorm"
)

type UpdateListInput struct {
	ID uint `json:"id" validate:"required" example:"1"`
	PartialUpdateListInput
}

type PartialUpdateListInput struct {
	Title       *string `json:"title" example:"My List"`
	Description *string `json:"description" example:"This is my list"`
	Priority    *int64  `json:"priority" example:"0"`
}

func (u *UpdateListInput) ToEntity() *entities.List {
	return &entities.List{
		Model: gorm.Model{
			ID: u.ID,
		},
		Title:       u.Title,
		Description: u.Description,
		Priority:    u.Priority,
	}
}
