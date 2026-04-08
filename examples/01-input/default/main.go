// Basic text input prompt with default settings.
//
// Run from the repository root:
//
//	go run ./examples/01-input/default
package main

import (
	"fmt"

	"github.com/indaco/prompti/input"
)

func main() {
	questionPrompt := &input.Config{
		Message:     "What's the name of your project?",
		Placeholder: "Please, provide a name for your project",
		ErrorMsg:    "Project name is mandatory",
	}

	result, _ := input.Run(questionPrompt)
	fmt.Println(result)
}
