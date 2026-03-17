package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/indaco/prompti/choose"
)

var (
	amber  = compat.AdaptiveColor{Light: lipgloss.Color("#f59e0b"), Dark: lipgloss.Color("#fbbf24")}
	purple = compat.AdaptiveColor{Light: lipgloss.Color("#7e22ce"), Dark: lipgloss.Color("#a855f7")}
	green  = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}

	myCustomStyle = choose.Styles{
		PrefixIcon:        "★",
		TitleStyle:        lipgloss.NewStyle().Background(green).Foreground(purple).Padding(0, 1),
		TitleBarStyle:     lipgloss.NewStyle(),
		ItemIcon:          "#",
		ItemStyle:         lipgloss.NewStyle().Foreground(amber),
		SelectedItemStyle: lipgloss.NewStyle().Foreground(purple),
	}
)

func main() {
	foodSelectionPrompt := &choose.Config{
		Title:    "What do you wanna eat tonight?",
		ErrorMsg: "Please, select your meal.",
		ShowHelp: true,
		Styles:   myCustomStyle,
	}

	entries := []choose.Item{
		{Name: "pizza", Desc: "It's always pizza time!"},
		{Name: "kebab", Desc: "I feel turkish today, kebab!"},
		{Name: "carbonara", Desc: "Carbonara, NO cream, please!"},
	}

	result, _ := choose.Run(foodSelectionPrompt, entries)
	fmt.Println(result)
}
