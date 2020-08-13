# k8s-commander

A golang library to create Kubernetes recipes following the command design pattern.

## Recipe Example

You can find an example of a k8s recipe [here](_examples/elasticsearch_recipe.go)

```bash
make run-example
```

Check [here](`_examples/main.go`) how to run recipe with the k8s commander. Example:

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