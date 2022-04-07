package roleservice

import "context"

type RoleService interface {
	GetRolePermissionService(ctx context.Context, role string) bool
}
