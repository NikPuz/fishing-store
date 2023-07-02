package entity

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type ISupplyService interface {
	CreateSupply(ctx context.Context, supply *Supply) (*Supply, error)
	ReadSupplies(ctx context.Context) ([]SupplyResponse, error)
}

type ISupplyRepository interface {
	TxInsertSupply(ctx context.Context, tx pgx.Tx, supply *Supply) (*Supply, error)
	SelectAllSupplies(ctx context.Context) ([]SupplyResponse, error)
}

type Supply struct {
	Id        int       `json:"id"`
	ProductId int       `json:"productId"`
	UnitPrice int       `json:"unitPrice"`
	Count     int       `json:"count"`
	Date      time.Time `json:"date"`
}

type SupplyResponse struct {
	Id                  int       `json:"id"`
	ProductName         *string   `json:"productName"`
	ProductCategory     *string   `json:"productCategory"`
	ProductManufacturer *string   `json:"productManufacturer"`
	UnitPrice           int       `json:"unitPrice"`
	Count               int       `json:"count"`
	Date                time.Time `json:"date"`
}
