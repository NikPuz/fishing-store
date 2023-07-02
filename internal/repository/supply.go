package repository

import (
	"context"
	"fishing-store/internal/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type supplyRepository struct {
	db *pgxpool.Pool
}

func NewSupplyRepository(db *pgxpool.Pool) entity.ISupplyRepository {
	supplyRepository := new(supplyRepository)
	supplyRepository.db = db
	return supplyRepository
}

func (r supplyRepository) TxInsertSupply(ctx context.Context, tx pgx.Tx, supply *entity.Supply) (*entity.Supply, error) {

	err := r.db.QueryRow(ctx, "insert into supplies(product_id, unit_price, count) values ($1, $2, $3) returning id",
		supply.ProductId, supply.UnitPrice, supply.Count).Scan(
		&supply.Id)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return supply, nil
}

func (r supplyRepository) SelectAllSupplies(ctx context.Context) ([]entity.SupplyResponse, error) {
	var supplies []entity.SupplyResponse

	rows, err := r.db.Query(ctx, `SELECT s.id, p.name, c.name, m.name, s.unit_price, s.count, s.date FROM supplies s
LEFT JOIN products p ON p.id = s.product_id
LEFT JOIN categories c ON c.id = p.category_id
LEFT JOIN manufacturers m ON m.id = p.manufacturer_id`)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		var supply entity.SupplyResponse
		err = rows.Scan(
			&supply.Id,
			&supply.ProductName,
			&supply.ProductCategory,
			&supply.ProductManufacturer,
			&supply.UnitPrice,
			&supply.Count,
			&supply.Date,
		)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		supplies = append(supplies, supply)
	}

	return supplies, nil
}

//func (r supplyRepository) TxInsertSupplies(ctx context.Context, tx pgx.Tx, supplies []entity.Supply) error {
//
//	rows := make([][]interface{}, 0, len(supplies))
//
//	for _, supply := range supplies {
//		rows = append(rows, []interface{}{supply.ProductId, supply.UnitPrice, supply.Count})
//	}
//
//	_, err := tx.CopyFrom(ctx, pgx.Identifier{"supplies"}, []string{"product_id", "unit_price", "count"},
//		pgx.CopyFromRows(rows))
//	if err != nil {
//		return entity.NewError(err, 500)
//	}
//
//	return nil
//}
