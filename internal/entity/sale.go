package entity

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type ISaleService interface {
	CreateSale(ctx context.Context, sales *SaleDTO) error
	UpdateSale(ctx context.Context, sales *SaleDTO) error
	ReadSales(ctx context.Context) ([]SaleResponse, error)
}

type ISaleRepository interface {
	GetTx(ctx context.Context) (pgx.Tx, error)
	TxInsertSale(ctx context.Context, tx pgx.Tx, sale *Sale) (int, error)
	TxInsertSaleItems(ctx context.Context, tx pgx.Tx, saleId int, saleItems []SaleItem) error
	UpdateRefund(ctx context.Context, id int, isRefund bool) error
	SelectAllSales(ctx context.Context) ([]Sale, error)
	SelectAllSaleItemsMap(ctx context.Context) (map[int][]SaleItem, error)
}

type SaleDTO struct {
	Id        int        `json:"id"`
	CashierId int        `json:"cashierId"`
	SaleItems []SaleItem `json:"saleItems"`
	PayType   string     `json:"payType"`
	Refund    bool       `json:"refund"`
}

type Sale struct {
	Id        int       `json:"id"`
	Sum       int       `json:"sum"`
	CashierId int       `json:"cashierId"`
	PayType   string    `json:"payType"`
	Refund    bool      `json:"refund"`
	Date      time.Time `json:"date"`
}

type SaleItem struct {
	Id          int     `json:"id"`
	SaleId      int     `json:"saleId"`
	ProductId   int     `json:"productId"`
	ProductName *string `json:"productName"`
	UnitPrice   int     `json:"unitPrice"`
	Count       int     `json:"count"`
}

type SaleResponse struct {
	Id        int        `json:"id"`
	Sum       int        `json:"sum"`
	CashierId int        `json:"cashierId"`
	PayType   string     `json:"payType"`
	Date      time.Time  `json:"date"`
	Refund    bool       `json:"refund"`
	SaleItems []SaleItem `json:"saleItems"`
}
