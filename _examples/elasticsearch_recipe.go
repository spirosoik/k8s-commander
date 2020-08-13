package main

import (
	"github.com/sirupsen/logrus"
	commander "github.com/spirosoik/k8s-commander"
	"github.com/spirosoik/k8s-commander/commands"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// Opts options for the recipe
type recipeOpts struct {
	Name           string
	Namespace      string
	ContainerImage string
	ContainerTag   string
	ContainerPort  int32
}

type elasticSearchRecipe struct {
	options   recipeOpts
	clientset *kubernetes.Clientset
	logger    logrus.FieldLogger
}

// NewElasticsearchDeployment creates a recipe for the provided
// deployment request to create Elasticsearch Deployment and service
func NewElasticsearchDeployment(options recipeOpts, clientset *kubernetes.Clientset, logger logrus.FieldLogger) commander.Recipe {
	return &elasticSearchRecipe{
		options:   options,
		clientset: clientset,
		logger:    logger,
	}
}

func (d *elasticSearchRecipe) Build() []commander.Command {
	labels := map[string]string{
		"app": d.options.Name,
	}
	cmds := make([]commander.Command, 0)
	return append(cmds,
		&commands.CreateDeployment{
			Name:           d.options.Name,
			Namespace:      d.options.Namespace,
			Labels:         labels,
			NumReplicats:   1,
			ContainerImage: d.options.ContainerImage,
			ContainerTag:   d.options.ContainerTag,
			ContainerPort:  d.options.ContainerPort,
			Clientset:      d.clientset,
			Logger:         d.logger,
		},
		&commands.CreateService{
			Name:      d.options.Name,
			Namespace: d.options.Namespace,
			ServicePort: []apiv1.ServicePort{
				{
					Port:     d.options.ContainerPort,
					Protocol: apiv1.ProtocolTCP,
				},
			},
			Labels:    labels,
			Clientset: d.clientset,
			Logger:    d.logger,
		},
	)
}
