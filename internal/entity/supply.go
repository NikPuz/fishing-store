package entity

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type ISupplyService interface {
	CreateSupply(ctx context.Context, supplyItems []SupplyItem) (*Supply, error)
	ReadSupplies(ctx context.Context) ([]SupplyResponse, error)
}

type ISupplyRepository interface {
	TxInsertSupply(ctx context.Context, tx pgx.Tx, supply *Supply) (*Supply, error)
	SelectAllSuppliesItemsResponseMap(ctx context.Context) (map[int][]SupplyItemResponse, error)
	SelectAllSupplies(ctx context.Context) ([]Supply, error)
	TxInsertSuppliesItems(ctx context.Context, tx pgx.Tx, supplyId int, supplyItems []SupplyItem) error
}

type Supply struct {
	Id   int       `json:"id"`
	Sum  int       `json:"sum"`
	Date time.Time `json:"date"`
}

type SupplyItem struct {
	Id        int `json:"id"`
	SupplyId  int `json:"supplyId"`
	ProductId int `json:"productId"`
	UnitPrice int `json:"unitPrice"`
	Count     int `json:"count"`
}

type SupplyItemResponse struct {
	ProductName         *string `json:"productName"`
	ProductCategory     *string `json:"productCategory"`
	ProductManufacturer *string `json:"productManufacturer"`
	UnitPrice           int     `json:"unitPrice"`
	Count               int     `json:"count"`
}

type SupplyResponse struct {
	Id          int                  `json:"id"`
	SupplyItems []SupplyItemResponse `json:"supplyItems"`
	Sum         int                  `json:"sum"`
	Date        time.Time            `json:"date"`
}
