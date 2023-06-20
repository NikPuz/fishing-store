package entity

import "context"

type IProductService interface {
	CreateProduct(ctx context.Context, product *Product) error
	ReadProduct(ctx context.Context, id int) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id int) error
	ReadProducts(ctx context.Context) ([]Product, error)
}

type IProductRepository interface {
	InsertProduct(ctx context.Context, product *Product) error
	SelectProduct(ctx context.Context, id int) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id int) error
	SelectAllProducts(ctx context.Context) ([]Product, error)
}

type Product struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Price          float32 `json:"price"`
	Stock          int     `json:"stock"`
	CategoryId     int     `json:"category_id"`
	ManufacturerId *int    `json:"manufacturer_id"`
}
