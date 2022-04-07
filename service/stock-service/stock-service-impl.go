package stockservice

import (
	"context"
	"errors"
	"math/rand"

	"github.com/Calmantara/go-inventory-poc/entity"
	stockrepository "github.com/Calmantara/go-inventory-poc/repository/stock-repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type StockServiceImpl struct {
	sugar           *zap.SugaredLogger
	stockRepository stockrepository.StockRepository
}

func NewStockService(
	sugar *zap.SugaredLogger,
	stockRepository stockrepository.StockRepository,
) StockService {
	return &StockServiceImpl{
		sugar:           sugar,
		stockRepository: stockRepository,
	}
}

func (s *StockServiceImpl) GetStockByIDService(ctx context.Context, stock *entity.StockEntity) (err error) {
	s.sugar.Infof("searching stock by id:%v", stock.ID)
	// error checking from repos
	if err = s.stockRepository.GetStockByID(ctx, stock); err != nil {
		return err
	}

	// check whether the payload is active or not
	if !stock.IsActive {
		err = errors.New("NOT_FOUND")
		return
	}

	return err
}

func (s *StockServiceImpl) StoreStockService(ctx context.Context, stock *entity.StockEntity) (err error) {
	storedID := (uuid.New())
	if stock.ID != nil {
		storedID = *stock.ID
	}

	// store stock payload to memory
	s.sugar.Infof("storing payload with id:%v", storedID)
	stock.ID = &storedID
	if err = s.stockRepository.StoreStock(ctx, stock); err != nil {
		return err
	}
	return err
}

func (s *StockServiceImpl) UpdateStockByIDService(ctx context.Context, stock *entity.StockEntity) (err error) {
	// check whether entity exist or not
	tempStock := entity.StockEntity{
		ID: stock.ID,
	}
	if err = s.GetStockByIDService(ctx, &tempStock); err != nil {
		return err
	}

	// generate update payload
	s.sugar.Infof("updating stock with id:%v", tempStock.ID)
	tempStock.Price = stock.Price
	if err = s.StoreStockService(ctx, &tempStock); err != nil {
		return err
	}
	return err
}

func (s *StockServiceImpl) UpdateAllStocksPrice(ctx context.Context) {
	var stocks map[string]entity.StockEntity
	if err := s.stockRepository.GetAllStocks(ctx, &stocks); err != nil {
		s.sugar.Errorf("error when fetch all stocks:%v", err.Error())
	}
	s.sugar.Infof("auto update with total payload:%v", len(stocks))
	go func() {
		for _, el := range stocks {
			r := 47500 + rand.Float64()*(52500-47500)
			el.Price = r
			if err := s.StoreStockService(ctx, &el); err != nil {
				s.sugar.Errorf("error auto update payload with id:%v err:%c", el.ID, err.Error())
			}
		}
	}()
}
