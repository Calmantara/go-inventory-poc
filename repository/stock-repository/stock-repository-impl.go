package stockrepository

import (
	"context"
	"errors"

	inappmemory "github.com/Calmantara/go-inventory-poc/configuration/in-app-memory"
	"github.com/Calmantara/go-inventory-poc/entity"
)

type StockRepositoryImpl struct {
	inAppMemory inappmemory.InAppMemory
}

func NewStockRepository(
	inAppMemory inappmemory.InAppMemory,
) StockRepository {
	return &StockRepositoryImpl{
		inAppMemory: inAppMemory,
	}
}

func (s *StockRepositoryImpl) GetAllStocks(ctx context.Context, stocks *map[string]entity.StockEntity) (err error) {
	*stocks = s.inAppMemory.Stocks
	return err
}

func (s *StockRepositoryImpl) GetStockByID(ctx context.Context, stock *entity.StockEntity) (err error) {
	*stock = s.inAppMemory.Stocks[stock.ID.String()]
	if stock == nil {
		err = errors.New("NOT_FOUND")
		return err
	}
	return err
}

func (s *StockRepositoryImpl) StoreStock(ctx context.Context, stock *entity.StockEntity) (err error) {
	if err = s.inAppMemory.StoreStock(stock); err != nil {
		err = errors.New("error storing new stock")
		return err
	}
	return err
}
