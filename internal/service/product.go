package service

import (
	"context"
	"fishing-store/internal/entity"
)

type productService struct {
	productRepo entity.IProductRepository
}

func NewProductService(productRepo entity.IProductRepository) entity.IProductService {
	productService := new(productService)
	productService.productRepo = productRepo
	return productService
}

func (s productService) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	return s.productRepo.InsertProduct(ctx, product)
}

func (s productService) ReadProduct(ctx context.Context, id int) (*entity.ProductResponse, error) {
	return s.productRepo.SelectProduct(ctx, id)
}

func (s productService) UpdateProduct(ctx context.Context, product *entity.Product) error {
	return s.productRepo.UpdateProduct(ctx, product)
}

func (s productService) DeleteProduct(ctx context.Context, id int) error {
	return s.productRepo.DeleteProduct(ctx, id)
}

func (s productService) ReadProducts(ctx context.Context) ([]entity.ProductResponse, error) {
	return s.productRepo.SelectAllProducts(ctx)
}
