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

func (r saleRepository) SelectAllSales(ctx context.Context) ([]entity.Sale, error) {
	var sales []entity.Sale

	rows, err := r.db.Query(ctx, `SELECT s.id, s.sum, s.cashier_id, s.date FROM sales s`)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		var sale entity.Sale
		err = rows.Scan(
			&sale.Id,
			&sale.Sum,
			&sale.CashierId,
			&sale.Date,
		)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		sales = append(sales, sale)
	}

	return sales, nil
}

func (r saleRepository) SelectAllSaleItemsMap(ctx context.Context) (map[int][]entity.SaleItem, error) {
	saleItems := make(map[int][]entity.SaleItem)

	rows, err := r.db.Query(ctx, `SELECT s.id, s.sale_id, s.product_id, p.name, s.unit_price, s.count FROM sales_items s
LEFT JOIN products p ON p.id = s.product_id`)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		var saleItem entity.SaleItem
		var id int
		err = rows.Scan(
			&saleItem.Id,
			&id,
			&saleItem.ProductId,
			&saleItem.ProductName,
			&saleItem.UnitPrice,
			&saleItem.Count,
		)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		if _, ok := saleItems[id]; ok {
			saleItems[id] = append(saleItems[id], saleItem)
		} else {
			saleItems[id] = []entity.SaleItem{saleItem}
		}
	}

	return saleItems, nil
}
