package inappmemory

import (
	"github.com/Calmantara/go-inventory-poc/entity"
)

type InAppMemoryImpl struct {
	stocks         map[string]entity.StockEntity
	urlPermissions map[string][]entity.UrlPermission
}

func NewInAppMemory() InAppMemory {
	return &InAppMemoryImpl{}
}

//stock impl
func (i *InAppMemoryImpl) GetStockData() map[string]entity.StockEntity {
	//return all recorded stocks
	return i.stocks
}
func (i *InAppMemoryImpl) StoreStockData(id string, stockEntity entity.StockEntity) map[string]entity.StockEntity {
	// add stock to hash map
	i.stocks[id] = stockEntity
	return map[string]entity.StockEntity{id: stockEntity}
}

// url permission impl
func (i *InAppMemoryImpl) GetUrlPermissionData() map[string][]entity.UrlPermission {
	// return all url permission
	return i.urlPermissions
}
