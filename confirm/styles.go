package confirm

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

// Styles is the struct representing the style configuration options.
type Styles struct {
	// --- Dialog-mode fields ---

	// Width sets the dialog box width (ModeDialog only).
	Width int
	// BorderColor sets the dialog border color (ModeDialog only).
	BorderColor color.Color
	// BorderStyle sets the dialog border style (ModeDialog only).
	BorderStyle lipgloss.Border
	// MessageStyle styles the optional message body (ModeDialog only).
	MessageStyle lipgloss.Style

	// --- Inline-mode fields ---

	// PrefixIcon is the icon shown before the question (ModeInline only).
	PrefixIcon string
	// PrefixIconColor is the color of the prefix icon (ModeInline only).
	PrefixIconColor color.Color
	// DividerStyle styles the separator between toggle options (ModeInline only).
	DividerStyle lipgloss.Style

	// --- Shared fields ---

	// QuestionStyle styles the question text.
	QuestionStyle lipgloss.Style
	// ButtonStyle styles the inactive option button.
	ButtonStyle lipgloss.Style
	// ActiveButtonStyle styles the currently selected option button.
	ActiveButtonStyle lipgloss.Style
	// DialogStyle is the outer container style.
	DialogStyle lipgloss.Style
}

// -----------------------------------------------------------------------
// Dialog-mode defaults
// -----------------------------------------------------------------------

var (
	dialogWidth         = 50
	dialogBorderStyle   = lipgloss.RoundedBorder()
	dialogMessageStyle  = lipgloss.NewStyle()
	dialogQuestionStyle = lipgloss.NewStyle().Bold(true)
	dialogButtonStyle   = lipgloss.NewStyle().
				Foreground(theme.Amber).
				Background(theme.Neutral).
				Padding(0, 3).
				Margin(1, 1)
	dialogActiveButtonStyle = dialogButtonStyle.
				Foreground(theme.Amber).
				Background(theme.Purple).
				Underline(true)
	dialogContainerStyle = lipgloss.NewStyle().
				Border(dialogBorderStyle).
				BorderForeground(theme.Purple).
				Margin(1, 0, 0, 0).
				Padding(1, 0).
				BorderTop(true).
				BorderLeft(true).
				BorderRight(true).
				BorderBottom(true).
				Width(dialogWidth).Align(lipgloss.Center)

	defaultDialogTheme = Styles{
		Width:             dialogWidth,
		BorderColor:       color.Color(theme.Purple),
		BorderStyle:       dialogBorderStyle,
		MessageStyle:      dialogMessageStyle,
		QuestionStyle:     dialogQuestionStyle,
		ButtonStyle:       dialogButtonStyle,
		ActiveButtonStyle: dialogActiveButtonStyle,
		DialogStyle:       dialogContainerStyle,
	}
)

// -----------------------------------------------------------------------
// Inline-mode defaults
// -----------------------------------------------------------------------

const (
	questionMark = theme.QuestionMark
)

var (
	prefixIconStyle = func(c color.Color) lipgloss.Style {
		return lipgloss.NewStyle().Bold(true).Foreground(c).PaddingRight(1)
	}
	inlineQuestionStyle = lipgloss.NewStyle().Bold(true).PaddingRight(1)
	inlineButtonStyle   = lipgloss.NewStyle().
				Foreground(theme.Muted).Margin(0)
	inlineActiveButtonStyle = inlineButtonStyle.
				Foreground(theme.Cyan).
				Underline(true)
	inlineDialogStyle  = lipgloss.NewStyle()
	inlineDividerStyle = lipgloss.NewStyle().
				Margin(0, 1).
				Foreground(theme.Muted)

	defaultInlineTheme = Styles{
		PrefixIcon:        questionMark,
		PrefixIconColor:   theme.Purple,
		QuestionStyle:     inlineQuestionStyle,
		ButtonStyle:       inlineButtonStyle,
		ActiveButtonStyle: inlineActiveButtonStyle,
		DialogStyle:       inlineDialogStyle,
		DividerStyle:      inlineDividerStyle,
	}
)

// -----------------------------------------------------------------------
// Public helpers
// -----------------------------------------------------------------------

func setCustomDialogStyles(t *Styles) lipgloss.Style {
	_width := dialogWidth
	if t.Width != 0 {
		_width = t.Width
	}

	_borderStyle := dialogBorderStyle
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

// DefaultStyles returns the default styles for ModeDialog.
func DefaultStyles() (s Styles) {
	return defaultDialogTheme
}

// DefaultInlineStyles returns the default styles for ModeInline.
func DefaultInlineStyles() (s Styles) {
	return defaultInlineTheme
}

func (t *Styles) setDefaults(mode Mode) {
	if mode == ModeInline {
		t.setInlineDefaults()
	} else {
		t.setDialogDefaults()
	}
}

func (t *Styles) setDialogDefaults() {
	if t.Width == 0 {
		t.Width = defaultDialogTheme.Width
	}

	if t.BorderColor == nil {
		t.BorderColor = defaultDialogTheme.BorderColor
	}

	if theme.IsZeroBorder(t.BorderStyle) {
		t.BorderStyle = defaultDialogTheme.BorderStyle
	}

	if theme.IsZeroStyle(t.MessageStyle) {
		t.MessageStyle = defaultDialogTheme.MessageStyle
	}

	if theme.IsZeroStyle(t.QuestionStyle) {
		t.QuestionStyle = defaultDialogTheme.QuestionStyle
	}

	if theme.IsZeroStyle(t.ButtonStyle) {
		t.ButtonStyle = defaultDialogTheme.ButtonStyle
	}

	if theme.IsZeroStyle(t.ActiveButtonStyle) {
		t.ActiveButtonStyle = defaultDialogTheme.ActiveButtonStyle
	}

	if theme.IsZeroStyle(t.DialogStyle) {
		t.DialogStyle = setCustomDialogStyles(t)
	}
}

func (t *Styles) setInlineDefaults() {
	if t.PrefixIcon == "" {
		t.PrefixIcon = defaultInlineTheme.PrefixIcon
	}
	if t.PrefixIconColor == nil {
		t.PrefixIconColor = defaultInlineTheme.PrefixIconColor
	}
	if theme.IsZeroStyle(t.QuestionStyle) {
		t.QuestionStyle = defaultInlineTheme.QuestionStyle
	}
	if theme.IsZeroStyle(t.ButtonStyle) {
		t.ButtonStyle = defaultInlineTheme.ButtonStyle
	}
	if theme.IsZeroStyle(t.ActiveButtonStyle) {
		t.ActiveButtonStyle = defaultInlineTheme.ActiveButtonStyle
	}
	if theme.IsZeroStyle(t.DialogStyle) {
		t.DialogStyle = defaultInlineTheme.DialogStyle
	}
	if theme.IsZeroStyle(t.DividerStyle) {
		t.DividerStyle = defaultInlineTheme.DividerStyle
	}
}
