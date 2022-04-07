package stockcontroller

import (
	ginroutergroup "github.com/Calmantara/go-inventory-poc/configuration/gin-router-group"
	stockhandler "github.com/Calmantara/go-inventory-poc/handler/stock-handler"
)

type StockControllerImpl struct {
	ginGroup     ginroutergroup.RouterGroup
	stockHandler stockhandler.StockHandler
}

func NewStockController(
	ginGroup ginroutergroup.RouterGroup,
	stockHandler stockhandler.StockHandler,
) StockController {
	return &StockControllerImpl{
		ginGroup:     ginGroup,
		stockHandler: stockHandler,
	}
}

func (s *StockControllerImpl) getStockByID() {
	s.ginGroup.GET(
		"/:id",
		s.stockHandler.GetStockByIDHandler,
	)
}

func (s *StockControllerImpl) storeStock() {
	s.ginGroup.POST(
		"",
		s.stockHandler.StoreStockHandler,
	)
}

func (s *StockControllerImpl) updateStock() {
	s.ginGroup.PUT(
		"",
		s.stockHandler.UpdateStockByIDHandler,
	)
}

func (s *StockControllerImpl) Controllers() {
	s.getStockByID()
	s.storeStock()
	s.updateStock()
}
