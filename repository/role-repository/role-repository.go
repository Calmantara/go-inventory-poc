package rolerepository

import "github.com/Calmantara/go-inventory-poc/entity"

type RoleRepository interface {
	GetRolePermission(role string, permission *[]entity.UrlPermission) (err error)
}
