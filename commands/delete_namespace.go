package commands

import (
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // // Import solely to initialize client auth plugins.
)

// DeleteNamespace deletes the namespace
type DeleteNamespace struct {
	Namespace string
	Clientset *kubernetes.Clientset
	Logger    logrus.FieldLogger
}

// Execute run the command
func (c *DeleteNamespace) Execute() error {
	deletePolicy := metav1.DeletePropagationForeground
	if err := c.Clientset.CoreV1().Namespaces().Delete(c.Namespace, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return err
	}
	c.Logger.WithField("name", c.Namespace).Info("namespace deleted")
	return nil
}
