package main

import (
	"fmt"

	"github.com/owenrumney/pubsub/internal/app/commands"
)

func main() {
	rootCmd := commands.RootCommand()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %v\n", r)
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
