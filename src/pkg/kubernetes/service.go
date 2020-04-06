package kubernetes

import (
	"fmt"
	"log"

	k8s_go "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc" // xxx
	k8s_clientcmd "k8s.io/client-go/tools/clientcmd"
)

// Kubernetes ...
type Kubernetes struct {
	clientset *k8s_go.Clientset
	Namespace string
}

func (k8s *Kubernetes) init() {
	if k8s.clientset != nil {
		return
	}

	kubeconfig := k8s_clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		k8s_clientcmd.NewDefaultClientConfigLoadingRules(),
		&k8s_clientcmd.ConfigOverrides{},
	)

	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("#1563220691: %v", err))
		return
	}

	k8s.clientset, err = k8s_go.NewForConfig(restconfig)
	if err != nil {
		log.Fatal(fmt.Sprintf("#1563220710: %v", err))
		return
	}
}
