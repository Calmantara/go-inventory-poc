package stockrepository

import (
	"context"
	"testing"

	inappmemory "github.com/Calmantara/go-inventory-poc/configuration/in-app-memory"
	"github.com/Calmantara/go-inventory-poc/entity"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

var (
	Iam  = inappmemory.NewInAppMemory()
	Repo = NewStockRepository(Iam)
)

func TestGetAllStocks(t *testing.T) {
	var stocks map[string]entity.StockEntity
	err := Repo.GetAllStocks(context.Background(), &stocks)

	if err != nil {
		t.Errorf("TestGetAllStocks FAILED. Expected error %v got error %v", nil, err.Error())
	} else {
		t.Logf("TestGetAllStocks PASSED. Expected error %v got error %v", nil, err)
	}
}

func TestGetStockByID(t *testing.T) {
	newID := uuid.New()
	newStock := entity.StockEntity{
		ID:           &newID,
		Name:         "stock-1",
		Price:        50000,
		Availability: 100,
		IsActive:     true,
	}
	err := Repo.StoreStock(context.Background(), &newStock)
	if err != nil {
		t.Errorf("TestGetStockByID FAILED. Expected error %v got error %v", nil, err.Error())
	} else {
		t.Logf("TestGetStockByID PASSED. Expected error %v got error %v", nil, err)
	}

	stock := entity.StockEntity{
		ID: &newID,
	}
	err = Repo.GetStockByID(context.Background(), &stock)

	// check equal value
	assert.Equal(t, newStock, stock)

	if err != nil {
		t.Errorf("TestGetStockByID FAILED. Expected error %v got error %v", nil, err.Error())
	} else {
		t.Logf("TestGetStockByID PASSED. Expected error %v got error %v", nil, err)
	}
}

func TestStoreStock(t *testing.T) {
	newID := uuid.New()
	newStock := entity.StockEntity{
		ID:           &newID,
		Name:         "stock-1",
		Price:        50000,
		Availability: 100,
		IsActive:     true,
	}
	err := Repo.StoreStock(context.Background(), &newStock)
	if err != nil {
		t.Errorf("TestGetStockByID FAILED. Expected error %v got error %v", nil, err.Error())
	} else {
		t.Logf("TestGetStockByID PASSED. Expected error %v got error %v", nil, err)
	}
}
