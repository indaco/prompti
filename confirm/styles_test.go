package confirm

import (
	"image/color"
	"reflect"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

// -----------------------------------------------------------------------
// Dialog-mode style tests
// -----------------------------------------------------------------------

func TestSetDefaults_ZeroValue(t *testing.T) {
	var s Styles
	s.setDefaults(ModeDialog)

	if s.Width != defaultDialogTheme.Width {
		t.Errorf("Width: got %d, want %d", s.Width, defaultDialogTheme.Width)
	}
	if s.BorderColor == nil {
		t.Error("BorderColor should not be nil after setDefaults")
	}
	if theme.IsZeroBorder(s.BorderStyle) {
		t.Error("BorderStyle should not be zero after setDefaults")
	}
	// MessageStyle default is lipgloss.NewStyle() which equals the zero value;
	// lipgloss.Style contains slices, so use reflect.DeepEqual.
	if !reflect.DeepEqual(s.MessageStyle, defaultDialogTheme.MessageStyle) {
		t.Error("MessageStyle should match defaultDialogTheme after setDefaults")
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
	if theme.IsZeroStyle(s.DialogStyle) {
		t.Error("DialogStyle should not be zero after setDefaults")
	}
}

func TestSetDefaults_CustomWidthAndBorderColorPreserved(t *testing.T) {
	customWidth := 80
	customColor := color.RGBA{R: 0, G: 0, B: 255, A: 255}

	s := Styles{
		Width:       customWidth,
		BorderColor: customColor,
	}
	s.setDefaults(ModeDialog)

	if s.Width != customWidth {
		t.Errorf("Width: got %d, want %d", s.Width, customWidth)
	}
	if !reflect.DeepEqual(s.BorderColor, customColor) {
		t.Error("BorderColor was overwritten")
	}
	// Other fields should still get defaults
	if theme.IsZeroStyle(s.QuestionStyle) {
		t.Error("QuestionStyle should get default when not set")
	}
	if theme.IsZeroStyle(s.ButtonStyle) {
		t.Error("ButtonStyle should get default when not set")
	}
}

// -----------------------------------------------------------------------
// Inline-mode style tests
// -----------------------------------------------------------------------

func TestSetInlineDefaults_ZeroValue(t *testing.T) {
	var s Styles
	s.setDefaults(ModeInline)

	if s.PrefixIcon != defaultInlineTheme.PrefixIcon {
		t.Errorf("PrefixIcon: got %q, want %q", s.PrefixIcon, defaultInlineTheme.PrefixIcon)
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

func TestSetInlineDefaults_CustomValuesPreserved(t *testing.T) {
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
	s.setDefaults(ModeInline)

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
