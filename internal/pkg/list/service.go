package list

import (
	"context"
	"errors"

	"github.com/tommyatchiron/togolist/internal/pkg/list/dto"
	"github.com/tommyatchiron/togolist/internal/pkg/list/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (ls *ListService) GetAll(ctx context.Context) ([]*dto.List, error) {
	var lists []*entities.List
	res := ls.db.WithContext(ctx).Find(&lists)
	if res.Error != nil {
		return nil, res.Error
	}
	var listsDto = make([]*dto.List, len(lists))
	for i, list := range lists {
		listsDto[i] = dto.FromEntity(list)
	}
	return listsDto, nil
}

func (ls *ListService) GetOne(ctx context.Context, id uint) (*dto.List, error) {
	var list entities.List
	res := ls.db.WithContext(ctx).First(&list, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return dto.FromEntity(&list), nil
}

func (ls *ListService) Update(ctx context.Context, input *dto.UpdateListInput) (*dto.List, error) {
	list := input.ToEntity()
	updatedList := &entities.List{Model: gorm.Model{ID: input.ID}}
	res := ls.db.WithContext(ctx).Model(&updatedList).Clauses(clause.Returning{}).Updates(list)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return dto.FromEntity(updatedList), nil
}

func (ls *ListService) Delete(ctx context.Context, id uint) (*dto.List, error) {
	var list entities.List
	res := ls.db.WithContext(ctx).Clauses(clause.Returning{}).Delete(&list, id)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return dto.FromEntity(&list), nil
}
