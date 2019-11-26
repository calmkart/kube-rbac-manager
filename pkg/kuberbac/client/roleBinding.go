package client

import "k8s.io/client-go/kubernetes"

/*
	we needn't a List method to request k8s-server here.
	because we list-watch listen to the k8s api-server, and stored it in cache.
	so if we want to list the resouces, just use cache.

*/

type RoleBindingInterface interface {
	Create() error
	Update() error
	Delete() error
	List() error
}

type RoleBindingClient struct {
	ClientSet *kubernetes.Clientset
}

func NewRoleBindingClient(clientset *kubernetes.Clientset) *RoleBindingClient {
	return &RoleBindingClient{
		ClientSet: clientset,
	}
}

func (rb *RoleBindingClient) Create() error {
	return nil
}

func (rb *RoleBindingClient) Delete() error {
	return nil
}

func (rb *RoleBindingClient) Update() error {
	return nil
}

func (rb *RoleBindingClient) List() error {
	// just read the cache is enough
	return nil
}
