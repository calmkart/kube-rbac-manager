package main

import (
	"kube-rbac-manager/pkg/logger"
	"kube-rbac-manager/pkg/router"
	"kube-rbac-manager/pkg/kuberbac"

	"github.com/gin-contrib/cors"
)

func main() {

	// init logger
	logger.InitLogger()

	// create a KubeOpts
	kOpts, err := kuberbac.NewKubeOpts()
	if err != nil {
		logger.Log.Errorf("Create kubeOpts error, err: %s", err.Error())
	}
	// start rbac list-watch listen
	go kOpts.startListen()

	// create gin router
	router := router.GetRouter()

	// allow cors
	router.Use(cors.Default())

	// run
	router.Run(":8080")
}
