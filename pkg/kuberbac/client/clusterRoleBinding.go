package client

import "k8s.io/client-go/kubernetes"

/*
	we needn't a List method to request k8s-server here.
	because we list-watch listen to the k8s api-server, and stored it in cache.
	so if we want to list the resouces, just use cache.

*/

type ClusterRoleBindingInterface interface {
	Create() error
	Update() error
	Delete() error
	List() error
}

type ClusterRoleBindingClient struct {
	ClientSet *kubernetes.Clientset
}

func NewClusterRoleBindingClient(clientset *kubernetes.Clientset) *ClusterRoleBindingClient {
	return &ClusterRoleBindingClient{
		ClientSet: clientset,
	}
}

func (crb *ClusterRoleBindingClient) Create() error {
	return nil
}

func (crb *ClusterRoleBindingClient) Delete() error {
	return nil
}

func (crb *ClusterRoleBindingClient) Update() error {
	return nil
}

func (crb *ClusterRoleBindingClient) List() error {
	// just read the cache is enough
	return nil
}
