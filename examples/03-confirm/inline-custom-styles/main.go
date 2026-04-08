// Inline confirm prompt with custom labels, divider, and styles.
//
// Run from the repository root:
//
//	go run ./examples/03-confirm/inline-custom-styles
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
	red   = compat.AdaptiveColor{Light: lipgloss.Color("#ef4444"), Dark: lipgloss.Color("#ef4444")}

	myCustomStyle = confirm.Styles{
		PrefixIcon:        "★",
		PrefixIconColor:   red,
		DialogStyle:       lipgloss.NewStyle().Margin(1, 0),
		ButtonStyle:       lipgloss.NewStyle().Bold(true).Foreground(cyan),
		ActiveButtonStyle: lipgloss.NewStyle().Foreground(green),
	}

	inlineConfig = &confirm.Config{
		Mode:              confirm.ModeInline,
		Question:          "How do you feel?",
		OkButtonLabel:     "I'm super ok",
		CancelButtonLabel: "Next question, please!",
		Divider:           "|",
		Styles:            myCustomStyle,
	}
)

func main() {
	result, _ := confirm.Run(inlineConfig)
	fmt.Println(result)
}
