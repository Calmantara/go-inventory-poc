package stockrepository

import (
	"context"

	"github.com/Calmantara/go-inventory-poc/entity"
)

type StockRepository interface {
	GetAllStocks(ctx context.Context, stocks *map[string]entity.StockEntity) (err error)
	GetStockByID(ctx context.Context, stock *entity.StockEntity) (err error)
	StoreStock(ctx context.Context, stock *entity.StockEntity) (err error)
}
