package roleservice

import (
	"context"

	"github.com/Calmantara/go-inventory-poc/entity"
	rolerepository "github.com/Calmantara/go-inventory-poc/repository/role-repository"
	"go.uber.org/zap"
)

type RoleServiceImpl struct {
	sugar          *zap.SugaredLogger
	roleRepository rolerepository.RoleRepository
}

func NewRoleService(
	sugar *zap.SugaredLogger,
	roleRepository rolerepository.RoleRepository,
) RoleService {
	return &RoleServiceImpl{
		sugar:          sugar,
		roleRepository: roleRepository,
	}
}

func (r *RoleServiceImpl) GetRolePermissionService(ctx context.Context, role string) bool {
	r.sugar.Infof("checking role permission:%v", role)

	// check in repository
	var permissions []entity.UrlPermission
	if err := r.roleRepository.GetRolePermission(role, &permissions); err != nil {
		return false
	}

	// iterating user role
	//TODO check value for URL and METHOD
	return true
}
