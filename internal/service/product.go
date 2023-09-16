package service

import (
	"context"
	"fishing-store/internal/entity"
	"github.com/nicholassm/go-ean"
	"strconv"
)

type productService struct {
	productRepo entity.IProductRepository
}

func NewProductService(productRepo entity.IProductRepository) entity.IProductService {
	productService := new(productService)
	productService.productRepo = productRepo
	return productService
}

func (s productService) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	if len(product.Barcode) == 0 {
		tx, err := s.productRepo.GetTx(ctx)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}
		defer tx.Rollback(ctx)

		// Записываем Продукт
		product, err = s.productRepo.TxInsertProduct(ctx, tx, product)

		// Получаем Штрихкод
		eanProductId := "00000"[len(strconv.Itoa(product.Id)):] + strconv.Itoa(product.Id)

		code, err := ean.ChecksumEan13("460" + "0000" + eanProductId + "0")
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		barcode := "460" + "0000" + eanProductId + strconv.Itoa(code)

		// Записываем Штрихкод
		err = s.productRepo.TxUpdateBarcode(ctx, tx, product.Id, barcode)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		err = tx.Commit(ctx)
		if err != nil {
			return nil, entity.NewError(err, 500)
		}

		return product, err
	} else {
		return s.productRepo.InsertProduct(ctx, product)
	}
}

func (s productService) ReadProduct(ctx context.Context, id int) (*entity.ProductResponse, error) {
	return s.productRepo.SelectProduct(ctx, id)
}

func (s productService) UpdateProduct(ctx context.Context, product *entity.Product) error {
	if len(product.Barcode) == 0 {
		eanProductId := "00000"[len(strconv.Itoa(product.Id)):] + strconv.Itoa(product.Id)

		code, err := ean.ChecksumEan13("460" + "0000" + eanProductId + "0")
		if err != nil {
			return entity.NewError(err, 500)
		}

		barcode := "460" + "0000" + eanProductId + strconv.Itoa(code)

		product.Barcode = barcode
	}

	return s.productRepo.UpdateProduct(ctx, product)
}

func (s productService) DeleteProduct(ctx context.Context, id int) error {
	return s.productRepo.DeleteProduct(ctx, id)
}

func (s productService) ReadProducts(ctx context.Context) ([]entity.ProductResponse, error) {
	return s.productRepo.SelectAllProducts(ctx)
}
