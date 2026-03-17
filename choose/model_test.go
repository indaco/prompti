package choose

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

// helper to build a model with a few items for testing.
func testModel() model {
	cfg := &Config{
		Title: "Pick one",
	}
	items := []Item{
		{Name: "alpha", Desc: "alpha"},
		{Name: "beta", Desc: "beta"},
		{Name: "gamma", Desc: "gamma"},
	}
	cfg.setDefaults(len(items))
	return cfg.initialModel(items)
}

func keyPress(code rune, mod tea.KeyMod, text string) tea.KeyPressMsg {
	return tea.KeyPressMsg{Code: code, Mod: mod, Text: text}
}

func TestUpdate_CtrlC_SetsQuitting(t *testing.T) {
	m := testModel()
	msg := keyPress('c', tea.ModCtrl, "")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if !um.quitting {
		t.Error("expected quitting to be true after ctrl+c")
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command")
	}
}

func TestUpdate_Esc_SetsQuitting(t *testing.T) {
	m := testModel()
	msg := keyPress(tea.KeyEscape, 0, "")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if !um.quitting {
		t.Error("expected quitting to be true after esc")
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command")
	}
}

func TestUpdate_Enter_SetsChoice(t *testing.T) {
	m := testModel()
	// The list starts with index 0, so the selected item is "alpha".
	msg := keyPress(tea.KeyEnter, 0, "")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if um.choice != "alpha" {
		t.Errorf("expected choice = %q, got %q", "alpha", um.choice)
	}
	if um.quitting {
		t.Error("expected quitting to be false after enter (choice made)")
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command")
	}
}

func TestView_NoChoice_ReturnsListView(t *testing.T) {
	m := testModel()
	v := m.View()

	if v.Content == "" {
		t.Error("expected non-empty view content when no choice is made")
	}
	// The list view should contain item text.
	if strings.Contains(v.Content, "alpha") == false &&
		strings.Contains(v.Content, "beta") == false {
		t.Error("expected view to contain list item text")
	}
}

func TestView_WithChoice_ReturnsChoiceString(t *testing.T) {
	m := testModel()
	m.choice = "beta"
	v := m.View()

	if v.Content != "beta" {
		t.Errorf("expected view content = %q, got %q", "beta", v.Content)
	}
}

func TestUpdate_WindowSizeMsg(t *testing.T) {
	m := testModel()
	msg := tea.WindowSizeMsg{Width: 120, Height: 40}
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if um.quitting {
		t.Error("window size msg should not set quitting")
	}
	if cmd != nil {
		t.Error("expected nil command from window size msg")
	}
}
