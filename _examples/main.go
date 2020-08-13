package main

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	commander "github.com/spirosoik/k8s-commander"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	logger := logrus.New()
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		logger.WithError(err).Error()
		os.Exit(-1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.WithError(err).Error()
		os.Exit(-1)
	}

	// Create a new commander
	cmer := commander.New()

	// create a recipe with a set of commands
	opts := recipeOpts{
		Name:           "es",
		Namespace:      "default",
		ContainerImage: "elasticsearch",
		ContainerPort:  9200,
		ContainerTag:   "latest",
	}
	recipe := NewElasticsearchDeployment(opts, clientset, logger)

	// Execute the recipe
	err = cmer.Execute(recipe)
	if err != nil {
		logger.WithError(err).Error()
		os.Exit(-1)
	}
}
