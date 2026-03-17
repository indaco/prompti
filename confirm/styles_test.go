package confirm

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/indaco/prompti/internal/theme"
)

func TestSetDefaults_ZeroValue(t *testing.T) {
	var s Styles
	s.setDefaults()

	if s.Width != defaultTheme.Width {
		t.Errorf("Width: got %d, want %d", s.Width, defaultTheme.Width)
	}
	if s.BorderColor == nil {
		t.Error("BorderColor should not be nil after setDefaults")
	}
	if theme.IsZeroBorder(s.BorderStyle) {
		t.Error("BorderStyle should not be zero after setDefaults")
	}
	// MessageStyle default is lipgloss.NewStyle() which equals the zero value;
	// lipgloss.Style contains slices, so use reflect.DeepEqual.
	if !reflect.DeepEqual(s.MessageStyle, defaultTheme.MessageStyle) {
		t.Error("MessageStyle should match defaultTheme after setDefaults")
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
	s.setDefaults()

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
