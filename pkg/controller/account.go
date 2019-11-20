package controller

import (
	"kube-rbac-manager/pkg/kuberbac"
	"kube-rbac-manager/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func (aCtl *AccountController) CreateAccount(c *gin.Context) {
	err := kuberbac.RbacOperator.CreateAccount()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (aCtl *AccountController) ListAccount(c *gin.Context) {
	err := kuberbac.RbacOperator.ListAccount()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (aCtl *AccountController) DeleteAccount(c *gin.Context) {
	err := kuberbac.RbacOperator.DeleteAccount()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (aCtl *AccountController) UpdateAccount(c *gin.Context) {
	err := kuberbac.RbacOperator.UpdateAccount()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}
