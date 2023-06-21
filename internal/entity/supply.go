package entity

import (
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type ISupplyService interface {
	CreateSupplies(ctx context.Context, supplies []Supply) error
}

type ISupplyRepository interface {
	TxInsertSupplies(ctx context.Context, tx pgx.Tx, supplies []Supply) error
}

type Supply struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id"`
	UnitPrice int       `json:"unit_price"`
	Count     int       `json:"count"`
	Date      time.Time `json:"date"`
}
