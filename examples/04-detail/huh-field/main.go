// Using the prompti detail component as a huh.Field in a form.
//
// Run from the repository root:
//
//	go run ./examples/04-detail/huh-field
package main

import (
	"fmt"
	"log"

	"charm.land/huh/v2"
	"github.com/indaco/prompti/detail"
)

func main() {
	var expanded bool

	form := huh.NewForm(
		huh.NewGroup(
			detail.NewField(&detail.Config{
				Summary: "Terms of Service",
				Content: "By using this software, you agree to the terms and conditions.\nPlease review the full agreement before proceeding.",
			}, &expanded),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Detail expanded:", expanded)
}
