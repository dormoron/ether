package casbin

type Access interface {
	Policy
	RBAC
	Permission
	Resources
}

type Policy interface {
	AddPolicy(sub, obj, act string) (bool, error)
	RemovePolicy(sub, obj, act string) (bool, error)
	UpdatePolicy(oldPolicy []string, newPolicy []string) (bool, error)
	HasPolicy(sub, obj, act string) (bool, error)
}

type RBAC interface {
	AddRoleForUser(user, role string) (bool, error)
	AddRolesForUser(user string, role []string) (bool, error)
	DeleteRoleForUser(user, role string) (bool, error)
	DeleteRolesForUser(user string) (bool, error)
	DeleteUser(user string) (bool, error)
	DeleteRole(role string) (bool, error)
	GetRolesForUser(user string) ([]string, error)
	GetUsersForRole(role string) ([]string, error)
	HasRoleForUser(user, role string) (bool, error)
	UpdateRoleForUser(user, oldRoles, newRoles string) (bool, error)
	UpdateRolesForUser(user string, oldRoles, newRoles []string) (bool, error)
}

type Permission interface {
	AddPermissionForUser(user, permission string) (bool, error)
	AddPermissionsForUser(user string, permissions []string) (bool, error)
	DeletePermissionForUser(user, permission string) (bool, error)
	DeletePermissionsForUser(user string) (bool, error)
	GetPermissionsForUser(user string) ([][]string, error)
	HasPermissionForUser(user string, permission []string) (bool, error)
	GetImplicitRolesForUser(user string) ([]string, error)
	GetImplicitUsersForRole(role string) ([]string, error)
	GetImplicitPermissionsForUser(user string) ([]string, error)
	GetImplicitUsersForPermission(permission string) ([]string, error)
}

type Resources interface {
	GetImplicitResourcesForUser(user string) ([][]string, error)
	GetImplicitUsersForResource(resource string) ([][]string, error)
}
