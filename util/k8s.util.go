package tyrk8s

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetClient returns the kubernetes client based on the current env
func GetClient() (*kubernetes.Clientset, error) {
	env := os.Getenv("ENV")
	var config *rest.Config
	var err error
	if env == "production" {
		fmt.Println("Getting K8S Config In-Cluster")
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("Using K8S Config Outside Cluster")
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}
