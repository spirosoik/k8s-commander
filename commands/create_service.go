package commands

import (
	"github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // // Import solely to initialize client auth plugins.
)

// CreateService creates an service for
// a POD
type CreateService struct {
	Name        string
	Namespace   string
	ServicePort []apiv1.ServicePort
	ServiceType apiv1.ServiceType
	Labels      map[string]string
	Annotations map[string]string
	Clientset   *kubernetes.Clientset
	Logger      logrus.FieldLogger
}

// Execute run the
func (c *CreateService) Execute() error {
	_, err := c.Clientset.CoreV1().Services(c.Namespace).Create(&apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        c.Name,
			Namespace:   c.Namespace,
			Labels:      c.Labels,
			Annotations: c.Annotations,
		},
		Spec: apiv1.ServiceSpec{
			Ports:    c.ServicePort,
			Type:     c.ServiceType,
			Selector: c.Labels,
		},
	})
	if err != nil {
		return err
	}
	c.Logger.WithField("name", c.Name).Info("service create")
	return nil
}
