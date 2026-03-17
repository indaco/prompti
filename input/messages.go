package input

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

var (
	// messages
	promptMessage = func(question string, icon string, iconColor color.Color) string {
		return lipgloss.NewStyle().MarginTop(1).Render(lipgloss.JoinHorizontal(lipgloss.Center,
			prefixIconStyle(iconColor).Render(icon),
			questionStyle.Render(question),
			promptStyle.Render(theme.PromptMark)))
	}
)
