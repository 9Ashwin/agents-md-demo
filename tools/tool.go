package tools

// Tool represents a callable tool for the agent.
type Tool struct {
	Name        string
	Description string
	Execute     func(input map[string]any) (any, error)
}
