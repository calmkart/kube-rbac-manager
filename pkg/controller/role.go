package controller

import (
	"github.com/gin-gonic/gin"
	"kube-rbac-manager/pkg/kuberbac"
	"kube-rbac-manager/pkg/logger"
)

type RoleController struct{}

func (rCtl *RoleController) Create(c *gin.Context) {
	err := kuberbac.RbacOperator.Role.Create()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *RoleController) List(c *gin.Context) {
	err := kuberbac.RbacOperator.Role.List()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *RoleController) Delete(c *gin.Context) {
	err := kuberbac.RbacOperator.Role.Delete()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *RoleController) Update(c *gin.Context) {
	err := kuberbac.RbacOperator.Role.Update()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}
