package list

import (
	"context"

	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"github.com/tommyatchiron/togolist/internal/pkg/list/dto"
)

type ListResolver struct {
	listService *ListService
}

func NewListResolver(listService *ListService) *ListResolver {
	return &ListResolver{listService: listService}
}

func registerListResolver(listResolver *ListResolver, schema *schemabuilder.Schema) {
	schema.Object("List", dto.List{})
	mut := schema.Mutation()
	mut.FieldFunc("createList", listResolver.Create)
}

func (lr *ListResolver) Create(ctx context.Context, createListInput dto.CreateListInput) (*dto.List, error) {
	return lr.listService.Create(ctx, &createListInput)
}
