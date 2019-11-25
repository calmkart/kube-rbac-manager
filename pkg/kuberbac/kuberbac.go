package kuberbac

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// we can use the kubeRbacOperator to do something about kubenetes
// such as list-watch listen, create/delete/update resouces, and so on.
var RbacOperator *KubeOpts

type KubeOpts struct {
	ClientSet *kubernetes.Clientset
	// do some work here
}

func NewKubeOpts(configBytes []byte, kubeconfig string) error {
	var restClient *rest.Config
	var err error
	if len(configBytes) > 0 {
		restClient, err = clientcmd.RESTConfigFromKubeConfig(configBytes)
		if err != nil {
			return err
		}
	} else {
		restClient, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return err
		}
	}
	// build clientset
	clientset, err := kubernetes.NewForConfig(restClient)
	if err != nil {
		return err
	}

	RbacOperator = &KubeOpts{
		//
		ClientSet: clientset,
	}
	return nil
}

func (kOpts *KubeOpts) StartListen() {
	// start list-watch the account/role/rolebinding status
	// and write to the cache
	// if we want to get the account/role/rolebinding list
	// we can just list the cache
}
