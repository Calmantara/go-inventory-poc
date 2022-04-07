package stockhandler

import "github.com/gin-gonic/gin"

type StockHandler interface {
	GetStockByIDHandler(ctx *gin.Context)
	StoreStockHandler(ctx *gin.Context)
	UpdateStockByIDHandler(ctx *gin.Context)
}
