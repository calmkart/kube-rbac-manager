package client

import "k8s.io/client-go/kubernetes"

/*
	we needn't a List method to request k8s-server here.
	because we list-watch listen to the k8s api-server, and stored it in cache.
	so if we want to list the resouces, just use cache.

*/

type ClusterRoleInterface interface {
	Create() error
	Update() error
	Delete() error
	List() error
}

type ClusterRoleClient struct {
	ClientSet *kubernetes.Clientset
}

func NewClusterRoleClient(clientset *kubernetes.Clientset) *ClusterRoleClient {
	return &ClusterRoleClient{
		ClientSet: clientset,
	}
}

func (cr *ClusterRoleClient) Create() error {
	return nil
}

func (cr *ClusterRoleClient) Delete() error {
	return nil
}

func (cr *ClusterRoleClient) Update() error {
	return nil
}

func (cr *ClusterRoleClient) List() error {
	// just read the cache is enough
	return nil
}
