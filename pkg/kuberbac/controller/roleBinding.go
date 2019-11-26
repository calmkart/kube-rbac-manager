package controller

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"kube-rbac-manager/pkg/logger"
)

func BuildRoleBindingsController(clientSet *kubernetes.Clientset) (controller cache.Controller) {
	watchlist := cache.NewListWatchFromClient(clientSet.RbacV1().RESTClient(), "roleBindings", v1.NamespaceAll, fields.Everything())
	_, controller = cache.NewInformer(watchlist, &v1.Service{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc:    addRoleBindingsHandler,
		DeleteFunc: deleteRoleBindingsHandler,
		UpdateFunc: updateRoleBindingsHandler,
	},
	)
	return
}


func addRoleBindingsHandler(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}
	svcNamespace, svcName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	logger.Log.Debugf("Add RoleBinding %s namespace %s.", svcName, svcNamespace)
}

func deleteRoleBindingsHandler(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}
	svcNamespace, svcName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	logger.Log.Debugf("Delete RoleBinding %s namespace %s.", svcName, svcNamespace)

}

func updateRoleBindingsHandler(old interface{}, new interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(new)
	if err == nil {
		logger.Log.Debugf("update RoleBinding %s.", key)
	}
}
