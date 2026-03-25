package detail_test

import (
	"fmt"

	"github.com/indaco/prompti/detail"
)

func ExampleDefaultStyles() {
	s := detail.DefaultStyles()
	fmt.Println("PrefixIcon:", s.PrefixIcon)
	// Output:
	// PrefixIcon: ?
}
