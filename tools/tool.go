package tools

import "fmt"

// Tool represents a callable tool registered with the agent.
type Tool struct {
	Name        string
	Description string
	Execute     func(input map[string]any) (any, error)
}

// Validate checks whether the tool has required fields set.
func (t Tool) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("tool name is required")
	}
	if t.Execute == nil {
		return fmt.Errorf("tool %q has no execute handler", t.Name)
	}
	return nil
}
