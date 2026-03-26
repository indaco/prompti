package detail

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

func testModel() model {
	cfg := &Config{
		Summary: "More info",
		Content: "Here are the details.",
	}
	cfg.setDefaults()
	return cfg.initialModel()
}

func keyPress(code rune, mod tea.KeyMod, text string) tea.KeyPressMsg {
	return tea.KeyPressMsg{Code: code, Mod: mod, Text: text}
}

func TestInitialExpanded_IsFalse(t *testing.T) {
	m := testModel()
	if m.expanded {
		t.Error("expected initial expanded to be false")
	}
}

func TestUpdate_Enter_TogglesExpanded(t *testing.T) {
	m := testModel() // expanded starts false
	msg := keyPress(tea.KeyEnter, 0, "")
	updated, _ := m.Update(msg)
	um := updated.(model)

	if !um.expanded {
		t.Error("expected expanded to be true after enter")
	}
	if um.quitting {
		t.Error("expected quitting to be false after enter")
	}

	// Toggle again.
	updated2, _ := um.Update(msg)
	um2 := updated2.(model)
	if um2.expanded {
		t.Error("expected expanded to be false after second enter toggle")
	}
}

func TestUpdate_Space_TogglesExpanded(t *testing.T) {
	m := testModel()
	msg := keyPress(' ', 0, " ")
	updated, _ := m.Update(msg)
	um := updated.(model)

	if !um.expanded {
		t.Error("expected expanded to be true after space")
	}
}

func TestUpdate_Right_SetsExpanded(t *testing.T) {
	m := testModel() // expanded starts false
	msg := keyPress(tea.KeyRight, 0, "")
	updated, _ := m.Update(msg)
	um := updated.(model)

	if !um.expanded {
		t.Error("expected expanded to be true after right")
	}
}

func TestUpdate_Left_CollapsesExpanded(t *testing.T) {
	m := testModel()
	m.expanded = true
	msg := keyPress(tea.KeyLeft, 0, "")
	updated, _ := m.Update(msg)
	um := updated.(model)

	if um.expanded {
		t.Error("expected expanded to be false after left")
	}
}

func TestUpdate_J_SetsExpanded(t *testing.T) {
	m := testModel()
	msg := keyPress('j', 0, "j")
	updated, _ := m.Update(msg)
	um := updated.(model)

	if !um.expanded {
		t.Error("expected expanded to be true after 'j'")
	}
}

func TestUpdate_K_CollapsesExpanded(t *testing.T) {
	m := testModel()
	m.expanded = true
	msg := keyPress('k', 0, "k")
	updated, _ := m.Update(msg)
	um := updated.(model)

	if um.expanded {
		t.Error("expected expanded to be false after 'k'")
	}
}

func TestUpdate_Y_SetsExpandedTrue_Quitting(t *testing.T) {
	m := testModel()
	msg := keyPress('y', 0, "y")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if !um.expanded {
		t.Error("expected expanded to be true after 'y'")
	}
	if !um.quitting {
		t.Error("expected quitting to be true after 'y'")
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command")
	}
}

func TestUpdate_N_SetsExpandedFalse_Quitting(t *testing.T) {
	m := testModel()
	m.expanded = true
	msg := keyPress('n', 0, "n")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if um.expanded {
		t.Error("expected expanded to be false after 'n'")
	}
	if !um.quitting {
		t.Error("expected quitting to be true after 'n'")
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

func TestUpdate_Q_SetsQuitting(t *testing.T) {
	m := testModel()
	msg := keyPress('q', 0, "q")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if !um.quitting {
		t.Error("expected quitting to be true after 'q'")
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

func TestView_Collapsed_DoesNotContainContent(t *testing.T) {
	m := testModel()
	m.expanded = false
	v := m.View()

	// The content text should not appear when collapsed.
	if containsSubstring(v.Content, "Here are the details.") {
		t.Error("expected collapsed view to not contain detail content")
	}
}

func TestView_Expanded_ContainsContent(t *testing.T) {
	m := testModel()
	m.expanded = true
	v := m.View()

	if !containsSubstring(v.Content, "Here are the details.") {
		t.Error("expected expanded view to contain detail content")
	}
}

// containsSubstring is a helper to check if a string contains a substring.
func containsSubstring(s, sub string) bool {
	return len(s) >= len(sub) && searchSubstring(s, sub)
}

func searchSubstring(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
