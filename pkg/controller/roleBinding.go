package controller

import (
	"kube-rbac-manager/pkg/kuberbac"
	"kube-rbac-manager/pkg/logger"

	"github.com/gin-gonic/gin"
)

type RoleBindingController struct{}

func (rbCtl *RoleBindingController) Create(c *gin.Context) {
	err := kuberbac.RbacOperator.RoleBinding.Create()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rbCtl *RoleBindingController) List(c *gin.Context) {
	err := kuberbac.RbacOperator.RoleBinding.List()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rbCtl *RoleBindingController) Delete(c *gin.Context) {
	err := kuberbac.RbacOperator.RoleBinding.Delete()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rbCtl *RoleBindingController) Update(c *gin.Context) {
	err := kuberbac.RbacOperator.RoleBinding.Update()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}
