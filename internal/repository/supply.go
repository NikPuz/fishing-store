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

func (r supplyRepository) TxInsertSupplies(ctx context.Context, tx pgx.Tx, supplies []entity.Supply) error {

	rows := make([][]interface{}, 0, len(supplies))

	for _, supply := range supplies {
		rows = append(rows, []interface{}{supply.ProductId, supply.UnitPrice, supply.Count})
	}

	_, err := tx.CopyFrom(ctx, pgx.Identifier{"supplies"}, []string{"product_id", "unit_price", "count"},
		pgx.CopyFromRows(rows))
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}
