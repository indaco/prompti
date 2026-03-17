package input

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

func testModel() model {
	cfg := &Config{
		Message:     "Enter your name",
		Placeholder: "type here...",
	}
	cfg.setDefaults()
	return cfg.initialModel()
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

func TestView_ContainsTextInputView(t *testing.T) {
	m := testModel()
	v := m.View()

	if v.Content == "" {
		t.Error("expected non-empty view content from textInput")
	}
}

func TestUpdate_Enter_QuitsWhenNoValidation(t *testing.T) {
	m := testModel()
	// Set a value so it does not hit the empty path.
	m.textInput.SetValue("hello")
	msg := keyPress(tea.KeyEnter, 0, "")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if um.quitting {
		t.Error("expected quitting to remain false (enter does not set quitting)")
	}
	if um.err != nil {
		t.Errorf("expected no error, got %v", um.err)
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command")
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
