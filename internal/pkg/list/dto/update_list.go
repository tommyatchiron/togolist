package dto

type UpdateListInput struct {
	ID uint `json:"id" validate:"required"`
	CreateListInput
}
