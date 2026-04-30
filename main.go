package main

import (
	"fmt"
	"os"
)

func main() {
	agent := NewAgent()
	result := agent.Run(os.Args[1:])
	fmt.Println(result)
}
