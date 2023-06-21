package repository

import (
	"context"
	"fishing-store/internal/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type saleRepository struct {
	db *pgxpool.Pool
}

func NewSaleRepository(db *pgxpool.Pool) entity.ISaleRepository {
	saleRepository := new(saleRepository)
	saleRepository.db = db
	return saleRepository
}

func (r saleRepository) GetTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return tx, nil
}

func (r saleRepository) TxInsertSale(ctx context.Context, tx pgx.Tx, sale *entity.Sale) (int, error) {
	var id int

	rows, err := tx.Query(ctx, "insert into sales(sum, cashier_id) values ($1, $2) returning id", sale.Sum, sale.CashierId)
	if err != nil {
		return 0, entity.NewError(err, 500)
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, entity.NewError(err, 500)
		}
	}

	return id, nil
}

func (r saleRepository) TxInsertSaleItems(ctx context.Context, tx pgx.Tx, saleId int, saleItems []entity.SaleItem) error {

	rows := make([][]interface{}, 0, len(saleItems))

	for _, sale := range saleItems {
		rows = append(rows, []interface{}{saleId, sale.ProductId, sale.UnitPrice, sale.Count})
	}

	_, err := tx.CopyFrom(ctx, pgx.Identifier{"sales_items"}, []string{"sale_id", "product_id", "unit_price", "count"},
		pgx.CopyFromRows(rows))
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}
