// Confirmation dialog with a message body and custom border colors.
//
// Run from the repository root:
//
//	go run ./examples/03-confirm/custom-styles
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/indaco/prompti/confirm"
)

var (
	cyan  = compat.AdaptiveColor{Light: lipgloss.Color("#4f46e5"), Dark: lipgloss.Color("#c7d2fe")}
	green = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}

	infoText = `Lorem ipsum dolor sit amet,
consectetur adipiscing elit %s...`

	Green   = lipgloss.NewStyle().Foreground(green).Render
	message = fmt.Sprintf(infoText, Green("elit"))

	myCustomStyle = confirm.Styles{
		Width:       60,
		BorderColor: cyan,
	}

	confirmConfig = &confirm.Config{
		Message:  message,
		Question: "Continue?",
		Styles:   myCustomStyle,
	}
)

func main() {
	result, _ := confirm.Run(confirmConfig)
	fmt.Println(result)
}
