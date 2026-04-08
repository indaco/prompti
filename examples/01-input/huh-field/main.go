// Using the prompti input component as a huh.Field in a form.
//
// Run from the repository root:
//
//	go run ./examples/01-input/huh-field
package main

import (
	"fmt"
	"log"

	"charm.land/huh/v2"
	"github.com/indaco/prompti/input"
)

func main() {
	var projectName string

	form := huh.NewForm(
		huh.NewGroup(
			input.NewField(&input.Config{
				Message:     "What's the name of your project?",
				Placeholder: "Please, provide a name for your project",
			}, &projectName),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Project name:", projectName)
}
