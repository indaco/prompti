package choose

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
		Title: "Pick one",
	}
	items := []Item{
		{Name: "alpha", Desc: "alpha"},
		{Name: "beta", Desc: "beta"},
		{Name: "gamma", Desc: "gamma"},
	}
	f := NewField(cfg, items, &val)
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

func TestField_View_ReturnsString(t *testing.T) {
	f, _ := testField()
	v := f.View()
	if v == "" {
		t.Error("expected non-empty view")
	}
}

func TestField_GetValue_Initial(t *testing.T) {
	f, _ := testField()
	val := f.GetValue()
	if val != "" {
		t.Errorf("expected empty initial value, got %q", val)
	}
}

func TestField_GetKey(t *testing.T) {
	f, _ := testField()
	f.Key = "pick"
	if f.GetKey() != "pick" {
		t.Errorf("expected key = %q, got %q", "pick", f.GetKey())
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
	result := f.WithWidth(60)
	if result != f {
		t.Error("expected same field returned from WithWidth")
	}
	if f.Width != 60 {
		t.Errorf("expected width = 60, got %d", f.Width)
	}
}

func TestField_WithHeight(t *testing.T) {
	f, _ := testField()
	result := f.WithHeight(15)
	if result != f {
		t.Error("expected same field returned from WithHeight")
	}
	if f.Height != 15 {
		t.Errorf("expected height = 15, got %d", f.Height)
	}
}

func TestField_WithPosition(t *testing.T) {
	f, _ := testField()
	pos := huh.FieldPosition{
		Field:      1,
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

func TestField_RunAccessible_ByNumber(t *testing.T) {
	var val string
	cfg := &Config{Title: "Pick"}
	items := []Item{
		{Name: "alpha", Desc: "alpha"},
		{Name: "beta", Desc: "beta"},
	}
	f := NewField(cfg, items, &val)

	var out bytes.Buffer
	in := strings.NewReader("2\n")
	err := f.RunAccessible(&out, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "beta" {
		t.Errorf("expected val = %q, got %q", "beta", val)
	}
}

func TestField_RunAccessible_ByName(t *testing.T) {
	var val string
	cfg := &Config{Title: "Pick"}
	items := []Item{
		{Name: "alpha", Desc: "alpha"},
		{Name: "beta", Desc: "beta"},
	}
	f := NewField(cfg, items, &val)

	var out bytes.Buffer
	in := strings.NewReader("alpha\n")
	err := f.RunAccessible(&out, in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "alpha" {
		t.Errorf("expected val = %q, got %q", "alpha", val)
	}
}

func TestField_RunAccessible_InvalidSelection(t *testing.T) {
	var val string
	cfg := &Config{Title: "Pick"}
	items := []Item{{Name: "alpha", Desc: "alpha"}}
	f := NewField(cfg, items, &val)

	var out bytes.Buffer
	in := strings.NewReader("bogus\n")
	err := f.RunAccessible(&out, in)
	if err == nil {
		t.Error("expected error for invalid selection")
	}
}
