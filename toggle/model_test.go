package toggle

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

func testModel() model {
	cfg := &Config{
		Question: "Continue?",
	}
	cfg.setDefaults()
	return cfg.initialModel()
}

func keyPress(code rune, mod tea.KeyMod, text string) tea.KeyPressMsg {
	return tea.KeyPressMsg{Code: code, Mod: mod, Text: text}
}

func TestInitialConfirmation_IsTrue(t *testing.T) {
	m := testModel()
	if !m.confirmation {
		t.Error("expected initial confirmation to be true")
	}
}

func TestUpdate_Y_SetsConfirmationTrue_Quitting(t *testing.T) {
	m := testModel()
	m.confirmation = false // start with false to verify "y" sets it true
	msg := keyPress('y', 0, "y")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if !um.confirmation {
		t.Error("expected confirmation to be true after 'y'")
	}
	if !um.quitting {
		t.Error("expected quitting to be true after 'y'")
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command")
	}
}

func TestUpdate_N_SetsConfirmationFalse_Quitting(t *testing.T) {
	m := testModel()
	msg := keyPress('n', 0, "n")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if um.confirmation {
		t.Error("expected confirmation to be false after 'n'")
	}
	if !um.quitting {
		t.Error("expected quitting to be true after 'n'")
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command")
	}
}

func TestUpdate_Left_TogglesConfirmation(t *testing.T) {
	m := testModel() // confirmation starts true
	msg := keyPress(tea.KeyLeft, 0, "")
	updated, _ := m.Update(msg)
	um := updated.(model)

	if um.confirmation {
		t.Error("expected confirmation to be false after 'left' (was true)")
	}

	// Toggle again.
	updated2, _ := um.Update(msg)
	um2 := updated2.(model)
	if !um2.confirmation {
		t.Error("expected confirmation to be true after second 'left' toggle")
	}
}

func TestUpdate_Right_TogglesConfirmation(t *testing.T) {
	m := testModel() // confirmation starts true
	msg := keyPress(tea.KeyRight, 0, "")
	updated, _ := m.Update(msg)
	um := updated.(model)

	if um.confirmation {
		t.Error("expected confirmation to be false after 'right' (was true)")
	}
}

func TestUpdate_Enter_QuitsWithCurrentState(t *testing.T) {
	m := testModel()
	m.confirmation = false // set to false before enter
	msg := keyPress(tea.KeyEnter, 0, "")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if um.confirmation {
		t.Error("expected confirmation to remain false after enter")
	}
	if !um.quitting {
		t.Error("expected quitting to be true after enter")
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command")
	}
}

func TestUpdate_CtrlC_SetsQuitting(t *testing.T) {
	m := testModel()
	msg := keyPress('c', tea.ModCtrl, "")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if um.confirmation {
		t.Error("expected confirmation to be false after ctrl+c")
	}
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

	if um.confirmation {
		t.Error("expected confirmation to be false after esc")
	}
	if !um.quitting {
		t.Error("expected quitting to be true after esc")
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command")
	}
}

func TestView_Quitting_ReturnsEmpty(t *testing.T) {
	m := testModel()
	m.quitting = true
	v := m.View()

	if v.Content != "" {
		t.Errorf("expected empty view content when quitting, got %q", v.Content)
	}
}

func TestView_NotQuitting_ReturnsNonEmpty(t *testing.T) {
	m := testModel()
	v := m.View()

	if v.Content == "" {
		t.Error("expected non-empty view content when not quitting")
	}
}
