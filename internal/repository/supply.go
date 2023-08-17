package repository

import (
	"context"
	"fishing-store/internal/entity"
	"fmt"
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
	fmt.Println(supply.Sum)
	err := tx.QueryRow(ctx, "insert into supplies(sum) values ($1) returning id",
		supply.Sum).Scan(
		&supply.Id)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return supply, nil
}

func (r supplyRepository) SelectAllSupplies(ctx context.Context) ([]entity.Supply, error) {
	var supplies []entity.Supply

	rows, err := r.db.Query(ctx, `SELECT s.id, s.sum, s.date FROM supplies s`)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		var supply entity.Supply
		err = rows.Scan(
			&supply.Id,
			&supply.Sum,
			&supply.Date,
		)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		supplies = append(supplies, supply)
	}

	return supplies, nil
}

func (r supplyRepository) SelectAllSuppliesItemsResponseMap(ctx context.Context) (map[int][]entity.SupplyItemResponse, error) {
	mapSuppliesItems := make(map[int][]entity.SupplyItemResponse)

	rows, err := r.db.Query(ctx, `SELECT si.supply_id, p.name, c.name, m.name, si.unit_price, si.count FROM supplies_items si
LEFT JOIN products p ON p.id = si.product_id
LEFT JOIN categories c ON c.id = p.category_id
LEFT JOIN manufacturers m ON m.id = p.manufacturer_id`)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var supplyResponse entity.SupplyItemResponse
		err = rows.Scan(
			&id,
			&supplyResponse.ProductName,
			&supplyResponse.ProductCategory,
			&supplyResponse.ProductManufacturer,
			&supplyResponse.UnitPrice,
			&supplyResponse.Count,
		)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		if _, ok := mapSuppliesItems[id]; ok {
			mapSuppliesItems[id] = append(mapSuppliesItems[id], supplyResponse)
		} else {
			mapSuppliesItems[id] = []entity.SupplyItemResponse{supplyResponse}
		}
	}

	return mapSuppliesItems, nil
}

func (r supplyRepository) TxInsertSuppliesItems(ctx context.Context, tx pgx.Tx, supplyId int, supplyItems []entity.SupplyItem) error {

	rows := make([][]interface{}, 0, len(supplyItems))

	for _, item := range supplyItems {
		rows = append(rows, []interface{}{supplyId, item.ProductId, item.UnitPrice, item.Count})
	}

	_, err := tx.CopyFrom(ctx, pgx.Identifier{"supplies_items"}, []string{"supply_id", "product_id", "unit_price", "count"},
		pgx.CopyFromRows(rows))
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}
