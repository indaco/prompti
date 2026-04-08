// Using multiple prompti components together in a multi-group huh form.
//
// Run from the repository root:
//
//	go run ./examples/06-huh-form
package main

import (
	"fmt"
	"log"

	"charm.land/huh/v2"
	"github.com/indaco/prompti/choose"
	"github.com/indaco/prompti/confirm"
	"github.com/indaco/prompti/input"
)

func main() {
	var (
		projectName string
		language    string
		useGit      bool
		confirmed   bool
	)

	languages := []choose.Item{
		{Name: "Go", Desc: "Fast and simple"},
		{Name: "Rust", Desc: "Safe and fast"},
		{Name: "TypeScript", Desc: "JavaScript, but typed"},
		{Name: "Python", Desc: "Batteries included"},
	}

	form := huh.NewForm(
		huh.NewGroup(
			input.NewField(&input.Config{
				Message:     "Project name",
				Placeholder: "my-awesome-project",
			}, &projectName),

			choose.NewField(&choose.Config{
				Title: "Pick a language",
			}, languages, &language),
		),

		huh.NewGroup(
			confirm.NewField(&confirm.Config{
				Mode:     confirm.ModeInline,
				Question: "Initialize a git repository?",
			}, &useGit),

			confirm.NewField(&confirm.Config{
				Question: "Create the project?",
			}, &confirmed),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nProject: %s\n", projectName)
	fmt.Printf("Language: %s\n", language)
	fmt.Printf("Git: %v\n", useGit)
	fmt.Printf("Confirmed: %v\n", confirmed)
}
