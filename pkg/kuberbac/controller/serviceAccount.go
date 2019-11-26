package controller

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"kube-rbac-manager/pkg/logger"
)

func BuildServiceAccountsController(clientSet *kubernetes.Clientset) (controller cache.Controller) {
	watchlist := cache.NewListWatchFromClient(clientSet.CoreV1().RESTClient(), "serviceAccounts", v1.NamespaceAll, fields.Everything())
	_, controller = cache.NewInformer(watchlist, &v1.Service{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc:    addServiceAccountsHandler,
		DeleteFunc: deleteServiceAccountsHandler,
		UpdateFunc: updateServiceAccountsHandler,
	},
	)
	return
}

func addServiceAccountsHandler(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}
	svcNamespace, svcName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	logger.Log.Debugf("Add ServiceAccount %s namespace %s.", svcName, svcNamespace)
}

func deleteServiceAccountsHandler(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}
	svcNamespace, svcName, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	logger.Log.Debugf("Delete ServiceAccount %s namespace %s.", svcName, svcNamespace)

}

func updateServiceAccountsHandler(old interface{}, new interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(new)
	if err == nil {
		logger.Log.Debugf("update ServiceAccount %s.", key)
	}
}
