package kuberbac

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"kube-rbac-manager/pkg/kuberbac/client"
	"kube-rbac-manager/pkg/kuberbac/controller"
	"os"
)

// we can use the kubeRbacOperator to do something about kubenetes
// such as list-watch listen, create/delete/update resouces, and so on.
var RbacOperator *KubeOpts

type KubeOpts struct {
	ClientSet          *kubernetes.Clientset
	ServiceAccount     client.ServiceAccountInterface
	Role               client.RoleInterface
	RoleBinding        client.RoleBindingInterface
	ClusterRole        client.ClusterRoleInterface
	ClusterRoleBinding client.ClusterRoleBindingInterface
	// do some work here
}

func NewKubeOpts(configBytes []byte, kubeconfig *string) error {
	var restClient *rest.Config
	var err error
	// build restClient
	if len(configBytes) > 0 {
		restClient, err = clientcmd.RESTConfigFromKubeConfig(configBytes)
		if err != nil {
			return err
		}
	} else {
		restClient, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			return err
		}
	}
	// build clientset
	clientset, err := kubernetes.NewForConfig(restClient)
	if err != nil {
		return err
	}
	// build some sub clients
	serviceAccountClient := client.NewServiceAccountClient(clientset)
	roleClient := client.NewRoleClient(clientset)
	roleBindingClient := client.NewRoleBindingClient(clientset)
	clusterRoleClient := client.NewClusterRoleClient(clientset)
	clusterRoleBindingClient := client.NewClusterRoleBindingClient(clientset)

	// build a Rbacoperator
	RbacOperator = &KubeOpts{
		ClientSet:          clientset,
		ServiceAccount:     serviceAccountClient,
		Role:               roleClient,
		RoleBinding:        roleBindingClient,
		ClusterRole:        clusterRoleClient,
		ClusterRoleBinding: clusterRoleBindingClient,
	}
	return nil
}

func (kOpts *KubeOpts) StartListen() {
	// start list-watch the account/role/rolebinding/clusterRole/clusterRoleBinding status
	// and write to the cache
	// if we want to get the account/role/rolebinding/clusterRole/clusterRoleBinding list
	// we can just list the cache

	// build controllers
	serviceAccountsController := controller.BuildServiceAccountsController(kOpts.ClientSet)
	roleBindingController := controller.BuildRoleBindingsController(kOpts.ClientSet)
	roleController := controller.BuildRolesController(kOpts.ClientSet)
	clusterRoleController := controller.BuildClusterRolesController(kOpts.ClientSet)
	clusterRoleBindingController := controller.BuildClusterRoleBindingsController(kOpts.ClientSet)


	// run controllers
	stopChannel := make(chan struct{})
	signals := make(chan os.Signal, 1)
	defer close(stopChannel)

	go serviceAccountsController.Run(stopChannel)
	go roleBindingController.Run(stopChannel)
	go roleController.Run(stopChannel)
	go clusterRoleController.Run(stopChannel)
	go clusterRoleBindingController.Run(stopChannel)

	<- signals
}