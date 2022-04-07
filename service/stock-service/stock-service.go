package stockservice

import (
	"context"

	"github.com/Calmantara/go-inventory-poc/entity"
)

type StockService interface {
	GetStockByIDService(ctx context.Context, stock *entity.StockEntity) (err error)
	StoreStockService(ctx context.Context, stock *entity.StockEntity) (err error)
	UpdateStockByIDService(ctx context.Context, stock *entity.StockEntity) (err error)

	// misc
	UpdateAllStocksPrice(ctx context.Context)
}
