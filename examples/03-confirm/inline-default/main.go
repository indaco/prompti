// Inline confirm prompt with default settings.
//
// Run from the repository root:
//
//	go run ./examples/03-confirm/inline-default
package main

import (
	"fmt"

	"github.com/indaco/prompti/confirm"
)

func main() {
	result, _ := confirm.Run(&confirm.Config{
		Mode:     confirm.ModeInline,
		Question: "Continue?",
	})
	fmt.Println(result)
}
