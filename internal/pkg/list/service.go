package list

import (
	"context"

	"github.com/tommyatchiron/togolist/internal/pkg/list/dto"
	"gorm.io/gorm"
)

type ListService struct {
	db *gorm.DB
}

func NewListService(db *gorm.DB) *ListService {
	return &ListService{db: db}
}

func (ls *ListService) Create(ctx context.Context, input *dto.CreateListInput) (*dto.List, error) {
	list := input.ToEntity()
	res := ls.db.WithContext(ctx).Create(list)
	if res.Error != nil {
		return nil, res.Error
	}
	return dto.FromEntity(list), nil
}
