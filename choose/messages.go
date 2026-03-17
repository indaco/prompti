package choose

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

var (
	// messages
	titleMessage = func(icon string, iconColor color.Color, titleStyle lipgloss.Style, title string) string {
		return lipgloss.NewStyle().Render(
			lipgloss.JoinHorizontal(lipgloss.Center,
				prefixIconStyle(iconColor).Render(icon),
				titleStyle.Render(theme.Whitespace),
				titleStyle.Render(title),
				titleStyle.Render(theme.Whitespace),
				titleStyle.Render(theme.PromptMark)))
	}
)
