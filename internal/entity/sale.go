package entity

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type ISaleService interface {
	CreateSale(ctx context.Context, sales *SaleDTO) error
}

type ISaleRepository interface {
	GetTx(ctx context.Context) (pgx.Tx, error)
	TxInsertSale(ctx context.Context, tx pgx.Tx, sale *Sale) (int, error)
	TxInsertSaleItems(ctx context.Context, tx pgx.Tx, saleId int, saleItems []SaleItem) error
}

type SaleDTO struct {
	CashierId int        `json:"cashier_id"`
	SaleItems []SaleItem `json:"sale_items"`
}

type Sale struct {
	Id        int       `json:"id"`
	Sum       int       `json:"sum"`
	CashierId int       `json:"cashier_id"`
	Date      time.Time `json:"date"`
}

type SaleItem struct {
	Id        int `json:"id"`
	SaleId    int `json:"sale_id"`
	ProductId int `json:"product_id"`
	UnitPrice int `json:"unit_price"`
	Count     int `json:"count"`
}
