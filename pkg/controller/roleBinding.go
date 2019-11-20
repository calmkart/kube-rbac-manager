package controller

import (
	"github.com/gin-gonic/gin"
	"kube-rbac-manager/pkg/logger"
)

type RoleController struct{}

func (rCtl *RoleController) CreateRole(c *gin.Context) {
	// logger.Log.Error(err.Error())
	logger.Log.Info("test create here")
	return
}

func (rCtl *RoleController) ListRole(c *gin.Context) {
	return
}

func (rCtl *RoleController) DeleteRole(c *gin.Context) {
	return
}
