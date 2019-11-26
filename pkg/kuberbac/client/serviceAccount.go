package client

import "k8s.io/client-go/kubernetes"

/*
	we needn't a List method to request k8s-server here.
	because we list-watch listen to the k8s api-server, and stored it in cache.
	so if we want to list the resouces, just use cache.

*/

type ServiceAccountInterface interface {
	Create() error
	Update() error
	Delete() error
	List() error
}

type ServiceAccountClient struct {
	ClientSet *kubernetes.Clientset
}

func NewServiceAccountClient(clientset *kubernetes.Clientset) *ServiceAccountClient {
	return &ServiceAccountClient{
		ClientSet: clientset,
	}
}

func (sa *ServiceAccountClient) Create() error {
	return nil
}

func (sa *ServiceAccountClient) Delete() error {
	return nil
}

func (sa *ServiceAccountClient) Update() error {
	return nil
}

func (sa *ServiceAccountClient) List() error {
	// just read the cache is enough
	return nil
}
