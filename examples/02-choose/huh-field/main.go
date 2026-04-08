// Using the prompti choose component as a huh.Field in a form.
//
// Run from the repository root:
//
//	go run ./examples/02-choose/huh-field
package main

import (
	"fmt"
	"log"

	"charm.land/huh/v2"
	"github.com/indaco/prompti/choose"
)

func main() {
	var meal string

	entries := []choose.Item{
		{Name: "pizza", Desc: "It's always pizza time!"},
		{Name: "kebab", Desc: "I feel turkish today, kebab!"},
		{Name: "carbonara", Desc: "Carbonara, NO cream, please!"},
	}

	form := huh.NewForm(
		huh.NewGroup(
			choose.NewField(&choose.Config{
				Title: "What do you wanna eat tonight?",
			}, entries, &meal),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Selected:", meal)
}
