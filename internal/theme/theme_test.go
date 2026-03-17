package theme

import (
	"testing"

	"charm.land/lipgloss/v2"
)

func TestIsZeroStyle(t *testing.T) {
	tests := []struct {
		name string
		s    lipgloss.Style
		want bool
	}{
		{"zero value", lipgloss.Style{}, true},
		{"new style with bold", lipgloss.NewStyle().Bold(true), false},
		{"new style with foreground", lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZeroStyle(tt.s); got != tt.want {
				t.Errorf("IsZeroStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsZeroBorder(t *testing.T) {
	tests := []struct {
		name string
		b    lipgloss.Border
		want bool
	}{
		{"zero value", lipgloss.Border{}, true},
		{"rounded border", lipgloss.RoundedBorder(), false},
		{"normal border", lipgloss.NormalBorder(), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZeroBorder(tt.b); got != tt.want {
				t.Errorf("IsZeroBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
