package stockservice

import (
	"context"
	"testing"

	inappmemory "github.com/Calmantara/go-inventory-poc/configuration/in-app-memory"
	zaplogger "github.com/Calmantara/go-inventory-poc/configuration/zap-logger"
	"github.com/Calmantara/go-inventory-poc/entity"
	stockrepository "github.com/Calmantara/go-inventory-poc/repository/stock-repository"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

var (
	sugar = zaplogger.NewWrappedZapLogger()
	Iam   = inappmemory.NewInAppMemory()
	Repo  = stockrepository.NewStockRepository(Iam)
	Serv  = NewStockService(sugar, Repo)
)

func TestStoreAndGetStockByID(t *testing.T) {
	newID := uuid.New()
	newStock := entity.StockEntity{
		ID:           &newID,
		Name:         "stock-1",
		Price:        50000,
		Availability: 100,
		IsActive:     true,
	}
	err := Serv.StoreStockService(context.Background(), &newStock)
	if err != nil {
		t.Errorf("TestGetStockByID FAILED. Expected error %v got error %v", nil, err.Error())
	} else {
		t.Logf("TestGetStockByID PASSED. Expected error %v got error %v", nil, err)
	}

	stock := entity.StockEntity{
		ID: &newID,
	}
	err = Serv.GetStockByIDService(context.Background(), &stock)

	// check equal value
	assert.Equal(t, newStock, stock)

	if err != nil {
		t.Errorf("TestStoreAndGetStockByID FAILED. Expected error %v got error %v", nil, err.Error())
	} else {
		t.Logf("TestStoreAndGetStockByID PASSED. Expected error %v got error %v", nil, err)
	}
}

func TestStoreUpdateAndGetStockByID(t *testing.T) {
	newID := uuid.New()
	newStock := entity.StockEntity{
		ID:           &newID,
		Name:         "stock-1",
		Price:        50000,
		Availability: 100,
		IsActive:     true,
	}
	err := Serv.StoreStockService(context.Background(), &newStock)
	if err != nil {
		t.Errorf("TestStoreUpdateAndGetStockByID FAILED. Expected error %v got error %v", nil, err.Error())
	} else {
		t.Logf("TestStoreUpdateAndGetStockByID PASSED. Expected error %v got error %v", nil, err)
	}

	newStock.Price = 100000
	err = Serv.UpdateStockByIDService(context.Background(), &newStock)
	if err != nil {
		t.Errorf("TestStoreUpdateAndGetStockByID FAILED. Expected error %v got error %v", nil, err.Error())
	} else {
		t.Logf("TestStoreUpdateAndGetStockByID PASSED. Expected error %v got error %v", nil, err)
	}

	stock := entity.StockEntity{
		ID: &newID,
	}
	err = Serv.GetStockByIDService(context.Background(), &stock)

	// check equal value
	assert.Equal(t, newStock, stock)

	if err != nil {
		t.Errorf("TestStoreUpdateAndGetStockByID FAILED. Expected error %v got error %v", nil, err.Error())
	} else {
		t.Logf("TestStoreUpdateAndGetStockByID PASSED. Expected error %v got error %v", nil, err)
	}
}
