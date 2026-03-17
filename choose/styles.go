package choose

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

// Styles is the struct representing the style configuration options.
type Styles struct {
	PrefixIcon        string
	PrefixIconColor   color.Color
	TitleStyle        lipgloss.Style
	TitleBarStyle     lipgloss.Style
	ItemIcon          string
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
}

var (
	// styles
	noStyle         = lipgloss.NewStyle()
	prefixIconStyle = func(c color.Color) lipgloss.Style {
		return lipgloss.NewStyle().Bold(true).Foreground(c)
	}
	// default theme styles
	defaultTitleStyle        = lipgloss.NewStyle().Bold(true)
	defaultTitleBarStyle     = noStyle.PaddingBottom(1)
	defaultItemIcon          = lipgloss.NewStyle().Render("▸")
	defaultItemStyle         = lipgloss.NewStyle().PaddingLeft(0)
	defaultSelectedItemStyle = defaultItemStyle.Foreground(theme.Cyan)

	defaultTheme = Styles{
		PrefixIcon:        theme.QuestionMark,
		PrefixIconColor:   theme.Purple,
		TitleStyle:        defaultTitleStyle,
		TitleBarStyle:     defaultTitleBarStyle,
		ItemIcon:          defaultItemIcon,
		ItemStyle:         defaultItemStyle,
		SelectedItemStyle: defaultSelectedItemStyle,
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
	if theme.IsZeroStyle(t.TitleStyle) {
		t.TitleStyle = defaultTheme.TitleStyle
	}
	if theme.IsZeroStyle(t.TitleBarStyle) {
		t.TitleBarStyle = defaultTheme.TitleBarStyle
	}
	if t.ItemIcon == "" {
		t.ItemIcon = defaultTheme.ItemIcon
	}
	if theme.IsZeroStyle(t.ItemStyle) {
		t.ItemStyle = defaultTheme.ItemStyle
	}
	if theme.IsZeroStyle(t.SelectedItemStyle) {
		t.SelectedItemStyle = defaultTheme.SelectedItemStyle
	}
}
