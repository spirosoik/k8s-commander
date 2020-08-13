package commands

import (
	"strings"

	"github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // // Import solely to initialize client auth plugins.
)

// CreateNamespace creates a namespace
type CreateNamespace struct {
	Name      string
	Clientset *kubernetes.Clientset
	Logger    logrus.FieldLogger
}

// Execute run the command
func (c *CreateNamespace) Execute() error {
	spec := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: strings.ToLower(c.Name),
		},
	}
	_, err := c.Clientset.CoreV1().Namespaces().Create(spec)
	if err != nil {
		return err
	}
	c.Logger.WithField("name", c.Name).Info("namespace created")
	return nil
}
