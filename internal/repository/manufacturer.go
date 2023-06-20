package repository

import (
	"context"
	"fishing-store/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type manufacturerRepository struct {
	db *pgxpool.Pool
}

func NewManufacturerRepository(db *pgxpool.Pool) entity.IManufacturerRepository {
	manufacturerRepository := new(manufacturerRepository)
	manufacturerRepository.db = db
	return manufacturerRepository
}

func (r manufacturerRepository) InsertManufacturer(ctx context.Context, manufacturer *entity.Manufacturer) error {

	_, err := r.db.Exec(ctx, "insert into manufacturer(name) values ($1)", manufacturer.Name)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r manufacturerRepository) SelectManufacturer(ctx context.Context, id int) (*entity.Manufacturer, error) {
	var manufacturer entity.Manufacturer

	err := r.db.QueryRow(ctx, `SELECT id, name FROM manufacturer WHERE id = $1`, id).Scan(&manufacturer.Id, &manufacturer.Name)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return &manufacturer, nil
}

func (r manufacturerRepository) UpdateManufacturer(ctx context.Context, manufacturer *entity.Manufacturer) error {

	_, err := r.db.Exec(ctx, `update manufacturer set name = $1 where id = $2`, manufacturer.Name, manufacturer.Id)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r manufacturerRepository) DeleteManufacturer(ctx context.Context, id int) error {

	_, err := r.db.Exec(ctx, `delete from manufacturer where id=$1`, id)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r manufacturerRepository) SelectAllManufacturers(ctx context.Context) ([]entity.Manufacturer, error) {
	var manufacturers []entity.Manufacturer

	rows, err := r.db.Query(ctx, "SELECT id, name FROM manufacturer")
	if err != nil {
		return nil, entity.NewError(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		var manufacturer entity.Manufacturer

		err = rows.Scan(
			&manufacturer.Id,
			&manufacturer.Name,
		)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		manufacturers = append(manufacturers, manufacturer)
	}

	return manufacturers, nil
}
