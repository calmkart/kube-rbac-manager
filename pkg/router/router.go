package router

import (
	"kube-rbac-manager/pkg/controller"

	"github.com/gin-gonic/gin"
)

var accountCtl = new(controller.AccountController)
var roleCtl = new(controller.RoleController)
var roleBindingCtl = new(controller.RoleBindingController)

// GetRouter gen gin router
func GetRouter() *gin.Engine {
	router := gin.Default()

	// Account router
	account := router.Group("/api/account")
	{
		account.POST("/createAccount", accountCtl.CreateAccount)
		account.GET("/ListAccount", accountCtl.ListAccount)
		account.POST("/deleteAccount", accountCtl.DeleteAccount)
		account.POST("/updateAccount", accountCtl.UpdateAccount)
	}

	// Role router
	role := router.Group("/api/role")
	{
		role.POST("/createRole", roleCtl.CreateRole)
		role.GET("/ListRole", roleCtl.ListRole)
		role.POST("/deleteRole", roleCtl.DeleteRole)
		role.POST("/updateRole", roleCtl.UpdateRole)
	}

	// RoleBinding router
	roleBinding := router.Group("/api/roleBinding")
	{
		roleBinding.POST("/createRoleBinding", roleBindingCtl.CreateRoleBinding)
		roleBinding.GET("/ListRoleBinding", roleBindingCtl.ListRoleBinding)
		roleBinding.POST("/deleteRoleBinding", roleBindingCtl.DeleteRoleBinding)
		roleBinding.POST("/updateRoleBinding", roleBindingCtl.UpdateRoleBinding)
	}

	return router
}
