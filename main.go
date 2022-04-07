package main

import (
	"context"

	ginrouter "github.com/Calmantara/go-inventory-poc/configuration/gin-router"
	inappmemory "github.com/Calmantara/go-inventory-poc/configuration/in-app-memory"
	zaplogger "github.com/Calmantara/go-inventory-poc/configuration/zap-logger"
	stockcontroller "github.com/Calmantara/go-inventory-poc/controller/stock-controller"
	commonhandler "github.com/Calmantara/go-inventory-poc/handler/common-handler"
	stockhandler "github.com/Calmantara/go-inventory-poc/handler/stock-handler"
	corsmiddleware "github.com/Calmantara/go-inventory-poc/middleware/cors-middleware"
	stockrepository "github.com/Calmantara/go-inventory-poc/repository/stock-repository"
	cronservice "github.com/Calmantara/go-inventory-poc/service/cron-service"
	stockservice "github.com/Calmantara/go-inventory-poc/service/stock-service"
	"github.com/gin-gonic/gin"
)

func main() {
	// enable auto update
	isEnable := true

	// initiate router, logger, and other configs
	ginRouter := ginrouter.NewGinRouter()
	sugarLogger := zaplogger.NewWrappedZapLogger()
	inAppMemory := inappmemory.NewInAppMemory()

	// repository
	stockRepository := stockrepository.NewStockRepository(inAppMemory)

	// service
	stockService := stockservice.NewStockService(sugarLogger, stockRepository)
	cronService := cronservice.NewCronJobService(sugarLogger)

	// middleware
	corsMiddleware := corsmiddleware.NewCorsMiddleware()

	// handler
	commonHandler := commonhandler.NewCommonHandler()
	stockHandler := stockhandler.NewStockHandler(sugarLogger, commonHandler, stockService)

	// auto update price every 30 seconds
	if isEnable {
		sugarLogger.Info("enabling auto update")
		cronService.CreateRepetitionSecondJob(
			context.Background(), 15,
			stockService.UpdateAllStocksPrice)
	}

	// set common middleware
	ginRouter.USE(
		gin.Logger(),
		gin.Recovery(),
		corsMiddleware.CorsRequest,
	)

	//serve controllers
	stockGroup := ginRouter.GROUP("/stock")
	stockcontroller.NewStockController(stockGroup, stockHandler).Controllers()

	ginRouter.SERVE(ginrouter.WithPort("8080"))
}
