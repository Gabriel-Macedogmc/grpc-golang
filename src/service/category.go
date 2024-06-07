package service

import (
	"context"

	"github.com/Gabriel-Macedogmc/grpc-golang/src/infra/gorm/interfaces"
	pb "github.com/Gabriel-Macedogmc/grpc-golang/src/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type categoryServiceServer struct {
	pb.UnimplementedCategoryServiceServer
	repository interfaces.CategoryRepository
}

func NewService(repository interfaces.CategoryRepository) *categoryServiceServer {
	return &categoryServiceServer{
		repository: repository,
	}
}

func (svc *categoryServiceServer) Save(ctx context.Context, c *pb.Category) (*pb.CategoryResponse, error) {
	category, _ := svc.repository.Create(c)

	return &pb.CategoryResponse{
		Id:          int64(category.ID),
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
		CreatedAt:   timestamppb.New(category.CreatedAt),
		UpdatedAt:   timestamppb.New(category.UpdatedAt),
	}, nil
}

func (svc *categoryServiceServer) Find(ctx context.Context, c *pb.Empty) (*pb.CategoryListResponse, error) {
	categories, _ := svc.repository.List()

	var categoriesResponse []*pb.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, &pb.CategoryResponse{
			Id:          int64(category.ID),
			Name:        category.Name,
			Description: category.Description,
			IsActive:    category.IsActive,
			CreatedAt:   timestamppb.New(category.CreatedAt),
			UpdatedAt:   timestamppb.New(category.UpdatedAt),
		})
	}

	return &pb.CategoryListResponse{
		Categories: categoriesResponse,
	}, nil
}
