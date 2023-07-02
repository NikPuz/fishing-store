package service

import (
	"context"
	"fishing-store/internal/entity"
)

type categoryService struct {
	categoryRepo entity.ICategoryRepository
	productRepo  entity.IProductRepository
}

func NewCategoryService(categoryRepo entity.ICategoryRepository, productRepo entity.IProductRepository) entity.ICategoryService {
	categoryService := new(categoryService)
	categoryService.categoryRepo = categoryRepo
	categoryService.productRepo = productRepo
	return categoryService
}

func (s categoryService) CreateCategory(ctx context.Context, category *entity.Category) error {
	return s.categoryRepo.InsertCategory(ctx, category)
}

func (s categoryService) ReadCategory(ctx context.Context, id int) (*entity.Category, error) {
	return s.categoryRepo.SelectCategory(ctx, id)
}

func (s categoryService) UpdateCategory(ctx context.Context, category *entity.Category) error {
	return s.categoryRepo.UpdateCategory(ctx, category)
}

func (s categoryService) DeleteCategory(ctx context.Context, id int) error {
	err := s.categoryRepo.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}

	return s.productRepo.SetDefaultCategoryByCategoryId(ctx, id)
}

func (s categoryService) ReadCategories(ctx context.Context) ([]entity.Category, error) {
	return s.categoryRepo.SelectAllCategories(ctx)
}
