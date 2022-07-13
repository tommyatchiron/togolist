package list

import (
	"context"

	"github.com/tommyatchiron/togolist/internal/pkg/list/dto"
	pb "github.com/tommyatchiron/togolist/internal/pkg/list/grpc"
	"google.golang.org/grpc"
)

type ListGrpcController struct {
	pb.UnimplementedListServiceServer
	listService *ListService
}

func NewListGrpcController(listService *ListService) *ListGrpcController {
	return &ListGrpcController{listService: listService}
}

func registerListGrpcServer(lc *ListGrpcController, s *grpc.Server) {
	pb.RegisterListServiceServer(s, lc)
}

func (lgc *ListGrpcController) Create(ctx context.Context, createListInput *pb.CreateListInput) (*pb.List, error) {
	createListInputDto := dto.CreateListInput{
		Title:       &createListInput.Title,
		Description: &createListInput.Description,
	}
	if createListInput.Priority != nil {
		priority := int64(*createListInput.Priority)
		createListInputDto.Priority = &priority
	}
	list, err := lgc.listService.Create(ctx, &createListInputDto)
	if err != nil {
		return nil, err
	}
	return &pb.List{
		Id:          uint64(list.ID),
		Title:       *list.Title,
		Description: *list.Description,
		Priority:    int64(*list.Priority),
	}, nil
}

func (lgc *ListGrpcController) GetAll(ctx context.Context, _ *pb.GetAllOptions) (*pb.RepeatedList, error) {
	lists, err := lgc.listService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	pbList := make([]*pb.List, len(lists))
	for i, list := range lists {
		pbList[i] = &pb.List{
			Id:          uint64(list.ID),
			Title:       *list.Title,
			Description: *list.Description,
			Priority:    int64(*list.Priority),
		}
	}
	return &pb.RepeatedList{Lists: pbList}, nil
}

func (lgc *ListGrpcController) GetOne(ctx context.Context, id *pb.ID) (*pb.OptionalList, error) {
	list, err := lgc.listService.GetOne(ctx, uint(id.Id))
	if err != nil {
		return nil, err
	}
	if list == nil {
		return &pb.OptionalList{}, nil
	}
	return &pb.OptionalList{
		List: &pb.List{
			Id:          uint64(list.ID),
			Title:       *list.Title,
			Description: *list.Description,
			Priority:    int64(*list.Priority),
		},
	}, nil
}

func (lgc *ListGrpcController) Update(ctx context.Context, updateListInput *pb.UpdateListInput) (*pb.OptionalList, error) {
	updateListInputDto := dto.UpdateListInput{
		ID: uint(updateListInput.Id),
		PartialUpdateListInput: dto.PartialUpdateListInput{
			Title:       updateListInput.Title,
			Description: updateListInput.Description,
		},
	}
	if updateListInput.Priority != nil {
		priority := int64(*updateListInput.Priority)
		updateListInputDto.Priority = &priority
	}
	list, err := lgc.listService.Update(ctx, &updateListInputDto)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return &pb.OptionalList{}, nil
	}
	return &pb.OptionalList{
		List: &pb.List{
			Id:          uint64(list.ID),
			Title:       *list.Title,
			Description: *list.Description,
			Priority:    int64(*list.Priority),
		},
	}, nil
}

func (lgc *ListGrpcController) Delete(ctx context.Context, id *pb.ID) (*pb.OptionalList, error) {
	list, err := lgc.listService.Delete(ctx, uint(id.Id))
	if err != nil {
		return nil, err
	}
	if list == nil {
		return &pb.OptionalList{}, nil
	}
	return &pb.OptionalList{
		List: &pb.List{
			Id:          uint64(list.ID),
			Title:       *list.Title,
			Description: *list.Description,
			Priority:    int64(*list.Priority),
		},
	}, nil
}
