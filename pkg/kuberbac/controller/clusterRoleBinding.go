package controller

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"kube-rbac-manager/pkg/logger"
)


func BuildClusterRoleBindingsController(clientSet *kubernetes.Clientset) (controller cache.Controller) {
	watchlist := cache.NewListWatchFromClient(clientSet.RbacV1().RESTClient(), "clusterRoleBindings", v1.NamespaceAll, fields.Everything())
	_, controller = cache.NewInformer(watchlist, &v1.Service{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc:    AddClusterRoleBindingsHandler,
		DeleteFunc: deleteClusterRoleBindingsHandler,
		UpdateFunc: updateClusterRoleBindingsHandler,
	},
	)
	return
}

func AddClusterRoleBindingsHandler(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}
	svcNamespace, svcName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	logger.Log.Debugf("Add ClusterRoleBinding %s namespace %s.", svcName, svcNamespace)
}

func deleteClusterRoleBindingsHandler(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}
	svcNamespace, svcName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	logger.Log.Debugf("Delete ClusterRoleBinding %s namespace %s.", svcName, svcNamespace)

}

func updateClusterRoleBindingsHandler(old interface{}, new interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(new)
	if err == nil {
		logger.Log.Debugf("update ClusterRoleBinding %s.", key)
	}
}
