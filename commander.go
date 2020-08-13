package commander

// Command describes a command
type Command interface {
	Execute() error
}

// Recipe will create the commands
// for the recipe execution
type Recipe interface {
	Build() []Command
}

// Executor executes recipes
type Executor interface {
	Execute(r Recipe) error
}

// Commander is a recipe executor
type Commander struct{}

// New factory method
func New() *Commander {
	return &Commander{}
}

// Execute will execute the set of commands of a recipe
func (*Commander) Execute(r Recipe) error {
	for _, c := range r.Build() {
		if err := c.Execute(); err != nil {
			return err
		}
	}
	return nil
}
