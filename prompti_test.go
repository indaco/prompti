package prompti_test

import (
	"errors"
	"fmt"

	"github.com/indaco/prompti"
)

func ExampleErrCancelled() {
	err := prompti.ErrCancelled
	fmt.Println(errors.Is(err, prompti.ErrCancelled))
	fmt.Println(err)
	// Output:
	// true
	// prompti: operation cancelled
}

func ExampleErrEmpty() {
	err := prompti.ErrEmpty
	fmt.Println(errors.Is(err, prompti.ErrEmpty))
	fmt.Println(err)
	// Output:
	// true
	// prompti: empty input
}
