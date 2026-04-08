package detail

import (
	"bytes"
	"strings"
	"testing"
	"unicode/utf8"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
)

func testField() (*Field, *bool) {
	var val bool
	cfg := &Config{
		Summary: "Terms of Service",
		Content: "You agree to these terms.",
	}
	f := NewField(cfg, &val)
	f.Focus()
	return f, &val
}

func TestField_ImplementsHuhField(t *testing.T) {
	f, _ := testField()
	// Compile-time check is implicit; runtime verify the interface.
	var _ huh.Field = f
	_ = f
}

func TestField_Init(t *testing.T) {
	f, _ := testField()
	cmd := f.Init()
	if cmd != nil {
		t.Error("expected nil cmd from Init")
	}
}

func TestField_FocusBlur(t *testing.T) {
	f, _ := testField()
	f.Blur()
	if f.Focused {
		t.Error("expected Focused to be false after Blur")
	}

	f.Focus()
	if !f.Focused {
		t.Error("expected Focused to be true after Focus")
	}
}

func TestField_Update_IgnoredWhenBlurred(t *testing.T) {
	f, _ := testField()
	f.Blur()

	msg := tea.KeyPressMsg{Code: tea.KeyEnter}
	_, cmd := f.Update(msg)
	if cmd != nil {
		t.Error("expected nil cmd when blurred")
	}
}

func TestField_Update_Toggle(t *testing.T) {
	f, _ := testField()
	// Initially expanded is false; toggle should make it true.
	toggleKeys := f.Keymap.Confirm.Toggle.Keys()
	if len(toggleKeys) == 0 {
		t.Fatal("no toggle keys defined")
	}
	r, _ := utf8.DecodeRuneInString(toggleKeys[0])
	msg := tea.KeyPressMsg{Code: r, Text: toggleKeys[0]}
	updated, _ := f.Update(msg)
	uf := updated.(*Field)
	if !uf.inner.expanded {
		t.Error("expected expanded to be true after toggle")
	}

	// Toggle again should collapse.
	updated2, _ := uf.Update(msg)
	uf2 := updated2.(*Field)
	if uf2.inner.expanded {
		t.Error("expected expanded to be false after second toggle")
	}
}

func TestField_Update_Next(t *testing.T) {
	f, val := testField()
	f.inner.expanded = true
	msg := tea.KeyPressMsg{Code: tea.KeyTab}
	_, cmd := f.Update(msg)
	if cmd == nil {
		t.Error("expected non-nil cmd from next")
	}
	if !*val {
		t.Error("expected value to be true after next with expanded detail")
	}
}

func TestField_Update_Prev(t *testing.T) {
	f, _ := testField()
	msg := tea.KeyPressMsg{Code: tea.KeyTab, Mod: tea.ModShift}
	_, cmd := f.Update(msg)
	if cmd == nil {
		t.Error("expected non-nil cmd from prev")
	}
}

func TestField_View_ReturnsString(t *testing.T) {
	f, _ := testField()
	v := f.View()
	if v == "" {
		t.Error("expected non-empty view")
	}
}

func TestField_GetValue(t *testing.T) {
	f, _ := testField()
	val := f.GetValue()
	if val != false {
		t.Errorf("expected initial value = false, got %v", val)
	}
}

func TestField_GetValue_Expanded(t *testing.T) {
	f, _ := testField()
	f.inner.expanded = true
	val := f.GetValue()
	if val != true {
		t.Errorf("expected value = true after expanding, got %v", val)
	}
}

func TestField_GetKey(t *testing.T) {
	f, _ := testField()
	f.Key = "detail"
	if f.GetKey() != "detail" {
		t.Errorf("expected key = %q, got %q", "detail", f.GetKey())
	}
}

func TestField_SkipZoom(t *testing.T) {
	f, _ := testField()
	if f.Skip() {
		t.Error("expected Skip() = false")
	}
	if f.Zoom() {
		t.Error("expected Zoom() = false")
	}
}

func TestField_Error(t *testing.T) {
	f, _ := testField()
	if f.Error() != nil {
		t.Error("expected nil error initially")
	}
}

func TestField_WithTheme(t *testing.T) {
	f, _ := testField()
	result := f.WithTheme(huh.ThemeFunc(huh.ThemeCharm))
	if result != f {
		t.Error("expected same field returned from WithTheme")
	}
	// Second call should be a no-op.
	result2 := f.WithTheme(huh.ThemeFunc(huh.ThemeBase))
	if result2 != f {
		t.Error("expected same field returned from second WithTheme")
	}
}

func TestField_WithKeyMap(t *testing.T) {
	f, _ := testField()
	km := huh.NewDefaultKeyMap()
	result := f.WithKeyMap(km)
	if result != f {
		t.Error("expected same field returned from WithKeyMap")
	}
}

func TestField_WithWidth(t *testing.T) {
	f, _ := testField()
	result := f.WithWidth(40)
	if result != f {
		t.Error("expected same field returned from WithWidth")
	}
	if f.Width != 40 {
		t.Errorf("expected width = 40, got %d", f.Width)
	}
}

func TestField_WithHeight(t *testing.T) {
	f, _ := testField()
	result := f.WithHeight(10)
	if result != f {
		t.Error("expected same field returned from WithHeight")
	}
	if f.Height != 10 {
		t.Errorf("expected height = 10, got %d", f.Height)
	}
}

func TestField_WithPosition(t *testing.T) {
	f, _ := testField()
	pos := huh.FieldPosition{
		Field:      0,
		FirstField: 0,
		LastField:  2,
		Group:      0,
		FirstGroup: 0,
		LastGroup:  0,
		GroupCount: 1,
	}
	result := f.WithPosition(pos)
	if result != f {
		t.Error("expected same field returned from WithPosition")
	}
}

func TestField_KeyBinds(t *testing.T) {
	f, _ := testField()
	binds := f.KeyBinds()
	if len(binds) == 0 {
		t.Error("expected non-empty key bindings")
	}
}

func TestField_Blur_SyncsValue(t *testing.T) {
	f, val := testField()
	f.inner.expanded = true
	f.Blur()
	if !*val {
		t.Error("expected value to be true after blur with expanded detail")
	}
}

func TestField_RunAccessible_Yes(t *testing.T) {
	var val bool
	cfg := &Config{
		Summary: "Terms of Service",
		Content: "You agree to these terms.",
	}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("y\n")
	err := f.RunAccessible(&out, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !val {
		t.Error("expected val = true")
	}
	if !strings.Contains(out.String(), "You agree to these terms.") {
		t.Error("expected content to be displayed when expanded")
	}
}

func TestField_RunAccessible_No(t *testing.T) {
	var val bool
	cfg := &Config{
		Summary: "Terms of Service",
		Content: "You agree to these terms.",
	}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("n\n")
	err := f.RunAccessible(&out, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val {
		t.Error("expected val = false")
	}
}

func TestField_RunAccessible_Default(t *testing.T) {
	var val bool
	cfg := &Config{
		Summary: "Terms of Service",
		Content: "You agree to these terms.",
	}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("\n")
	err := f.RunAccessible(&out, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val {
		t.Error("expected val = false (default)")
	}
}

func TestField_RunAccessible_Invalid(t *testing.T) {
	var val bool
	cfg := &Config{
		Summary: "Terms of Service",
		Content: "You agree to these terms.",
	}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("maybe\n")
	err := f.RunAccessible(&out, in)
	if err == nil {
		t.Error("expected error for invalid input")
	}
}
