package input

import (
	"bytes"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
)

func testField() (*Field, *string) {
	var val string
	cfg := &Config{
		Message:     "Enter name",
		Placeholder: "type here...",
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

	msg := tea.KeyPressMsg{Code: 'a', Text: "a"}
	updated, cmd := f.Update(msg)
	if cmd != nil {
		t.Error("expected nil cmd when blurred")
	}
	_ = updated
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
	if val != "" {
		t.Errorf("expected empty initial value, got %q", val)
	}
}

func TestField_GetKey(t *testing.T) {
	f, _ := testField()
	f.Key = "name"
	if f.GetKey() != "name" {
		t.Errorf("expected key = %q, got %q", "name", f.GetKey())
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

func TestField_RunAccessible(t *testing.T) {
	var val string
	cfg := &Config{
		Message:     "Enter name",
		Placeholder: "your name",
	}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("Alice\n")
	err := f.RunAccessible(&out, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "Alice" {
		t.Errorf("expected val = %q, got %q", "Alice", val)
	}
}

func TestField_RunAccessible_UsesDefault(t *testing.T) {
	var val string
	cfg := &Config{
		Message: "Enter name",
		Initial: "Bob",
	}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("\n")
	err := f.RunAccessible(&out, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "Bob" {
		t.Errorf("expected val = %q, got %q", "Bob", val)
	}
}

func TestField_RunAccessible_Validation(t *testing.T) {
	var val string
	cfg := &Config{
		Message:      "Enter email",
		ValidateFunc: ValidateEmail,
	}
	f := NewField(cfg, &val)

	var out bytes.Buffer
	in := strings.NewReader("not-an-email\n")
	err := f.RunAccessible(&out, in)
	if err == nil {
		t.Error("expected validation error")
	}
}
