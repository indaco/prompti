package toggle

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
	if theme.IsZeroStyle(s.QuestionStyle) {
		t.Error("QuestionStyle should not be zero after setDefaults")
	}
	if theme.IsZeroStyle(s.ButtonStyle) {
		t.Error("ButtonStyle should not be zero after setDefaults")
	}
	if theme.IsZeroStyle(s.ActiveButtonStyle) {
		t.Error("ActiveButtonStyle should not be zero after setDefaults")
	}
	if theme.IsZeroStyle(s.DividerStyle) {
		t.Error("DividerStyle should not be zero after setDefaults")
	}
}

func TestSetDefaults_CustomValuesPreserved(t *testing.T) {
	customIcon := ">>>"
	customColor := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	customStyle := lipgloss.NewStyle().Italic(true)

	s := Styles{
		PrefixIcon:        customIcon,
		PrefixIconColor:   customColor,
		QuestionStyle:     customStyle,
		ButtonStyle:       customStyle,
		ActiveButtonStyle: customStyle,
		DialogStyle:       customStyle,
		DividerStyle:      customStyle,
	}
	s.setDefaults()

	if s.PrefixIcon != customIcon {
		t.Errorf("PrefixIcon: got %q, want %q", s.PrefixIcon, customIcon)
	}
	if s.PrefixIconColor != customColor {
		t.Error("PrefixIconColor was overwritten")
	}
	if s.QuestionStyle.GetItalic() != true {
		t.Error("QuestionStyle custom value was overwritten")
	}
	if s.ButtonStyle.GetItalic() != true {
		t.Error("ButtonStyle custom value was overwritten")
	}
	if s.ActiveButtonStyle.GetItalic() != true {
		t.Error("ActiveButtonStyle custom value was overwritten")
	}
	if s.DialogStyle.GetItalic() != true {
		t.Error("DialogStyle custom value was overwritten")
	}
	if s.DividerStyle.GetItalic() != true {
		t.Error("DividerStyle custom value was overwritten")
	}
}
