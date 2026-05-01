package main

import (
	"strings"
	"sync"

	"github.com/9Ashwin/agents-md-demo/tools"
)

// Agent is the core AI agent that processes queries using registered tools.
type Agent struct {
	tools     []tools.Tool
	verbose   bool
	mu        sync.Mutex
	execCount int
}

// Stats holds agent execution statistics.
type Stats struct {
	Tools    int
	Executed int
}

// NewAgent creates a new Agent with default tools.
func NewAgent() *Agent {
	return &Agent{
		tools: []tools.Tool{
			{Name: "web_search", Description: "Search the web for information"},
			{Name: "read_file", Description: "Read a file from disk"},
			{Name: "write_file", Description: "Write content to a file"},
			{Name: "exec_cmd", Description: "Execute a shell command"},
			{Name: "send_message", Description: "Send a message to the user"},
		},
	}
}

// SetVerbose enables or disables verbose logging.
func (a *Agent) SetVerbose(v bool) {
	a.verbose = v
}

// ListTools returns a copy of the registered tools.
func (a *Agent) ListTools() []tools.Tool {
	result := make([]tools.Tool, len(a.tools))
	copy(result, a.tools)
	return result
}

// Stats returns current agent statistics.
func (a *Agent) Stats() Stats {
	a.mu.Lock()
	defer a.mu.Unlock()
	return Stats{Tools: len(a.tools), Executed: a.execCount}
}

// Run executes the agent with the given query words.
func (a *Agent) Run(args []string) string {
	if len(args) == 0 {
		return "no query provided"
	}

	query := strings.Join(args, " ")

	a.mu.Lock()
	a.execCount++
	a.mu.Unlock()

	if a.verbose {
		return "Agent processing: " + query + " ... done"
	}
	return "Agent: processed '" + query + "'"
}
