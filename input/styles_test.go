package input

import (
	"image/color"
	"reflect"
	"testing"

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
	// TextStyle and CursorStyle defaults are lipgloss.NewStyle() which is the
	// zero value, so we verify with reflect.DeepEqual (lipgloss.Style contains
	// slices and is not comparable with ==).
	if !reflect.DeepEqual(s.TextStyle, defaultTheme.TextStyle) {
		t.Error("TextStyle should match defaultTheme after setDefaults")
	}
	if theme.IsZeroStyle(s.PlaceholderStyle) {
		t.Error("PlaceholderStyle should not be zero after setDefaults")
	}
	if !reflect.DeepEqual(s.CursorStyle, defaultTheme.CursorStyle) {
		t.Error("CursorStyle should match defaultTheme after setDefaults")
	}
}

func TestSetDefaults_CustomPrefixIconPreserved(t *testing.T) {
	customIcon := ">>>"
	customColor := color.RGBA{R: 128, G: 128, B: 0, A: 255}

	s := Styles{
		PrefixIcon:      customIcon,
		PrefixIconColor: customColor,
	}
	s.setDefaults()

	if s.PrefixIcon != customIcon {
		t.Errorf("PrefixIcon: got %q, want %q", s.PrefixIcon, customIcon)
	}
	if !reflect.DeepEqual(s.PrefixIconColor, customColor) {
		t.Error("PrefixIconColor was overwritten")
	}
	// Other fields should still get defaults
	if theme.IsZeroStyle(s.PlaceholderStyle) {
		t.Error("PlaceholderStyle should get default when not set")
	}
}
