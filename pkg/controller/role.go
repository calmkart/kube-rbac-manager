package controller

import (
	"github.com/gin-gonic/gin"
	"kube-rbac-manager/pkg/kuberbac"
	"kube-rbac-manager/pkg/logger"
)

type RoleController struct{}

func (rCtl *RoleController) CreateRole(c *gin.Context) {
	err := kuberbac.RbacOperator.CreateRole()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *RoleController) ListRole(c *gin.Context) {
	err := kuberbac.RbacOperator.ListRole()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *RoleController) DeleteRole(c *gin.Context) {
	err := kuberbac.RbacOperator.DeleteRole()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *RoleController) UpdateRole(c *gin.Context) {
	err := kuberbac.RbacOperator.UpdateRole()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}
