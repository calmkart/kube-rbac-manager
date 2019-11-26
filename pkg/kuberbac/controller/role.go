package controller

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)


func BuildRolesController(clientSet *kubernetes.Clientset) (controller cache.Controller) {
	watchlist := cache.NewListWatchFromClient(kOpts.ClientSet.RbacV1().RESTClient(), "roles", v1.NamespaceAll, fields.Everything())
	_, controller = cache.NewInformer(watchlist, &v1.Service{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc:    addRolesHandler,
		DeleteFunc: deleteRolesHandler,
		UpdateFunc: updateRolesHandler,
	},
	)
	return
}

func addRolesHandler(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}
	svcNamespace, svcName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	logger.Log.Debugf("Add Role %s namespace %s.", svcName, svcNamespace)
}

func deleteRolesHandler(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}
	svcNamespace, svcName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	logger.Log.Debugf("Delete Role %s namespace %s.", svcName, svcNamespace)

}

func updateRolesHandler(old interface{}, new interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(new)
	if err == nil {
		logger.Log.Debugf("update Role %s.", key)
	}
}
