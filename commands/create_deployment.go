package commands

import (
	"fmt"

	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // // Import solely to initialize client auth plugins.
)

// CreateDeployment creates a deployment
// with a POD
type CreateDeployment struct {
	Name           string
	Namespace      string
	NumReplicats   int32
	ContainerImage string
	ContainerTag   string
	ContainerPort  int32
	MemoryRequest  string
	CPURequest     string
	MemoryLimit    string
	CPULimit       string
	EnvVars        []apiv1.EnvVar
	Labels         map[string]string
	Clientset      *kubernetes.Clientset
	Logger         logrus.FieldLogger
}

// Execute run the command
func (c *CreateDeployment) Execute() error {
	deploymentsClient := c.Clientset.AppsV1().Deployments(c.Namespace)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-deployment", c.Name),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &c.NumReplicats,
			Selector: &metav1.LabelSelector{
				MatchLabels: c.Labels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: c.Labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  c.Name,
							Image: fmt.Sprintf("%s:%s", c.ContainerImage, c.ContainerTag),
							Env:   c.EnvVars,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: c.ContainerPort,
								},
							},
							Resources: getResourceRequirements(getResourceList(c.CPURequest, c.MemoryRequest), getResourceList(c.CPULimit, c.MemoryLimit)),
						},
					},
				},
			},
		},
	}
	_, err := deploymentsClient.Create(deployment)
	if err != nil {
		return err
	}
	c.Logger.WithField("name", c.Name).Info("deployment created")
	return nil
}

func getResourceRequirements(requests, limits apiv1.ResourceList) apiv1.ResourceRequirements {
	res := apiv1.ResourceRequirements{}
	res.Requests = requests
	res.Limits = limits
	return res
}

func getResourceList(cpu, memory string) apiv1.ResourceList {
	res := apiv1.ResourceList{}
	if cpu != "" {
		res[apiv1.ResourceCPU] = resource.MustParse(cpu)
	}
	if memory != "" {
		res[apiv1.ResourceMemory] = resource.MustParse(memory)
	}
	return res
}
