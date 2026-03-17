package progressbar

import (
	"testing"

	"charm.land/lipgloss/v2"
)

func TestNewProgressBarModel_DefaultBlend(t *testing.T) {
	cfg := &Config{
		Items: []string{"a", "b"},
		Styles: Styles{
			GradientFrom: nil,
			GradientTo:   nil,
		},
	}

	model := newProgressBarModel(cfg)

	// The model should be usable; verify it has the expected width.
	if model.Width() != 40 {
		t.Errorf("Width: got %d, want 40", model.Width())
	}
}

func TestNewProgressBarModel_CustomGradients(t *testing.T) {
	cfg := &Config{
		Items: []string{"a", "b"},
		Styles: Styles{
			GradientFrom: lipgloss.Color("#ff0000"),
			GradientTo:   lipgloss.Color("#00ff00"),
		},
	}

	model := newProgressBarModel(cfg)

	if model.Width() != 40 {
		t.Errorf("Width: got %d, want 40", model.Width())
	}
}
