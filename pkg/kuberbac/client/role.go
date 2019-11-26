package client

import "k8s.io/client-go/kubernetes"

/*
	we needn't a List method to request k8s-server here.
	because we list-watch listen to the k8s api-server, and stored it in cache.
	so if we want to list the resouces, just use cache.

*/

type RoleInterface interface {
	Create() error
	Update() error
	Delete() error
	List() error
}

type RoleClient struct {
	ClientSet *kubernetes.Clientset
}

func NewRoleClient(clientset *kubernetes.Clientset) *RoleClient {
	return &RoleClient{
		ClientSet: clientset,
	}
}

func (r *RoleClient) Create() error {
	return nil
}

func (r *RoleClient) Delete() error {
	return nil
}

func (r *RoleClient) Update() error {
	return nil
}

func (r *RoleClient) List() error {
	// just read the cache is enough
	return nil
}
