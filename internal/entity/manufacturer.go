package entity

import "context"

type IManufacturerService interface {
	CreateManufacturer(ctx context.Context, manufacturer *Manufacturer) error
	ReadManufacturer(ctx context.Context, id int) (*Manufacturer, error)
	UpdateManufacturer(ctx context.Context, manufacturer *Manufacturer) error
	DeleteManufacturer(ctx context.Context, id int) error
	ReadManufacturers(ctx context.Context) ([]Manufacturer, error)
}

type IManufacturerRepository interface {
	InsertManufacturer(ctx context.Context, manufacturer *Manufacturer) error
	SelectManufacturer(ctx context.Context, id int) (*Manufacturer, error)
	UpdateManufacturer(ctx context.Context, manufacturer *Manufacturer) error
	DeleteManufacturer(ctx context.Context, id int) error
	SelectAllManufacturers(ctx context.Context) ([]Manufacturer, error)
}

type Manufacturer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
