// Collapsible detail/summary prompt with default settings.
//
// Run from the repository root:
//
//	go run ./examples/04-detail/default
package main

import (
	"fmt"

	"github.com/indaco/prompti/detail"
)

func main() {
	expanded, err := detail.Run(&detail.Config{
		Summary: "What is prompti?",
		Content: "prompti is a collection of interactive terminal UI prompts\nbuilt on the charmbracelet bubbletea framework.",
	})
	fmt.Println("expanded:", expanded, "err:", err)
}
