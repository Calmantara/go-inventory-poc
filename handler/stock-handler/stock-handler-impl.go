package stockhandler

import (
	"encoding/json"

	"github.com/Calmantara/go-inventory-poc/entity"
	commonhandler "github.com/Calmantara/go-inventory-poc/handler/common-handler"
	stockservice "github.com/Calmantara/go-inventory-poc/service/stock-service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type StockHandlerImpl struct {
	sugar         *zap.SugaredLogger
	commonHandler commonhandler.CommonHandler
	stockService  stockservice.StockService
}

func NewStockHandler(
	sugar *zap.SugaredLogger,
	commonHandler commonhandler.CommonHandler,
	stockService stockservice.StockService,
) StockHandler {
	return &StockHandlerImpl{
		sugar:         sugar,
		commonHandler: commonHandler,
		stockService:  stockService,
	}
}

func (s *StockHandlerImpl) GetStockByIDHandler(ctx *gin.Context) {
	// get id from param
	paramID := ctx.Param("id")

	// check uuid
	uuid_, err := uuid.Parse(paramID)
	if err != nil {
		s.sugar.Errorf("error parsing uuid:%v", err.Error())
		s.commonHandler.CommonErrorResponseBuilder(ctx, commonhandler.BAD_REQUEST, err.Error())
		return
	}

	// get stock
	stock := entity.StockEntity{
		ID: &(uuid_),
	}
	if err := s.stockService.GetStockByIDService(ctx.Copy(), &stock); err != nil {
		s.sugar.Errorf("error when processing get stock with id:%v err:%v", uuid_, err.Error())
		s.commonHandler.CommonErrorResponseBuilder(ctx, commonhandler.ErrorType(err.Error()), err.Error())
		return
	}

	// success response
	s.commonHandler.CommonResponseBuilder(ctx, commonhandler.SUCCESS, stock)
}

func (s *StockHandlerImpl) StoreStockHandler(ctx *gin.Context) {
	// bind payload
	var stock entity.StockEntity
	if err := ctx.ShouldBind(&stock); err != nil {
		s.sugar.Errorf("error when binding stock payload:%v", err.Error())
		s.commonHandler.CommonErrorResponseBuilder(ctx, commonhandler.BAD_REQUEST, err.Error())
		return
	}

	// process service
	if err := s.stockService.StoreStockService(ctx.Copy(), &stock); err != nil {
		s.sugar.Errorf("error when processing store service:%v", err.Error())
		s.commonHandler.CommonErrorResponseBuilder(ctx, commonhandler.INTERNAL, err.Error())
		return
	}

	//success response
	s.commonHandler.CommonResponseBuilder(ctx, commonhandler.SUCCESS,
		map[string]interface{}{
			"message": "success fully store stock",
			"payload": stock,
		})
}

func (s *StockHandlerImpl) UpdateStockByIDHandler(ctx *gin.Context) {
	// bind payload
	stock := struct {
		ID    uuid.UUID `json:"id" binding:"required"`
		Price float64   `json:"price" binding:"required,gte=1"`
	}{}
	if err := ctx.ShouldBind(&stock); err != nil {
		s.sugar.Errorf("error when binding stock payload:%v", err.Error())
		s.commonHandler.CommonErrorResponseBuilder(ctx, commonhandler.BAD_REQUEST, err.Error())
		return
	}

	// transform stock payload to entity
	b, err := json.Marshal(&stock)
	if err != nil {
		s.sugar.Errorf("error when marshaling payload:%v", err.Error())
		s.commonHandler.CommonErrorResponseBuilder(ctx, commonhandler.INTERNAL, err.Error())
		return
	}
	var stock_ entity.StockEntity
	err = json.Unmarshal(b, &stock_)
	if err != nil {
		s.sugar.Errorf("error when unmarshaling payload:%v", err.Error())
		s.commonHandler.CommonErrorResponseBuilder(ctx, commonhandler.INTERNAL, err.Error())
		return
	}

	// process service
	if err := s.stockService.UpdateStockByIDService(ctx.Copy(), &stock_); err != nil {
		s.sugar.Errorf("error when processing update service:%v", err.Error())
		s.commonHandler.CommonErrorResponseBuilder(ctx, commonhandler.INTERNAL, err.Error())
		return
	}

	//success response
	s.commonHandler.CommonResponseBuilder(ctx, commonhandler.SUCCESS,
		map[string]interface{}{
			"message": "success fully update stock",
			"payload": stock,
		})
}
