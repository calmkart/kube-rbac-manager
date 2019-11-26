package controller

import (
	"github.com/gin-gonic/gin"
	"kube-rbac-manager/pkg/kuberbac"
	"kube-rbac-manager/pkg/logger"
)

type ClusterRoleBindingController struct{}

func (rCtl *ClusterRoleBindingController) Create(c *gin.Context) {
	err := kuberbac.RbacOperator.RoleBinding.Create()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *ClusterRoleBindingController) List(c *gin.Context) {
	err := kuberbac.RbacOperator.RoleBinding.List()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *ClusterRoleBindingController) Delete(c *gin.Context) {
	err := kuberbac.RbacOperator.RoleBinding.Delete()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}

func (rCtl *ClusterRoleBindingController) Update(c *gin.Context) {
	err := kuberbac.RbacOperator.RoleBinding.Update()
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, gin.H{})
}
