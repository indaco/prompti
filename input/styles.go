package input

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

// Styles is the struct representing the style configuration options.
type Styles struct {
	PrefixIcon       string
	PrefixIconColor  color.Color
	PromptStyle      lipgloss.Style
	TextStyle        lipgloss.Style
	BackgroundStyle  lipgloss.Style
	PlaceholderStyle lipgloss.Style
	CursorStyle      lipgloss.Style
}

// default styles
var defaultTheme = Styles{
	PrefixIcon:       theme.QuestionMark,
	PrefixIconColor:  theme.Purple,
	TextStyle:        noStyle,
	PlaceholderStyle: placeholderStyle,
	CursorStyle:      cursorStyle,
}

var (
	// styles
	noStyle          = lipgloss.NewStyle()
	placeholderStyle = noStyle.Faint(true)
	cursorStyle      = lipgloss.NewStyle()
	prefixIconStyle  = func(c color.Color) lipgloss.Style {
		return lipgloss.NewStyle().MarginRight(1).Bold(true).Foreground(c)
	}
	questionStyle   = lipgloss.NewStyle().MarginRight(1).Bold(true)
	promptStyle     = lipgloss.NewStyle().Faint(true)
	cancelMarkStyle = lipgloss.NewStyle().MarginRight(1)
	errorStyle      = lipgloss.NewStyle().MarginRight(1).Bold(true).Foreground(theme.Red)
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
	if theme.IsZeroStyle(t.PromptStyle) {
		t.PromptStyle = defaultTheme.PromptStyle
	}
	if theme.IsZeroStyle(t.TextStyle) {
		t.TextStyle = defaultTheme.TextStyle
	}
	if theme.IsZeroStyle(t.PlaceholderStyle) {
		t.PlaceholderStyle = defaultTheme.PlaceholderStyle
	}
	if theme.IsZeroStyle(t.CursorStyle) {
		t.CursorStyle = defaultTheme.CursorStyle
	}
}
