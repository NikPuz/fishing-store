package service

import (
	"context"
	"fishing-store/internal/entity"
)

type supplyService struct {
	supplyRepo  entity.ISupplyRepository
	productRepo entity.IProductRepository
}

func NewSupplyService(supplyRepo entity.ISupplyRepository, productRepo entity.IProductRepository) entity.ISupplyService {
	supplyService := new(supplyService)
	supplyService.supplyRepo = supplyRepo
	supplyService.productRepo = productRepo
	return supplyService
}

func (s supplyService) CreateSupplies(ctx context.Context, supplies []entity.Supply) error {
	tx, err := s.productRepo.GetTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, supply := range supplies {
		err = s.productRepo.TxUpdateProductAddStock(ctx, tx, supply.ProductId, supply.Count)
		if err != nil {
			return err
		}
	}

	err = s.supplyRepo.TxInsertSupplies(ctx, tx, supplies)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}
