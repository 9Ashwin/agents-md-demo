package main

import (
	"strings"

	"github.com/9Ashwin/agents-md-demo/tools"
)

// Agent is the core AI agent.
type Agent struct {
	tools []tools.Tool
}

// NewAgent creates a new Agent with default tools.
func NewAgent() *Agent {
	return &Agent{
		tools: []tools.Tool{
			{Name: "web_search", Description: "Search the web"},
			{Name: "read_file", Description: "Read a file from disk"},
			{Name: "write_file", Description: "Write content to a file"},
		},
	}
}

// Run executes the agent with the given query.
func (a *Agent) Run(args []string) string {
	query := strings.Join(args, " ")
	_ = query
	return "Agent response"
}
