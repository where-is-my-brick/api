package services

import (
	"context"
	"github.com/where-is-my-brick/api/grpc/category_service"
	"log"
	"os"
)

type CategoryServiceServer struct {
	PathPrefix string
	category_service.UnimplementedCategoryServiceServer
}

func (s *CategoryServiceServer) CreateCategory(ctx context.Context, req *category_service.CreateCategoryRequest) (*category_service.CreateCategoryResponse, error) {
	parentPath := req.Parent
	categoryName := req.Category

	err := os.MkdirAll(s.PathPrefix+parentPath+"/"+categoryName, 0755)
	if err != nil {
		log.Fatal(err)
		return &category_service.CreateCategoryResponse{
			Success: false,
		}, err
	}

	return &category_service.CreateCategoryResponse{
		Success: true,
	}, nil
}

func (s *CategoryServiceServer) ListCategories(ctx context.Context, req *category_service.ListCategoriesRequest) (*category_service.ListCategoriesResponse, error) {
	var categories []string
	parentPath := req.Parent

	dirs, err := os.ReadDir(s.PathPrefix + parentPath)
	if err != nil {
		log.Fatal(err)
		return &category_service.ListCategoriesResponse{
			Categories: categories,
		}, err
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			categories = append(categories, dir.Name())
		}
	}

	return &category_service.ListCategoriesResponse{
		Categories: categories,
	}, nil
}
