package router

import (
	"kube-rbac-manager/pkg/controller"

	"github.com/gin-gonic/gin"
)

var serviceAccountCtl = new(controller.ServiceAccountController)
var roleCtl = new(controller.RoleController)
var roleBindingCtl = new(controller.RoleBindingController)
var clusterRoleCtl = new(controller.ClusterRoleController)
var clusterRoleBindingCtl = new(controller.ClusterRoleBindingController)

// GetRouter gen gin router
func GetRouter() *gin.Engine {
	router := gin.Default()

	// Account router
	account := router.Group("/api/serviceAccount")
	{
		account.POST("/create", serviceAccountCtl.Create)
		account.GET("/List", serviceAccountCtl.List)
		account.POST("/delete", serviceAccountCtl.Delete)
		account.POST("/update", serviceAccountCtl.Update)
	}

	// Role router
	role := router.Group("/api/role")
	{
		role.POST("/create", roleCtl.Create)
		role.GET("/List", roleCtl.List)
		role.POST("/delete", roleCtl.Delete)
		role.POST("/update", roleCtl.Update)
	}

	// RoleBinding router
	roleBinding := router.Group("/api/roleBinding")
	{
		roleBinding.POST("/create", roleBindingCtl.Create)
		roleBinding.GET("/List", roleBindingCtl.List)
		roleBinding.POST("/delete", roleBindingCtl.Delete)
		roleBinding.POST("/update", roleBindingCtl.Update)
	}

	// ClusterRole router
	clusterRole := router.Group("/api/clusterRole")
	{
		clusterRole.POST("/create", clusterRoleCtl.Create)
		clusterRole.GET("/List", clusterRoleCtl.List)
		clusterRole.POST("/delete", clusterRoleCtl.Delete)
		clusterRole.POST("/update", clusterRoleCtl.Update)
	}

	// ClusterRoleBinding router
	clusterRoleBinding := router.Group("/api/clusterRoleBinding")
	{
		clusterRoleBinding.POST("/create", clusterRoleBindingCtl.Create)
		clusterRoleBinding.GET("/List", clusterRoleBindingCtl.List)
		clusterRoleBinding.POST("/delete", clusterRoleBindingCtl.Delete)
		clusterRoleBinding.POST("/update", clusterRoleBindingCtl.Update)
	}


	return router
}
