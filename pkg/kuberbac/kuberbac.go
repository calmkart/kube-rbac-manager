package kuberbac

import (
	"k8s.io/client-go/kubernetes"
)

// we can use the kubeRbacOperator to do something about kubenetes
// such as list-watch listen, create/delete/update resouces, and so on.
var RbacOperator *KubeOpts

type KubeOpts struct {
	ClientSet *kubernetes.Clientset
	// do some work here
}

func NewKubeOpts() error {
	//
	RbacOperator = &KubeOpts{
		//
	}
	return nil
}

func (kOpts *KubeOpts) StartListen() {
	// start list-watch the accout/role/rolebinding status
	// and write to the cache
	// if we want to get the accout/role/rolebinding list
	// we can just list the cache
}
