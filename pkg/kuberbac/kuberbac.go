package kuberbac

import {
	"k8s.io/client-go/kubernetes"
}

type KubeOpts struct{
	ClientSet    *kubernetes.Clientset
	// do some work here
}

func NewKubeOpts() *KubeOpts, err {
	return KubeOpts{
		//
	}, nil
}

func (kOpts *KubeOpts) StartListen() {
	// start list-watch the accout/role/rolebinding status
	// and write to the cache
}
