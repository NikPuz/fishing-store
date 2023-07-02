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

func (s supplyService) CreateSupply(ctx context.Context, supply *entity.Supply) (*entity.Supply, error) {
	tx, err := s.productRepo.GetTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	supply, err = s.supplyRepo.TxInsertSupply(ctx, tx, supply)
	if err != nil {
		return nil, err
	}

	err = s.productRepo.TxUpdateProductAddStock(ctx, tx, supply.ProductId, supply.Count)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return supply, nil
}

func (s supplyService) ReadSupplies(ctx context.Context) ([]entity.SupplyResponse, error) {
	return s.supplyRepo.SelectAllSupplies(ctx)
}
