package tyrk8s

import (
	"fmt"
	"os"
	"path/filepath"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // req for gcp auth
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetClient returns the kubernetes client based on the current env
func GetClient() (client *kubernetes.Clientset, err error) {
	env := os.Getenv("ENV")
	var config *rest.Config
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
	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// CreateJob creates a k8s job based on the submission id
func CreateJob(id string) (job string, err error) {
	falseVal := false //bc spec needs a *bool
	jobName := fmt.Sprintf("grader-job-%s", id)
	podName := fmt.Sprintf("grader-pod-%s", id)
	client, err := GetClient()
	if err != nil {
		return "", err
	}
	jobsClient := client.BatchV1().Jobs("default")
	batchJob := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:   jobName,
			Labels: make(map[string]string),
		},
		Spec: batchv1.JobSpec{
			// Optional: Parallelism:,
			// Optional: Completions:,
			// Optional: ActiveDeadlineSeconds:,
			// Optional: Selector:,
			// Optional: ManualSelector:,
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   jobName,
					Labels: make(map[string]string),
				},
				Spec: apiv1.PodSpec{
					InitContainers: []apiv1.Container{},
					Containers: []apiv1.Container{
						{
							Name:    podName,
							Image:   "perl",
							Command: []string{"perl", "-Mbignum=bpi", "-wle", "print bpi(2000)"},
							SecurityContext: &apiv1.SecurityContext{
								Privileged: &falseVal,
							},
							ImagePullPolicy: apiv1.PullPolicy(apiv1.PullIfNotPresent),
							Env:             []apiv1.EnvVar{},
							VolumeMounts:    []apiv1.VolumeMount{},
						},
					},
					RestartPolicy:    apiv1.RestartPolicyNever,
					Volumes:          []apiv1.Volume{},
					ImagePullSecrets: []apiv1.LocalObjectReference{},
				},
			},
		},
	}

	newJob, err := jobsClient.Create(batchJob)
	if err != nil {
		return "", err
	}

	return newJob.Name, nil
}
