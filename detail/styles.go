package detail

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

// Styles is the struct representing the style configuration options.
type Styles struct {
	PrefixIcon      string
	PrefixIconColor color.Color
	SummaryStyle    lipgloss.Style
	IndicatorStyle  lipgloss.Style
	ContentStyle    lipgloss.Style
	DialogStyle     lipgloss.Style
	HintStyle       lipgloss.Style
}

const (
	questionMark              = theme.QuestionMark
	collapsedIndicatorDefault = "\u25b6"
	expandedIndicatorDefault  = "\u25bc"
)

var (
	// Styles
	prefixIconStyle = func(c color.Color) lipgloss.Style {
		return lipgloss.NewStyle().Bold(true).Foreground(c).PaddingRight(1)
	}
	summaryStyle   = lipgloss.NewStyle().Bold(true).PaddingRight(1)
	indicatorStyle = lipgloss.NewStyle().
			Foreground(theme.Cyan).PaddingRight(1)
	contentStyle = lipgloss.NewStyle().
			Foreground(theme.Neutral).
			PaddingLeft(3).
			MarginTop(1)
	dialogStyle = lipgloss.NewStyle()
	hintStyle   = lipgloss.NewStyle().
			Foreground(theme.Muted).
			MarginTop(1)

	defaultTheme = Styles{
		PrefixIcon:      questionMark,
		PrefixIconColor: theme.Purple,
		SummaryStyle:    summaryStyle,
		IndicatorStyle:  indicatorStyle,
		ContentStyle:    contentStyle,
		DialogStyle:     dialogStyle,
		HintStyle:       hintStyle,
	}
)

// DefaultStyles sets the default styles theme.
func DefaultStyles() (s Styles) {
	return defaultTheme
}

func (t *Styles) setDefaults() {
	if t.PrefixIcon == "" {
		t.PrefixIcon = defaultTheme.PrefixIcon
	}
	if t.PrefixIconColor == nil {
		t.PrefixIconColor = defaultTheme.PrefixIconColor
	}
	if theme.IsZeroStyle(t.SummaryStyle) {
		t.SummaryStyle = defaultTheme.SummaryStyle
	}
	if theme.IsZeroStyle(t.IndicatorStyle) {
		t.IndicatorStyle = defaultTheme.IndicatorStyle
	}
	if theme.IsZeroStyle(t.ContentStyle) {
		t.ContentStyle = defaultTheme.ContentStyle
	}
	if theme.IsZeroStyle(t.DialogStyle) {
		t.DialogStyle = defaultTheme.DialogStyle
	}
	if theme.IsZeroStyle(t.HintStyle) {
		t.HintStyle = defaultTheme.HintStyle
	}
}
