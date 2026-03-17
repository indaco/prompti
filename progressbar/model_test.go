package progressbar

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

func testModel() model {
	cfg := &Config{
		Items:          []string{"item1", "item2", "item3"},
		OnCompletedMsg: "All done!",
		OnProgressMsg:  "Processing",
	}
	cfg.setDefaults()
	return cfg.initialModel()
}

func keyPress(code rune, mod tea.KeyMod, text string) tea.KeyPressMsg {
	return tea.KeyPressMsg{Code: code, Mod: mod, Text: text}
}

func TestUpdate_KeyPress_Quits(t *testing.T) {
	m := testModel()
	msg := keyPress('q', 0, "q")
	_, cmd := m.Update(msg)

	if cmd == nil {
		t.Error("expected a non-nil quit command on any key press")
	}
}

func TestUpdate_IncrementMsg_AdvancesIndex(t *testing.T) {
	m := testModel()
	msg := IncrementMsg("item1")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if um.index != 1 {
		t.Errorf("expected index = 1, got %d", um.index)
	}
	if um.done {
		t.Error("expected done to be false after first increment")
	}
	if cmd == nil {
		t.Error("expected a non-nil command after increment")
	}
}

func TestUpdate_IncrementMsg_LastItem_SetsDone(t *testing.T) {
	m := testModel()
	m.index = len(m.items) - 1 // set to last item
	msg := IncrementMsg("item3")
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if !um.done {
		t.Error("expected done to be true when index >= len(items)-1")
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command when done")
	}
}

func TestUpdate_IncrementErrMsg_SetsError(t *testing.T) {
	m := testModel()
	errMsg := IncrementErrMsg{Err: errTest}
	updated, cmd := m.Update(errMsg)
	um := updated.(model)

	if um.err == nil {
		t.Error("expected error to be set")
	}
	if cmd == nil {
		t.Error("expected a non-nil quit command on error")
	}
}

var errTest = &testError{}

type testError struct{}

func (e *testError) Error() string { return "test error" }

func TestUpdate_WindowSizeMsg(t *testing.T) {
	m := testModel()
	msg := tea.WindowSizeMsg{Width: 120, Height: 40}
	updated, cmd := m.Update(msg)
	um := updated.(model)

	if um.done {
		t.Error("window size msg should not set done")
	}
	if cmd != nil {
		t.Error("expected nil command from window size msg")
	}
}

func TestView_WhenDone_ShowsCompletedMsg(t *testing.T) {
	m := testModel()
	m.done = true
	v := m.View()

	if !strings.Contains(v.Content, "All done!") {
		t.Errorf("expected view to contain completed message, got %q", v.Content)
	}
}

func TestView_WhenError_ShowsError(t *testing.T) {
	m := testModel()
	m.err = errTest
	v := m.View()

	if !strings.Contains(v.Content, "test error") {
		t.Errorf("expected view to contain error message, got %q", v.Content)
	}
}

func TestView_InProgress_ShowsProgressBar(t *testing.T) {
	m := testModel()
	v := m.View()

	if v.Content == "" {
		t.Error("expected non-empty view content during progress")
	}
}
