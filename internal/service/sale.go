package service

import (
	"context"
	"fishing-store/internal/entity"
)

type saleService struct {
	saleRepo    entity.ISaleRepository
	productRepo entity.IProductRepository
}

func NewSaleService(saleRepo entity.ISaleRepository, productRepo entity.IProductRepository) entity.ISaleService {
	saleService := new(saleService)
	saleService.saleRepo = saleRepo
	saleService.productRepo = productRepo
	return saleService
}

func (s saleService) CreateSale(ctx context.Context, sales *entity.SaleDTO) error {
	tx, err := s.saleRepo.GetTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var sum int
	for _, item := range sales.SaleItems {
		sum += item.UnitPrice * item.Count
	}

	for _, sale := range sales.SaleItems {
		err = s.productRepo.TxUpdateProductAddStock(ctx, tx, sale.ProductId, -sale.Count)
		if err != nil {
			return err
		}
	}

	saleId, err := s.saleRepo.TxInsertSale(ctx, tx, &entity.Sale{Sum: sum, CashierId: sales.CashierId})
	if err != nil {
		return err
	}

	err = s.saleRepo.TxInsertSaleItems(ctx, tx, saleId, sales.SaleItems)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return entity.NewError(err, 500)
	}

	return nil
}
