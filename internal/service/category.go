package service

import (
	"context"

	"github.com/jonathanmdr/gRPC/internal/database"
	"github.com/jonathanmdr/gRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.CategoryResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) GetCategories(context.Context, *pb.EmptyRequest) (*pb.CategoriesResponse, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoriesResponse []*pb.CategoryResponse
	for _, category := range categories {
		categoryResponse := &pb.CategoryResponse{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}
	return &pb.CategoriesResponse{
		Categories: categoriesResponse,
	}, nil
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	return &pb.CategoryResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}
