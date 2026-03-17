package choose

import (
	"image/color"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

func TestSetDefaults_ZeroValue(t *testing.T) {
	var s Styles
	s.setDefaults()

	if s.PrefixIcon != defaultTheme.PrefixIcon {
		t.Errorf("PrefixIcon: got %q, want %q", s.PrefixIcon, defaultTheme.PrefixIcon)
	}
	if s.PrefixIconColor == nil {
		t.Error("PrefixIconColor should not be nil after setDefaults")
	}
	if theme.IsZeroStyle(s.TitleStyle) {
		t.Error("TitleStyle should not be zero after setDefaults")
	}
	if theme.IsZeroStyle(s.TitleBarStyle) {
		t.Error("TitleBarStyle should not be zero after setDefaults")
	}
	if s.ItemIcon == "" {
		t.Error("ItemIcon should not be empty after setDefaults")
	}
	if theme.IsZeroStyle(s.ItemStyle) {
		t.Error("ItemStyle should not be zero after setDefaults")
	}
	if theme.IsZeroStyle(s.SelectedItemStyle) {
		t.Error("SelectedItemStyle should not be zero after setDefaults")
	}
}

func TestSetDefaults_CustomValuesPreserved(t *testing.T) {
	customIcon := ">>>"
	customColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	customStyle := lipgloss.NewStyle().Italic(true)
	customItemIcon := "-"

	s := Styles{
		PrefixIcon:        customIcon,
		PrefixIconColor:   customColor,
		TitleStyle:        customStyle,
		TitleBarStyle:     customStyle,
		ItemIcon:          customItemIcon,
		ItemStyle:         customStyle,
		SelectedItemStyle: customStyle,
	}
	s.setDefaults()

	if s.PrefixIcon != customIcon {
		t.Errorf("PrefixIcon: got %q, want %q", s.PrefixIcon, customIcon)
	}
	if s.PrefixIconColor != customColor {
		t.Errorf("PrefixIconColor was overwritten")
	}
	if s.ItemIcon != customItemIcon {
		t.Errorf("ItemIcon: got %q, want %q", s.ItemIcon, customItemIcon)
	}
	// Style fields: verify they were not reset to the default by checking
	// they still differ from the default theme values. We use the Italic
	// property which the default theme does not set.
	if s.TitleStyle.GetItalic() != true {
		t.Error("TitleStyle custom value was overwritten")
	}
	if s.ItemStyle.GetItalic() != true {
		t.Error("ItemStyle custom value was overwritten")
	}
	if s.SelectedItemStyle.GetItalic() != true {
		t.Error("SelectedItemStyle custom value was overwritten")
	}
}
