package entity

import "context"

type ICategoryService interface {
	CreateCategory(ctx context.Context, category *Category) error
	ReadCategory(ctx context.Context, id int) (*Category, error)
	UpdateCategory(ctx context.Context, category *Category) error
	DeleteCategory(ctx context.Context, id int) error
	ReadCategories(ctx context.Context) ([]Category, error)
}

type ICategoryRepository interface {
	InsertCategory(ctx context.Context, category *Category) error
	SelectCategory(ctx context.Context, id int) (*Category, error)
	UpdateCategory(ctx context.Context, category *Category) error
	DeleteCategory(ctx context.Context, id int) error
	SelectAllCategories(ctx context.Context) ([]Category, error)
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
