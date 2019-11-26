package controller

import (
	"kube-rbac-manager/pkg/kuberbac"
	"kube-rbac-manager/pkg/logger"

	"github.com/gin-gonic/gin"
)

type ServiceAccountController struct{}

func (aCtl *ServiceAccountController) Create(c *gin.Context) {
	err := kuberbac.RbacOperator.ServiceAccount.Create()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (aCtl *ServiceAccountController) List(c *gin.Context) {
	err := kuberbac.RbacOperator.ServiceAccount.List()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (aCtl *ServiceAccountController) Delete(c *gin.Context) {
	err := kuberbac.RbacOperator.ServiceAccount.Delete()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (aCtl *ServiceAccountController) Update(c *gin.Context) {
	err := kuberbac.RbacOperator.ServiceAccount.Update()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}
