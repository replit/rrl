package rrl

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// PodCounter counts the number of healthy pods running.
type PodCounter interface {
	PodCount(ctx context.Context) (int64, error)
}

// KubernetesClient is a wrapper over the kubernetes client used to
// communicate with the kubernetes API.
type KubernetesClient struct {
	clientset *kubernetes.Clientset
	namespace string
	app       string
}

// NewKubernetesClient creates a new kubernetes client using in-cluster authn/authz.
func NewKubernetesClient(namespace, app string) (*KubernetesClient, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("could not get kubernetes client config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("could not create kubernetes client: %w", err)
	}

	return &KubernetesClient{
		clientset: clientset,
		namespace: namespace,
		app:       app,
	}, nil
}

// PodCount returns the number of healthy pods.
func (c *KubernetesClient) PodCount(ctx context.Context) (int64, error) {
	pods, err := c.clientset.CoreV1().Pods(c.namespace).List(ctx, metav1.ListOptions{
		FieldSelector: "status.phase=Running",
		LabelSelector: fmt.Sprintf("app=%s", c.app),
	})
	if err != nil {
		return 0, err
	}

	ready := int64(0)
	for _, pod := range pods.Items {
		for _, cond := range pod.Status.Conditions {
			if cond.Type == corev1.PodReady && cond.Status == corev1.ConditionTrue {
				ready++
				break
			}
		}
	}

	return ready, nil
}
