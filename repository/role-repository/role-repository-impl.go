package rolerepository

import (
	"errors"

	inappmemory "github.com/Calmantara/go-inventory-poc/configuration/in-app-memory"
	"github.com/Calmantara/go-inventory-poc/entity"
)

type RoleRepositoryImpl struct {
	inAppMemory inappmemory.InAppMemory
}

func (r *RoleRepositoryImpl) GetRolePermission(role string, permission *[]entity.UrlPermission) (err error) {
	*permission = r.inAppMemory.UrlPermissions[role]
	if permission == nil {
		err = errors.New("role is not found")
		return err
	}
	return err
}
