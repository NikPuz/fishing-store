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

func (s supplyService) CreateSupply(ctx context.Context, supplyItems []entity.SupplyItem) (*entity.Supply, error) {
	//tx, err := s.productRepo.GetTx(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//defer tx.Rollback(ctx)
	//
	//supply, err = s.supplyRepo.TxInsertSupply(ctx, tx, supply)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = s.productRepo.TxUpdateProductAddStock(ctx, tx, supply.ProductId, supply.Count)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = tx.Commit(ctx)
	//if err != nil {
	//	return nil, entity.NewError(err, 500)
	//}
	//
	//return supply, nil

	tx, err := s.productRepo.GetTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var sum int
	for _, item := range supplyItems {
		sum += item.UnitPrice * item.Count
	}

	for _, item := range supplyItems {
		err = s.productRepo.TxUpdateProductAddStock(ctx, tx, item.ProductId, item.Count)
		if err != nil {
			return nil, err
		}
	}

	supply, err := s.supplyRepo.TxInsertSupply(ctx, tx, &entity.Supply{Sum: sum})
	if err != nil {
		return nil, err
	}

	err = s.supplyRepo.TxInsertSuppliesItems(ctx, tx, supply.Id, supplyItems)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, entity.NewError(err, 500)
	}

	return nil, nil
}

func (s supplyService) ReadSupplies(ctx context.Context) ([]entity.SupplyResponse, error) {

	supplies, err := s.supplyRepo.SelectAllSupplies(ctx)
	if err != nil {
		return nil, err
	}

	supplyResponse := make([]entity.SupplyResponse, 0, len(supplies))

	suppliesItemsResponseMap, err := s.supplyRepo.SelectAllSuppliesItemsResponseMap(ctx)
	if err != nil {
		return nil, err
	}

	for _, supply := range supplies {
		supplyResponse = append(supplyResponse, entity.SupplyResponse{
			Id:          supply.Id,
			SupplyItems: suppliesItemsResponseMap[supply.Id],
			Sum:         supply.Sum,
			Date:        supply.Date,
		})
	}

	return supplyResponse, nil
}
