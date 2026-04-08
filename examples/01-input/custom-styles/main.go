// Text input prompt with custom styles.
//
// Run from the repository root:
//
//	go run ./examples/01-input/custom-styles
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/indaco/prompti/input"
)

func main() {
	questionPrompt := &input.Config{
		Message:     "What's the name of your project?",
		Placeholder: "Please, provide a name for your project",
		ErrorMsg:    "Project name is mandatory",
		Styles: input.Styles{
			PrefixIcon:      "*",
			PrefixIconColor: compat.AdaptiveColor{Light: lipgloss.Color("#ef4444"), Dark: lipgloss.Color("#ef4444")},
			PlaceholderStyle: lipgloss.NewStyle().
				Background(compat.AdaptiveColor{Light: lipgloss.Color("#3b82f6"), Dark: lipgloss.Color("#3b82f6")}).
				Foreground(compat.AdaptiveColor{Light: lipgloss.Color("#fde68a"), Dark: lipgloss.Color("#fffbeb")}),
		},
	}

	result, _ := input.Run(questionPrompt)
	fmt.Println(result)
}
