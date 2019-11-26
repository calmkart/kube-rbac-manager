package main

import (
	"kube-rbac-manager/pkg/kuberbac"
	"kube-rbac-manager/pkg/logger"
	"kube-rbac-manager/pkg/router"

	"github.com/gin-contrib/cors"
)

func main() {

	// init logger
	logger.InitLogger()

	// create a KubeOpts
	err := kuberbac.NewKubeOpts(nil, nil)
	if err != nil {
		logger.Log.Errorf("Create kubeOpts error, err: %s", err.Error())
		return
	}
	// start rbac list-watch listen
	go kuberbac.RbacOperator.StartListen()

	// create gin router
	router := router.GetRouter()

	// allow cors
	router.Use(cors.Default())

	// run
	router.Run(":8080")
}
