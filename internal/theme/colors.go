package theme

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
)

var (
	Cyan    = compat.AdaptiveColor{Light: lipgloss.Color("#4f46e5"), Dark: lipgloss.Color("#c7d2fe")}
	Green   = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}
	Red     = compat.AdaptiveColor{Light: lipgloss.Color("#ef4444"), Dark: lipgloss.Color("#ef4444")}
	Purple  = compat.AdaptiveColor{Light: lipgloss.Color("#7e22ce"), Dark: lipgloss.Color("#a855f7")}
	Neutral = compat.AdaptiveColor{Light: lipgloss.Color("#737373"), Dark: lipgloss.Color("#a3a3a3")}
	Amber   = compat.AdaptiveColor{Light: lipgloss.Color("#fef3c7"), Dark: lipgloss.Color("#fef3c7")}
	Muted   = compat.AdaptiveColor{Light: lipgloss.Color("#D9DCCF"), Dark: lipgloss.Color("#383838")}
)
