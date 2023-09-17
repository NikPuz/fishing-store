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

func (r productRepository) InsertProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {

	err := r.db.QueryRow(ctx, "insert into products(name, price, barcode, stock, description, category_id, manufacturer_id) values ($1, $2, $3, $4, $5, $6, $7) returning id",
		product.Name, product.Price, product.Barcode, product.Stock, product.Description, product.CategoryId, product.ManufacturerId).Scan(
		&product.Id)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return product, nil
}

func (r productRepository) SelectProduct(ctx context.Context, id int) (*entity.ProductResponse, error) {
	var product entity.ProductResponse

	err := r.db.QueryRow(ctx,
		`SELECT p.id, p.name, p.price, p.barcode, p.description, p.stock, c.name, p.name FROM products p
LEFT JOIN categories c ON c.id = p.category_id
LEFT JOIN manufacturers m ON m.id = p.manufacturer_id
WHERE p.id = $1`, id).Scan(
		&product.Id, &product.Name, &product.Price, &product.Barcode, &product.Description, &product.Stock, &product.Category, &product.Manufacturer)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return &product, nil
}

func (r productRepository) UpdateProduct(ctx context.Context, product *entity.Product) error {

	_, err := r.db.Exec(ctx,
		`update products set name = $2, price = $3, barcode = $4, description = $5, stock = $6, category_id = $7, manufacturer_id = $8 where id = $1`,
		product.Id, product.Name, product.Price, product.Barcode, product.Description, product.Stock, product.CategoryId, product.ManufacturerId)
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

func (r productRepository) SelectAllProducts(ctx context.Context) ([]entity.ProductResponse, error) {
	products := make([]entity.ProductResponse, 0)

	rows, err := r.db.Query(ctx, `SELECT p.id, p.name, p.price, p.barcode, p.description, p.stock, c.name, m.name FROM products p
LEFT JOIN categories c ON c.id = p.category_id
LEFT JOIN manufacturers m ON m.id = p.manufacturer_id`)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.ProductResponse
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Barcode,
			&product.Description,
			&product.Stock,
			&product.Category,
			&product.Manufacturer,
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

func (r productRepository) SetDefaultManufacturerByManufacturerId(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `update products set manufacturer_id = 0 where manufacturer_id = $1`, id)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r productRepository) SetDefaultCategoryByCategoryId(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `update products set category_id = 0 where category_id = $1`, id)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}

func (r productRepository) TxInsertProduct(ctx context.Context, tx pgx.Tx, product *entity.Product) (*entity.Product, error) {
	err := tx.QueryRow(ctx, "insert into products(name, price, barcode, stock, description, category_id, manufacturer_id) values ($1, $2, $3, $4, $5, $6, $7) returning id",
		product.Name, product.Price, product.Barcode, product.Stock, product.Description, product.CategoryId, product.ManufacturerId).Scan(
		&product.Id)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return product, nil
}

func (r productRepository) TxUpdateBarcode(ctx context.Context, tx pgx.Tx, id int, barcode string) error {

	_, err := tx.Exec(ctx, `update products set barcode = $2 where id = $1`, id, barcode)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}
