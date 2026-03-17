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

	pbConfig := &progressbar.Config{Items: fruits, RunConcurrently: true}

	fmt.Println("Run commands concurrently")
	if err := progressbar.Run(pbConfig); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
