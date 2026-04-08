// Animated progress bar with custom gradient colors and item label.
//
// Run from the repository root:
//
//	go run ./examples/05-progressbar/custom-styles
package main

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/progressbar"
)

func main() {
	fruits := []string{
		"apple",
		"banana",
		"orange",
		"grapes",
		"mellon",
		"strawberry",
		"mango",
		"lemon",
		"apricot",
		"peach",
		"papaya",
		"kiwi",
		"pear",
		"guava",
		"almond",
		"coconut",
		"blackberry",
		"cherry",
		"grapes",
	}

	pbConfig := &progressbar.Config{
		Items:         fruits,
		OnProgressMsg: "Eating:",
		Styles: progressbar.Styles{
			CurrentItemStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("411")),
			ShowLabel:        true,
			GradientFrom:     lipgloss.Color("#FF7CCB"),
			GradientTo:       lipgloss.Color("#FDFF8C"),
		}}

	if err := progressbar.Run(pbConfig); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
