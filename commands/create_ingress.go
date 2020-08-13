package commands

import (
	"fmt"

	"github.com/sirupsen/logrus"
	extv1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // // Import solely to initialize client auth plugins.
)

// CreateIngressCommand creates an ingress for
// a service
type CreateIngressCommand struct {
	Name        string
	Namespace   string
	Host        string
	Path        string
	ServicePort int
	IngressPath string
	Labels      map[string]string
	Annotations map[string]string
	Clientset   *kubernetes.Clientset
	Logger      logrus.FieldLogger
}

// Execute run the command
func (c *CreateIngressCommand) Execute() error {
	_, err := c.Clientset.ExtensionsV1beta1().Ingresses(c.Namespace).Create(&extv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:        c.Name,
			Namespace:   c.Namespace,
			Labels:      c.Labels,
			Annotations: c.Annotations,
		},
		Spec: extv1.IngressSpec{
			Rules: []extv1.IngressRule{
				{
					Host: fmt.Sprintf("%s.%s", c.Namespace, c.Host),
					IngressRuleValue: extv1.IngressRuleValue{
						HTTP: &extv1.HTTPIngressRuleValue{
							Paths: []extv1.HTTPIngressPath{
								{
									Path: c.Path,
									Backend: extv1.IngressBackend{
										ServiceName: c.Name,
										ServicePort: intstr.FromInt(c.ServicePort),
									},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}
	c.Logger.WithField("name", c.Name).Info("ingress created")
	return nil
}
