package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/indaco/prompti/detail"
)

var (
	cyan   = compat.AdaptiveColor{Light: lipgloss.Color("#4f46e5"), Dark: lipgloss.Color("#c7d2fe")}
	green  = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}
	red    = compat.AdaptiveColor{Light: lipgloss.Color("#ef4444"), Dark: lipgloss.Color("#ef4444")}
	purple = compat.AdaptiveColor{Light: lipgloss.Color("#7e22ce"), Dark: lipgloss.Color("#a855f7")}

	myCustomStyle = detail.Styles{
		PrefixIcon:      "★",
		PrefixIconColor: red,
		SummaryStyle:    lipgloss.NewStyle().Bold(true).Foreground(purple).PaddingRight(1),
		IndicatorStyle:  lipgloss.NewStyle().Foreground(green).PaddingRight(1),
		ContentStyle:    lipgloss.NewStyle().Foreground(cyan).PaddingLeft(3).MarginTop(1),
		DialogStyle:     lipgloss.NewStyle().Margin(1, 0),
	}

	detailConfig = &detail.Config{
		Summary: "Release Notes v0.3.0",
		Content: "- Added detail/summary component\n- Improved theme system\n- Bug fixes and performance improvements",
		Styles:  myCustomStyle,
	}
)

func main() {
	expanded, _ := detail.Run(detailConfig)
	fmt.Println("expanded:", expanded)
}
