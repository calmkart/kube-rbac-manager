package kuberbac

/*
	we needn't a List method to request k8s-server here.
	because we list-watch listen to the k8s api-server, and stored it in cache.
	so if we want to list the resouces, just use cache.

*/

func (kOpts *KubeOpts) CreateRole() error {
	return nil
}

func (kOpts *KubeOpts) DeleteRole() error {
	return nil
}

func (kOpts *KubeOpts) UpdateRole() error {
	return nil
}

func (kOpts *KubeOpts) ListRole() error {
	// just read the cache is enough
	return nil
}
