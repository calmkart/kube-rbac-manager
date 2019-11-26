package controller

import (
	"github.com/gin-gonic/gin"
	"kube-rbac-manager/pkg/kuberbac"
	"kube-rbac-manager/pkg/logger"
)

type ClusterRoleController struct{}

func (rCtl *ClusterRoleController) Create(c *gin.Context) {
	err := kuberbac.RbacOperator.ClusterRole.Create()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *ClusterRoleController) List(c *gin.Context) {
	err := kuberbac.RbacOperator.ClusterRole.List()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *ClusterRoleController) Delete(c *gin.Context) {
	err := kuberbac.RbacOperator.ClusterRole.Delete()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *ClusterRoleController) Update(c *gin.Context) {
	err := kuberbac.RbacOperator.ClusterRole.Update()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}
