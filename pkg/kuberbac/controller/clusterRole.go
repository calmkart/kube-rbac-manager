package controller

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"kube-rbac-manager/pkg/logger"
)


func BuildClusterRolesController(clientSet *kubernetes.Clientset) (controller cache.Controller) {
	watchlist := cache.NewListWatchFromClient(clientSet.RbacV1().RESTClient(), "clusterRoles", v1.NamespaceAll, fields.Everything())
	_, controller = cache.NewInformer(watchlist, &v1.Service{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc:    addClusterRolesHandler,
		DeleteFunc: deleteClusterRolesHandler,
		UpdateFunc: updateClusterRolesHandler,
	},
	)
	return
}

func addClusterRolesHandler(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}
	svcNamespace, svcName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	logger.Log.Debugf("Add ClusterRole %s namespace %s.", svcName, svcNamespace)
}

func deleteClusterRolesHandler(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}
	svcNamespace, svcName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	logger.Log.Debugf("Delete ClusterRole %s namespace %s.", svcName, svcNamespace)

}

func updateClusterRolesHandler(old interface{}, new interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(new)
	if err == nil {
		logger.Log.Debugf("update ClusterRole %s.", key)
	}
}
