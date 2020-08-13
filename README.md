
[![PkgGoDev](https://pkg.go.dev/badge/github.com/spirosoik/k8s-commander)](https://pkg.go.dev/github.com/spirosoik/k8s-commander)
[![report card](https://img.shields.io/badge/report%20card-a%2B-ff3333.svg?style=flat-square)](http://goreportcard.com/report/spirosoik/k8s-commander)

# k8s-commander

A golang library to create Kubernetes recipes following the command design pattern.

## Recipe Example

You can find an example of a k8s recipe [here](_examples/elasticsearch_recipe.go)

```bash
make run-example
```

Check [here](_examples/main.go) how to run recipe with the k8s commander. Example:

```go
	cm := commander.New()

	// create a recipe with a set of commands
	opts := recipeOpts{
		Name:           "es",
		Namespace:      "default",
		ContainerImage: "elastic/elasticsearch",
		ContainerPort:  9200,
		ContainerTag:   "7.8.1",
	}
	recipe := NewElasticsearchDeployment(opts, clientset, logger)

	// Execute the recipe
	err = cm.Execute(recipe)
	if err != nil {
		logger.WithError(err).Error()
		os.Exit(-1)
	}
```

## Available Commands

- CreateNamespace
- DeleteNamespace
- CreateDeployment
- CreateIngressCommand
- CreateService