// Using the prompti confirm component in inline mode as a huh.Field in a form.
//
// Run from the repository root:
//
//	go run ./examples/03-confirm/inline-huh-field
package main

import (
	"fmt"
	"log"

	"charm.land/huh/v2"
	"github.com/indaco/prompti/confirm"
)

func main() {
	var accepted bool

	form := huh.NewForm(
		huh.NewGroup(
			confirm.NewField(&confirm.Config{
				Mode:     confirm.ModeInline,
				Question: "Accept the terms?",
			}, &accepted),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Accepted:", accepted)
}
