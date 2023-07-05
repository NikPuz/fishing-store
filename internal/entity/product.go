package entity

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type IProductService interface {
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	ReadProduct(ctx context.Context, id int) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id int) error
	ReadProducts(ctx context.Context) ([]ProductResponse, error)
}

type IProductRepository interface {
	GetTx(ctx context.Context) (pgx.Tx, error)
	InsertProduct(ctx context.Context, product *Product) (*Product, error)
	SelectProduct(ctx context.Context, id int) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id int) error
	SelectAllProducts(ctx context.Context) ([]ProductResponse, error)
	TxUpdateProductAddStock(ctx context.Context, tx pgx.Tx, id, addStock int) error
	SetDefaultManufacturerByManufacturerId(ctx context.Context, id int) error
	SetDefaultCategoryByCategoryId(ctx context.Context, id int) error
}

type Product struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Description    string `json:"description"`
	Stock          int    `json:"stock"`
	CategoryId     int    `json:"categoryId"`
	ManufacturerId int    `json:"manufacturerId"`
}

type ProductResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Description  string `json:"description"`
	Stock        int    `json:"stock"`
	Category     string `json:"category"`
	Manufacturer string `json:"manufacturer"`
}
