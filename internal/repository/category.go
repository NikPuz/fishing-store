package repository

import (
	"context"
	"fishing-store/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) entity.ICategoryRepository {
	categoryRepository := new(categoryRepository)
	categoryRepository.db = db
	return categoryRepository
}

func (r categoryRepository) InsertCategory(ctx context.Context, category *entity.Category) error {

	_, err := r.db.Exec(ctx, "insert into categories(name) values ($1)", category.Name)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r categoryRepository) SelectCategory(ctx context.Context, id int) (*entity.Category, error) {
	var category entity.Category

	err := r.db.QueryRow(ctx, `SELECT id, name FROM categories WHERE id = $1`, id).Scan(&category.Id, &category.Name)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return &category, nil
}

func (r categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {

	_, err := r.db.Exec(ctx, `update categories set name = $1 where id = $2`, category.Name, category.Id)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r categoryRepository) DeleteCategory(ctx context.Context, id int) error {

	_, err := r.db.Exec(ctx, `delete from categories where id=$1`, id)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r categoryRepository) SelectAllCategories(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category

	rows, err := r.db.Query(ctx, "SELECT id, name FROM categories")
	if err != nil {
		return nil, entity.NewError(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		var category entity.Category

		err = rows.Scan(
			&category.Id,
			&category.Name,
		)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		categories = append(categories, category)
	}

	return categories, nil
}
