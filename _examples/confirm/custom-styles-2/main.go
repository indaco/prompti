package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/indaco/prompti/confirm"
)

var (
	cyan   = compat.AdaptiveColor{Light: lipgloss.Color("#4f46e5"), Dark: lipgloss.Color("#c7d2fe")}
	green  = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}
	red    = compat.AdaptiveColor{Light: lipgloss.Color("#ef4444"), Dark: lipgloss.Color("#ef4444")}
	purple = compat.AdaptiveColor{Light: lipgloss.Color("#7e22ce"), Dark: lipgloss.Color("#a855f7")}

	myCustomStyle = confirm.Styles{
		ActiveButtonStyle: lipgloss.NewStyle().Padding(0, 3).
			Margin(1, 1).Background(green).Foreground(purple),
		DialogStyle: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderTopForeground(purple).
			BorderBottomForeground(green).
			BorderLeftForeground(cyan).
			BorderRightForeground(cyan).
			Margin(1, 0, 1, 0).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true).
			Width(50).
			Align(lipgloss.Center),
	}

	confirmConfig = &confirm.Config{
		Question: "Continue?",
		Styles:   myCustomStyle,
	}
)

func main() {
	result, _ := confirm.Run(confirmConfig)
	fmt.Println(result)
}
