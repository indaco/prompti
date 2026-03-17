package confirm

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

// Styles is the struct representing the style configuration options.
type Styles struct {
	Width             int
	BorderColor       color.Color
	BorderStyle       lipgloss.Border
	MessageStyle      lipgloss.Style
	QuestionStyle     lipgloss.Style
	ButtonStyle       lipgloss.Style
	ActiveButtonStyle lipgloss.Style
	DialogStyle       lipgloss.Style
}

var (
	// Styles
	width         = 50
	borderStyle   = lipgloss.RoundedBorder()
	messageStyle  = lipgloss.NewStyle()
	questionStyle = lipgloss.NewStyle().Bold(true)
	buttonStyle   = lipgloss.NewStyle().
			Foreground(theme.Amber).
			Background(theme.Neutral).
			Padding(0, 3).
			Margin(1, 1)
	activeButtonStyle = buttonStyle.
				Foreground(theme.Amber).
				Background(theme.Purple).
				Underline(true)
	dialogStyle = lipgloss.NewStyle().
			Border(borderStyle).
			BorderForeground(theme.Purple).
			Margin(1, 0, 0, 0).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true).
			Width(width).Align(lipgloss.Center)

	defaultTheme = Styles{
		Width:             width,
		BorderColor:       color.Color(theme.Purple),
		BorderStyle:       borderStyle,
		MessageStyle:      messageStyle,
		QuestionStyle:     questionStyle,
		ButtonStyle:       buttonStyle,
		ActiveButtonStyle: activeButtonStyle,
		DialogStyle:       dialogStyle,
	}
)

func setCustomDialogStyles(t *Styles) lipgloss.Style {
	_width := width
	if t.Width != 0 {
		_width = t.Width
	}

	_borderStyle := borderStyle
	if !theme.IsZeroBorder(t.BorderStyle) {
		_borderStyle = t.BorderStyle
	}

	_borderColor := color.Color(theme.Purple)
	if t.BorderColor != nil {
		_borderColor = t.BorderColor
	}

	return lipgloss.NewStyle().
		Margin(1, 0, 0, 0).
		Padding(1, 0).
		Border(_borderStyle).
		BorderForeground(_borderColor).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		Width(_width).Align(lipgloss.Center)
}

// DefaultStyles sets the default styles theme.
func DefaultStyles() (s Styles) {
	return defaultTheme
}

func (t *Styles) setDefaults() {
	if t.Width == 0 {
		t.Width = defaultTheme.Width
	}

	if t.BorderColor == nil {
		t.BorderColor = defaultTheme.BorderColor
	}

	if theme.IsZeroBorder(t.BorderStyle) {
		t.BorderStyle = defaultTheme.BorderStyle
	}

	if theme.IsZeroStyle(t.MessageStyle) {
		t.MessageStyle = defaultTheme.MessageStyle
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
		t.DialogStyle = setCustomDialogStyles(t)
	}
}
