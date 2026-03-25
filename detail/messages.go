package detail

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

var (
	// messages
	summaryMessage = func(icon string, iconColor color.Color, sStyle lipgloss.Style, iStyle lipgloss.Style, indicator string, summary string) string {
		return lipgloss.NewStyle().Render(
			lipgloss.JoinHorizontal(lipgloss.Center,
				prefixIconStyle(iconColor).Render(icon),
				sStyle.Render(theme.Whitespace),
				iStyle.Render(indicator),
				sStyle.Render(summary)))
	}
)
