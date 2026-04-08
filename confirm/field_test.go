package confirm

import (
	"bytes"
	"strings"
	"testing"

	"unicode/utf8"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
)

// -----------------------------------------------------------------------
// Dialog-mode field tests
// -----------------------------------------------------------------------

func testField() (*Field, *bool) {
	var val bool
	cfg := &Config{
		Message:  "Are you sure?",
		Question: "Proceed?",
	}
	f := NewField(cfg, &val)
	f.Focus()
	return f, &val
}

func TestField_ImplementsHuhField(t *testing.T) {
	f, _ := testField()
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
	// Initial confirmation is true; toggle should make it false.
	toggleKeys := f.Keymap.Confirm.Toggle.Keys()
	if len(toggleKeys) == 0 {
		t.Fatal("no toggle keys defined")
	}
	r, _ := utf8.DecodeRuneInString(toggleKeys[0])
	msg := tea.KeyPressMsg{Code: r, Text: toggleKeys[0]}
	updated, _ := f.Update(msg)
	uf := updated.(*Field)
	if uf.inner.confirmation {
		t.Error("expected confirmation to be false after toggle")
	}
}

func TestField_Update_Accept(t *testing.T) {
	f, val := testField()
	f.inner.confirmation = false
	msg := tea.KeyPressMsg{Code: 'y', Text: "y"}
	_, cmd := f.Update(msg)
	if cmd == nil {
		t.Error("expected non-nil cmd from accept")
	}
	if !*val {
		t.Error("expected value to be true after accept")
	}
}

func TestField_Update_Reject(t *testing.T) {
	f, val := testField()
	f.inner.confirmation = true
	msg := tea.KeyPressMsg{Code: 'n', Text: "n"}

	// Need to check if reject key matches 'n'. Default reject keys are "n", "N".
	if !key.Matches(msg, f.Keymap.Confirm.Reject) {
		t.Skip("reject key does not match 'n'")
	}
	_, cmd := f.Update(msg)
	if cmd == nil {
		t.Error("expected non-nil cmd from reject")
	}
	if *val {
		t.Error("expected value to be false after reject")
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
	if val != true {
		t.Errorf("expected initial value = true, got %v", val)
	}
}

func TestField_GetKey(t *testing.T) {
	f, _ := testField()
	f.Key = "confirm"
	if f.GetKey() != "confirm" {
		t.Errorf("expected key = %q, got %q", "confirm", f.GetKey())
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

func TestField_WithTheme(t *testing.T) {
	f, _ := testField()
	result := f.WithTheme(huh.ThemeFunc(huh.ThemeCharm))
	if result != f {
		t.Error("expected same field returned from WithTheme")
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
	result := f.WithWidth(50)
	if result != f {
		t.Error("expected same field returned from WithWidth")
	}
	if f.Width != 50 {
		t.Errorf("expected width = 50, got %d", f.Width)
	}
}

func TestField_WithHeight(t *testing.T) {
	f, _ := testField()
	result := f.WithHeight(5)
	if result != f {
		t.Error("expected same field returned from WithHeight")
	}
	if f.Height != 5 {
		t.Errorf("expected height = 5, got %d", f.Height)
	}
}

func TestField_WithPosition(t *testing.T) {
	f, _ := testField()
	pos := huh.FieldPosition{
		Field: 0, FirstField: 0, LastField: 1,
		Group: 0, FirstGroup: 0, LastGroup: 0, GroupCount: 1,
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

func TestField_RunAccessible_Yes(t *testing.T) {
	var val bool
	cfg := &Config{Question: "Continue?"}
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
}

func TestField_RunAccessible_No(t *testing.T) {
	var val bool
	cfg := &Config{Question: "Continue?"}
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
	cfg := &Config{Question: "Continue?"}
	f := NewField(cfg, &val)
	// Default confirmation is true.

	var out bytes.Buffer
	in := strings.NewReader("\n")
	err := f.RunAccessible(&out, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !val {
		t.Error("expected val = true (default)")
	}
}

func TestField_RunAccessible_Invalid(t *testing.T) {
	var val bool
	cfg := &Config{Question: "Continue?"}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("maybe\n")
	err := f.RunAccessible(&out, in)
	if err == nil {
		t.Error("expected error for invalid input")
	}
}

// -----------------------------------------------------------------------
// Inline-mode field tests
// -----------------------------------------------------------------------

func testInlineField() (*Field, *bool) {
	var val bool
	cfg := &Config{
		Mode:     ModeInline,
		Question: "Continue?",
	}
	f := NewField(cfg, &val)
	f.Focus()
	return f, &val
}

func TestInlineField_ImplementsHuhField(t *testing.T) {
	f, _ := testInlineField()
	var _ huh.Field = f
	_ = f
}

func TestInlineField_Init(t *testing.T) {
	f, _ := testInlineField()
	cmd := f.Init()
	if cmd != nil {
		t.Error("expected nil cmd from Init")
	}
}

func TestInlineField_FocusBlur(t *testing.T) {
	f, _ := testInlineField()
	f.Blur()
	if f.Focused {
		t.Error("expected Focused to be false after Blur")
	}
	f.Focus()
	if !f.Focused {
		t.Error("expected Focused to be true after Focus")
	}
}

func TestInlineField_Update_IgnoredWhenBlurred(t *testing.T) {
	f, _ := testInlineField()
	f.Blur()
	msg := tea.KeyPressMsg{Code: tea.KeyEnter}
	_, cmd := f.Update(msg)
	if cmd != nil {
		t.Error("expected nil cmd when blurred")
	}
}

func TestInlineField_Update_Toggle(t *testing.T) {
	f, _ := testInlineField()
	toggleKeys := f.Keymap.Confirm.Toggle.Keys()
	if len(toggleKeys) == 0 {
		t.Fatal("no toggle keys defined")
	}
	r, _ := utf8.DecodeRuneInString(toggleKeys[0])
	msg := tea.KeyPressMsg{Code: r, Text: toggleKeys[0]}
	updated, _ := f.Update(msg)
	uf := updated.(*Field)
	if uf.inner.confirmation {
		t.Error("expected confirmation to be false after toggle")
	}
}

func TestInlineField_Update_Accept(t *testing.T) {
	f, val := testInlineField()
	f.inner.confirmation = false
	msg := tea.KeyPressMsg{Code: 'y', Text: "y"}
	_, cmd := f.Update(msg)
	if cmd == nil {
		t.Error("expected non-nil cmd from accept")
	}
	if !*val {
		t.Error("expected value to be true after accept")
	}
}

func TestInlineField_Update_Reject(t *testing.T) {
	f, val := testInlineField()
	f.inner.confirmation = true
	msg := tea.KeyPressMsg{Code: 'n', Text: "n"}

	if !key.Matches(msg, f.Keymap.Confirm.Reject) {
		t.Skip("reject key does not match 'n'")
	}
	_, cmd := f.Update(msg)
	if cmd == nil {
		t.Error("expected non-nil cmd from reject")
	}
	if *val {
		t.Error("expected value to be false after reject")
	}
}

func TestInlineField_View_ReturnsString(t *testing.T) {
	f, _ := testInlineField()
	v := f.View()
	if v == "" {
		t.Error("expected non-empty view")
	}
}

func TestInlineField_GetValue(t *testing.T) {
	f, _ := testInlineField()
	val := f.GetValue()
	if val != true {
		t.Errorf("expected initial value = true, got %v", val)
	}
}

func TestInlineField_GetKey(t *testing.T) {
	f, _ := testInlineField()
	f.Key = "toggle"
	if f.GetKey() != "toggle" {
		t.Errorf("expected key = %q, got %q", "toggle", f.GetKey())
	}
}

func TestInlineField_SkipZoom(t *testing.T) {
	f, _ := testInlineField()
	if f.Skip() {
		t.Error("expected Skip() = false")
	}
	if f.Zoom() {
		t.Error("expected Zoom() = false")
	}
}

func TestInlineField_WithTheme(t *testing.T) {
	f, _ := testInlineField()
	result := f.WithTheme(huh.ThemeFunc(huh.ThemeCharm))
	if result != f {
		t.Error("expected same field returned from WithTheme")
	}
}

func TestInlineField_WithKeyMap(t *testing.T) {
	f, _ := testInlineField()
	km := huh.NewDefaultKeyMap()
	result := f.WithKeyMap(km)
	if result != f {
		t.Error("expected same field returned from WithKeyMap")
	}
}

func TestInlineField_WithWidth(t *testing.T) {
	f, _ := testInlineField()
	result := f.WithWidth(60)
	if result != f {
		t.Error("expected same field returned from WithWidth")
	}
	if f.Width != 60 {
		t.Errorf("expected width = 60, got %d", f.Width)
	}
}

func TestInlineField_WithHeight(t *testing.T) {
	f, _ := testInlineField()
	result := f.WithHeight(3)
	if result != f {
		t.Error("expected same field returned from WithHeight")
	}
	if f.Height != 3 {
		t.Errorf("expected height = 3, got %d", f.Height)
	}
}

func TestInlineField_WithPosition(t *testing.T) {
	f, _ := testInlineField()
	pos := huh.FieldPosition{
		Field: 0, FirstField: 0, LastField: 1,
		Group: 0, FirstGroup: 0, LastGroup: 0, GroupCount: 1,
	}
	result := f.WithPosition(pos)
	if result != f {
		t.Error("expected same field returned from WithPosition")
	}
}

func TestInlineField_KeyBinds(t *testing.T) {
	f, _ := testInlineField()
	binds := f.KeyBinds()
	if len(binds) == 0 {
		t.Error("expected non-empty key bindings")
	}
}

func TestInlineField_RunAccessible_Yes(t *testing.T) {
	var val bool
	cfg := &Config{Mode: ModeInline, Question: "Enable?"}
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
}

func TestInlineField_RunAccessible_No(t *testing.T) {
	var val bool
	cfg := &Config{Mode: ModeInline, Question: "Enable?"}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("no\n")
	err := f.RunAccessible(&out, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val {
		t.Error("expected val = false")
	}
}

func TestInlineField_RunAccessible_Default(t *testing.T) {
	var val bool
	cfg := &Config{Mode: ModeInline, Question: "Enable?"}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("\n")
	err := f.RunAccessible(&out, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !val {
		t.Error("expected val = true (default)")
	}
}

func TestInlineField_RunAccessible_Invalid(t *testing.T) {
	var val bool
	cfg := &Config{Mode: ModeInline, Question: "Enable?"}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("maybe\n")
	err := f.RunAccessible(&out, in)
	if err == nil {
		t.Error("expected error for invalid input")
	}
}
