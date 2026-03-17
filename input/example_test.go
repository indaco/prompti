package input_test

import (
	"fmt"

	"github.com/indaco/prompti/input"
)

func ExampleValidateAlphanumeric() {
	err := input.ValidateAlphanumeric("hello123")
	fmt.Println("hello123:", err)

	err = input.ValidateAlphanumeric("not-valid!")
	fmt.Println("not-valid!:", err)
	// Output:
	// hello123: <nil>
	// not-valid!: it must be an alphanumeric value
}

func ExampleValidateEmail() {
	err := input.ValidateEmail("user@example.com")
	fmt.Println("user@example.com:", err)

	err = input.ValidateEmail("not-an-email")
	fmt.Println("not-an-email:", err)
	// Output:
	// user@example.com: <nil>
	// not-an-email: it must be a valid email address
}

func ExampleValidateURL() {
	err := input.ValidateURL("https://example.com")
	fmt.Println("https://example.com:", err)

	err = input.ValidateURL("not-a-url")
	fmt.Println("not-a-url:", err)
	// Output:
	// https://example.com: <nil>
	// not-a-url: it must be a valid URL
}
