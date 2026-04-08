// Using the prompti confirm component as a huh.Field in a form.
//
// Run from the repository root:
//
//	go run ./examples/03-confirm/huh-field
package main

import (
	"fmt"
	"log"

	"charm.land/huh/v2"
	"github.com/indaco/prompti/confirm"
)

func main() {
	var confirmed bool

	form := huh.NewForm(
		huh.NewGroup(
			confirm.NewField(&confirm.Config{
				Question: "Do you want to continue?",
			}, &confirmed),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Confirmed:", confirmed)
}
