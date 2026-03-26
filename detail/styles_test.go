package detail

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
	if theme.IsZeroStyle(s.SummaryStyle) {
		t.Error("SummaryStyle should not be zero after setDefaults")
	}
	if theme.IsZeroStyle(s.IndicatorStyle) {
		t.Error("IndicatorStyle should not be zero after setDefaults")
	}
	if theme.IsZeroStyle(s.ContentStyle) {
		t.Error("ContentStyle should not be zero after setDefaults")
	}
	if theme.IsZeroStyle(s.HintStyle) {
		t.Error("HintStyle should not be zero after setDefaults")
	}
}

func TestSetDefaults_CustomValuesPreserved(t *testing.T) {
	customIcon := ">>>"
	customColor := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	customStyle := lipgloss.NewStyle().Italic(true)

	s := Styles{
		PrefixIcon:      customIcon,
		PrefixIconColor: customColor,
		SummaryStyle:    customStyle,
		IndicatorStyle:  customStyle,
		ContentStyle:    customStyle,
		DialogStyle:     customStyle,
		HintStyle:       customStyle,
	}
	s.setDefaults()

	if s.PrefixIcon != customIcon {
		t.Errorf("PrefixIcon: got %q, want %q", s.PrefixIcon, customIcon)
	}
	if s.PrefixIconColor != customColor {
		t.Error("PrefixIconColor was overwritten")
	}
	if s.SummaryStyle.GetItalic() != true {
		t.Error("SummaryStyle custom value was overwritten")
	}
	if s.IndicatorStyle.GetItalic() != true {
		t.Error("IndicatorStyle custom value was overwritten")
	}
	if s.ContentStyle.GetItalic() != true {
		t.Error("ContentStyle custom value was overwritten")
	}
	if s.DialogStyle.GetItalic() != true {
		t.Error("DialogStyle custom value was overwritten")
	}
	if s.HintStyle.GetItalic() != true {
		t.Error("HintStyle custom value was overwritten")
	}
}
