// Confirmation dialog with default settings.
//
// Run from the repository root:
//
//	go run ./examples/03-confirm/default
package main

import (
	"fmt"

	"github.com/indaco/prompti/confirm"
)

func main() {
	result, _ := confirm.Run(&confirm.Config{Question: "Continue?"})
	fmt.Println(result)
}
