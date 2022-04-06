package inappmemory

import (
	"github.com/Calmantara/go-inventory-poc/entity"
)

type InAppMemory interface {
	//stock interface
	GetStockData() map[string]entity.StockEntity
	StoreStockData(id string, stockEntity entity.StockEntity) map[string]entity.StockEntity

	// url permission interface
	GetUrlPermissionData() map[string][]entity.UrlPermission
}
