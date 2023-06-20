package service

import (
	"context"
	"fishing-store/internal/entity"
)

type categoryService struct {
	categoryRepo entity.ICategoryRepository
}

func NewCategoryService(categoryRepo entity.ICategoryRepository) entity.ICategoryService {
	categoryService := new(categoryService)
	categoryService.categoryRepo = categoryRepo
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
	return s.categoryRepo.DeleteCategory(ctx, id)
}

func (s categoryService) ReadCategories(ctx context.Context) ([]entity.Category, error) {
	return s.categoryRepo.SelectAllCategories(ctx)
}
