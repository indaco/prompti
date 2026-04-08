// Animated progress bar with default settings.
//
// Run from the repository root:
//
//	go run ./examples/05-progressbar/default
package main

import (
	"fmt"
	"os"

	"github.com/indaco/prompti/progressbar"
)

func main() {
	fruits := []string{
		"apple",
		"banana",
		"orange",
		"grapes",
		"mellon",
		"strawberry",
		"mango",
		"lemon",
		"apricot",
		"peach",
		"papaya",
		"kiwi",
		"pear",
		"guava",
		"almond",
		"coconut",
		"blackberry",
		"cherry",
		"grapes",
	}

	pbConfig := &progressbar.Config{Items: fruits}

	if err := progressbar.Run(pbConfig); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
