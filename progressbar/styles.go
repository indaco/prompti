package progressbar

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// Styles is the struct representing the style configuration options.
type Styles struct {
	ShowLabel        bool
	CurrentItemStyle lipgloss.Style
	GradientFrom     color.Color
	GradientTo       color.Color
}
