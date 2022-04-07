package inappmemory

import (
	"github.com/Calmantara/go-inventory-poc/entity"
)

type InAppMemory struct {
	Stocks         map[string]entity.StockEntity
	UrlPermissions map[string][]entity.UrlPermission
}

func NewInAppMemory() InAppMemory {
	// define seeds for role
	return InAppMemory{
		Stocks: make(map[string]entity.StockEntity),
		UrlPermissions: map[string][]entity.UrlPermission{
			"role1": {
				entity.UrlPermission{
					Path:   "/stock",
					Method: "POST",
				},
				entity.UrlPermission{
					Path:   "/stock",
					Method: "PUT",
				},
				entity.UrlPermission{
					Path:   "/stock/",
					Method: "GET",
				},
			},
			"role2": {
				entity.UrlPermission{
					Path:   "/stock",
					Method: "PUT",
				},
				entity.UrlPermission{
					Path:   "/stock/:id",
					Method: "GET",
				},
			},
			"role3": {
				entity.UrlPermission{
					Path:   "/stock/:id",
					Method: "GET",
				},
			},
		},
	}
}

func (i *InAppMemory) StoreStock(stock *entity.StockEntity) (err error) {
	i.Stocks[stock.ID.String()] = *stock
	return err
}
