package toggle

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

// Styles is the struct representing the style configuration options.
type Styles struct {
	PrefixIcon        string
	PrefixIconColor   color.Color
	QuestionStyle     lipgloss.Style
	ButtonStyle       lipgloss.Style
	ActiveButtonStyle lipgloss.Style
	DialogStyle       lipgloss.Style
	DividerStyle      lipgloss.Style
}

const (
	questionMark = theme.QuestionMark
	okLabel      = "Yes"
	cancelLabel  = "No"
	cursorLabel  = ">"
	divider      = "/"
)

var (
	// Styles
	prefixIconStyle = func(c color.Color) lipgloss.Style {
		return lipgloss.NewStyle().Bold(true).Foreground(c).PaddingRight(1)
	}
	questionStyle = lipgloss.NewStyle().Bold(true).PaddingRight(1)
	buttonStyle   = lipgloss.NewStyle().
			Foreground(theme.Muted).Margin(0)
	activeButtonStyle = buttonStyle.
				Foreground(theme.Cyan).
				Underline(true)
	dialogStyle  = lipgloss.NewStyle()
	dividerStyle = lipgloss.NewStyle().
			Margin(0, 1).
			Foreground(theme.Muted)

	defaultTheme = Styles{
		PrefixIcon:        questionMark,
		PrefixIconColor:   theme.Purple,
		QuestionStyle:     questionStyle,
		ButtonStyle:       buttonStyle,
		ActiveButtonStyle: activeButtonStyle,
		DialogStyle:       dialogStyle,
		DividerStyle:      dividerStyle,
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
	if theme.IsZeroStyle(t.QuestionStyle) {
		t.QuestionStyle = defaultTheme.QuestionStyle
	}
	if theme.IsZeroStyle(t.ButtonStyle) {
		t.ButtonStyle = defaultTheme.ButtonStyle
	}
	if theme.IsZeroStyle(t.ActiveButtonStyle) {
		t.ActiveButtonStyle = defaultTheme.ActiveButtonStyle
	}
	if theme.IsZeroStyle(t.DialogStyle) {
		t.DialogStyle = defaultTheme.DialogStyle
	}
	if theme.IsZeroStyle(t.DividerStyle) {
		t.DividerStyle = defaultTheme.DividerStyle
	}
}
