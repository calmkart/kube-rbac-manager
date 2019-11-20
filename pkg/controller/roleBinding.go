package controller

import (
	"kube-rbac-manager/pkg/kuberbac"
	"kube-rbac-manager/pkg/logger"

	"github.com/gin-gonic/gin"
)

type RoleBindingController struct{}

func (rbCtl *RoleBindingController) CreateRoleBinding(c *gin.Context) {
	err := kuberbac.RbacOperator.CreateRoleBinding()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rbCtl *RoleBindingController) ListRoleBinding(c *gin.Context) {
	err := kuberbac.RbacOperator.ListRoleBinding()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rbCtl *RoleBindingController) DeleteRoleBinding(c *gin.Context) {
	err := kuberbac.RbacOperator.DeleteRoleBinding()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rbCtl *RoleBindingController) UpdateRoleBinding(c *gin.Context) {
	err := kuberbac.RbacOperator.UpdateRoleBinding()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}
