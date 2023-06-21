package repository

import (
	"context"
	"fishing-store/internal/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) entity.IProductRepository {
	productRepository := new(productRepository)
	productRepository.db = db
	return productRepository
}

func (r productRepository) GetTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return tx, nil
}

func (r productRepository) InsertProduct(ctx context.Context, product *entity.Product) error {

	_, err := r.db.Exec(ctx, "insert into products(name, price, stock, category_id, manufacturer_id) values ($1, $2, $3, $4, $5)",
		product.Name, product.Price, product.Stock, product.CategoryId, product.ManufacturerId)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r productRepository) SelectProduct(ctx context.Context, id int) (*entity.Product, error) {
	var product entity.Product

	err := r.db.QueryRow(ctx,
		`SELECT id, name, price, stock, category_id, manufacturer_id FROM products WHERE id = $1`, id).Scan(
		&product.Id, &product.Name, &product.Price, &product.Stock, &product.CategoryId, &product.ManufacturerId)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return &product, nil
}

func (r productRepository) UpdateProduct(ctx context.Context, product *entity.Product) error {

	_, err := r.db.Exec(ctx,
		`update products set name = $2, price = $3, stock = $4, category_id = $5, manufacturer_id = $6 where id = $1`,
		product.Id, product.Name, product.Price, product.Stock, product.CategoryId, product.ManufacturerId)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r productRepository) DeleteProduct(ctx context.Context, id int) error {

	_, err := r.db.Exec(ctx, `delete from products where id=$1`, id)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r productRepository) SelectAllProducts(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product

	rows, err := r.db.Query(ctx, "SELECT id, name, price, stock, category_id, manufacturer_id FROM products")
	if err != nil {
		return nil, entity.NewError(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product

		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CategoryId,
			&product.ManufacturerId,
		)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		products = append(products, product)
	}

	return products, nil
}

func (r productRepository) TxUpdateProductAddStock(ctx context.Context, tx pgx.Tx, id, addStock int) error {

	_, err := tx.Exec(ctx, `update products set stock = stock+$2 where id = $1`, id, addStock)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}
