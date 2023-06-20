package service

import (
	"context"
	"fishing-store/internal/entity"
)

type manufacturerService struct {
	manufacturerRepo entity.IManufacturerRepository
}

func NewManufacturerService(manufacturerRepo entity.IManufacturerRepository) entity.IManufacturerService {
	manufacturerService := new(manufacturerService)
	manufacturerService.manufacturerRepo = manufacturerRepo
	return manufacturerService
}

func (s manufacturerService) CreateManufacturer(ctx context.Context, manufacturer *entity.Manufacturer) error {
	return s.manufacturerRepo.InsertManufacturer(ctx, manufacturer)
}

func (s manufacturerService) ReadManufacturer(ctx context.Context, id int) (*entity.Manufacturer, error) {
	return s.manufacturerRepo.SelectManufacturer(ctx, id)
}

func (s manufacturerService) UpdateManufacturer(ctx context.Context, manufacturer *entity.Manufacturer) error {
	return s.manufacturerRepo.UpdateManufacturer(ctx, manufacturer)
}

func (s manufacturerService) DeleteManufacturer(ctx context.Context, id int) error {
	return s.manufacturerRepo.DeleteManufacturer(ctx, id)
}

func (s manufacturerService) ReadManufacturers(ctx context.Context) ([]entity.Manufacturer, error) {
	return s.manufacturerRepo.SelectAllManufacturers(ctx)
}
