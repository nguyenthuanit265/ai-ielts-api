package middleware

type ApiMiddleware struct {
	//roleService           services.RoleService
	//userService           services.UserService
	//rolePermissionService services.RolePermissionService
}

func NewMiddleware(
//roleService services.RoleService,
//userService services.UserService,
//rolePermissionService services.RolePermissionService,
) *ApiMiddleware {

	return &ApiMiddleware{
		//roleService:           roleService,
		//userService:           userService,
		//rolePermissionService: rolePermissionService,
	}
}
