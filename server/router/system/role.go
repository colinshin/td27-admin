package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	roleRouter := Router.Group("role").Use(middleware.OperationRecord())
	roleApi := api.ApiGroupApp.SystemApiGroup.RoleApi
	{
		roleRouter.POST("getRoles", roleApi.GetRoles)
		roleRouter.POST("addRole", roleApi.AddRole)
		roleRouter.POST("deleteRole", roleApi.DeleteRole)
		roleRouter.POST("editRole", roleApi.EditRole)
		roleRouter.POST("editRoleMenu", roleApi.EditRoleMenu)
	}
	return roleRouter
}
