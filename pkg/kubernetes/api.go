package kubernetes

import (
	api_v1 "k8s.io/api/core/v1"
	v1_beta1 "k8s.io/api/extensions/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

type NamespacesOutput struct {
	Items []api_v1.Namespace `json:"items"`
}

func (k8s *Kubernetes) GetNamespaces() (NamespacesOutput, error) {
	k8s.init()
	output := NamespacesOutput{}
	namespaces, err := k8s.clientset.CoreV1().Namespaces().List(meta_v1.ListOptions{})
	if err != nil {
		return output, err
	}

	output.Items = namespaces.Items
	return output, nil
}

type IngressesOutput struct {
	Items []v1_beta1.Ingress `json:"items"`
}

func (k8s *Kubernetes) GetIngresses() (IngressesOutput, error) {
	k8s.init()
	output := IngressesOutput{}

	ingresses, err := k8s.clientset.ExtensionsV1beta1().Ingresses(k8s.Namespace).List(meta_v1.ListOptions{})
	if err != nil {
		return output, err
	}

	output.Items = ingresses.Items
	return output, nil
}
