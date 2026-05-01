package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: agent-demo <command> [args...]")
		fmt.Println("Commands: run, version, stats, tools")
		return
	}

	agent := NewAgent()

	switch args[0] {
	case "run":
		query := strings.Join(args[1:], " ")
		result := agent.Run(strings.Fields(query))
		fmt.Println(result)
	case "version":
		fmt.Println(Version())
	case "stats":
		stats := agent.Stats()
		fmt.Printf("Tools: %d, Executed: %d\n", stats.Tools, stats.Executed)
	case "tools":
		for _, t := range agent.ListTools() {
			fmt.Printf("  %s - %s\n", t.Name, t.Description)
		}
	case "verbose":
		if len(args) < 2 {
			fmt.Println("verbose requires a subcommand")
			return
		}
		agent.SetVerbose(true)
		query := strings.Join(args[1:], " ")
		result := agent.Run(strings.Fields(query))
		fmt.Println(result)
	default:
		fmt.Printf("unknown command: %s\n", args[0])
	}
}

// Version returns the agent version.
func Version() string {
	return "0.2.0"
}
