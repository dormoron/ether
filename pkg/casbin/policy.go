package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type AccessCasbin struct {
	enforcer *casbin.Enforcer
}

func (a *AccessCasbin) AddPermissionForUser(user, permission string) (bool, error) {
	return a.enforcer.AddPermissionForUser(user, permission)
}

func (a *AccessCasbin) AddPermissionsForUser(user string, permissions []string) (bool, error) {
	return a.enforcer.AddPermissionsForUser(user, permissions)
}

func (a *AccessCasbin) DeletePermissionForUser(user, permission string) (bool, error) {
	return a.enforcer.DeletePermissionForUser(user, permission)
}

func (a *AccessCasbin) DeletePermissionsForUser(user string) (bool, error) {
	return a.enforcer.DeletePermissionsForUser(user)
}

func (a *AccessCasbin) GetPermissionsForUser(user string) ([][]string, error) {
	return a.enforcer.GetPermissionsForUser(user)
}

func (a *AccessCasbin) HasPermissionForUser(user string, permission []string) (bool, error) {
	return a.enforcer.HasPermissionForUser(user, permission...)
}

func (a *AccessCasbin) GetImplicitRolesForUser(user string) ([]string, error) {
	return a.enforcer.GetImplicitRolesForUser(user)
}

func (a *AccessCasbin) GetImplicitUsersForRole(role string) ([]string, error) {
	return a.enforcer.GetImplicitUsersForRole(role)
}

func (a *AccessCasbin) GetImplicitPermissionsForUser(user string) ([]string, error) {
	return a.enforcer.GetImplicitRolesForUser(user)
}

func (a *AccessCasbin) GetImplicitUsersForPermission(permission string) ([]string, error) {
	return a.enforcer.GetImplicitUsersForPermission(permission)
}

func (a *AccessCasbin) GetImplicitResourcesForUser(user string) ([][]string, error) {
	return a.enforcer.GetImplicitResourcesForUser(user)
}

func (a *AccessCasbin) GetImplicitUsersForResource(resource string) ([][]string, error) {
	return a.enforcer.GetImplicitUsersForResource(resource)
}

func (a *AccessCasbin) AddPolicy(sub, obj, act string) (bool, error) {
	return a.enforcer.AddPolicy(sub, obj, act)
}

func (a *AccessCasbin) RemovePolicy(sub, obj, act string) (bool, error) {
	return a.enforcer.RemovePolicy(sub, obj, act)
}

func (a *AccessCasbin) UpdatePolicy(oldPolicy []string, newPolicy []string) (bool, error) {
	return a.enforcer.UpdatePolicy(oldPolicy, newPolicy)
}

func (a *AccessCasbin) HasPolicy(sub, obj, act string) (bool, error) {
	return a.enforcer.HasPolicy(sub, obj, act)
}

func (a *AccessCasbin) AddRoleForUser(user, role string) (bool, error) {
	return a.enforcer.AddRoleForUser(user, role)
}

func (a *AccessCasbin) AddRolesForUser(user string, role []string) (bool, error) {
	return a.enforcer.AddRolesForUser(user, role)
}

func (a *AccessCasbin) DeleteRoleForUser(user, role string) (bool, error) {
	return a.enforcer.DeleteRoleForUser(user, role)
}

func (a *AccessCasbin) DeleteRolesForUser(user string) (bool, error) {
	return a.enforcer.DeleteRolesForUser(user)
}

func (a *AccessCasbin) DeleteUser(user string) (bool, error) {
	return a.enforcer.DeleteUser(user)
}

func (a *AccessCasbin) DeleteRole(role string) (bool, error) {
	return a.enforcer.DeleteRole(role)
}

func (a *AccessCasbin) GetRolesForUser(user string) ([]string, error) {
	return a.enforcer.GetRolesForUser(user)
}

func (a *AccessCasbin) GetUsersForRole(role string) ([]string, error) {
	return a.enforcer.GetUsersForRole(role)
}

func (a *AccessCasbin) HasRoleForUser(user, role string) (bool, error) {
	return a.enforcer.HasRoleForUser(user, role)
}

func (a *AccessCasbin) UpdateRoleForUser(user, oldRole, newRole string) (bool, error) {
	// Delete old role
	_, err := a.DeleteRoleForUser(user, oldRole)
	if err != nil {
		return false, err
	}

	return a.AddRoleForUser(user, newRole)
}

func (a *AccessCasbin) UpdateRolesForUser(user string, oldRoles, newRoles []string) (bool, error) {
	success := true
	// Delete old roles
	for _, oldRole := range oldRoles {
		if _, err := a.DeleteRoleForUser(user, oldRole); err != nil {
			success = false
		}
	}
	if _, err := a.AddRolesForUser(user, newRoles); err != nil {
		success = false
	}

	return success, nil
}

func NewAccessCasbin(db *gorm.DB, modelPath string) (Access, error) {
	// 初始化 MySQL 适配器
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// 加载模型
	m, err := model.NewModelFromFile(modelPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load model: %w", err)
	}

	// 初始化 Enforcer
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, fmt.Errorf("failed to create enforcer: %w", err)
	}

	// 加载策略
	if err = e.LoadPolicy(); err != nil {
		return nil, fmt.Errorf("failed to load policy: %w", err)
	}

	// 启用 Auto-Save 和 Auto-Load，减少网络开销
	e.EnableAutoSave(true)

	return &AccessCasbin{
		enforcer: e,
	}, err
}
