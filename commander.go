package orchestrator

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

// Cooker is a recipe executor
type Cooker struct{}

// NewCommander factory method
func NewCommander() *Cooker {
	return &Cooker{}
}

// Execute will execute the set of commands of a recipe
func (*Cooker) Execute(r Recipe) error {
	for _, c := range r.Build() {
		if err := c.Execute(); err != nil {
			return err
		}
	}
	return nil
}
