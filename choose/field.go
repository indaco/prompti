package choose

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"github.com/indaco/prompti/internal/huhfield"
)

// Field is a huh.Field adapter wrapping the prompti choose component.
type Field struct {
	huhfield.Base
	inner model
	items []Item
	value *string
}

// NewField creates a new choose Field for use within a huh form.
// The selected item name is written into the string pointed to by value.
func NewField(cfg *Config, items []Item, value *string) *Field {
	c := *cfg
	c.setDefaults(len(items))
	return &Field{
		inner: c.initialModel(items),
		items: items,
		value: value,
		Base: huhfield.Base{
			Keymap: huh.NewDefaultKeyMap(),
		},
	}
}

// Init initialises the field.
func (f *Field) Init() tea.Cmd {
	return nil
}

// Update handles messages for the huh form integration.
func (f *Field) Update(msg tea.Msg) (huh.Model, tea.Cmd) {
	if !f.Focused {
		return f, nil
	}

	if msg, ok := msg.(tea.KeyPressMsg); ok {
		km := f.Keymap.Select

		switch {
		case key.Matches(msg, km.Prev):
			return f, huh.PrevField
		case key.Matches(msg, km.Next, km.Submit):
			// Select current item and advance.
			i, ok := f.inner.list.SelectedItem().(Item)
			if ok {
				f.inner.choice = i.Name
				f.syncValue()
			}
			return f, huh.NextField
		}
	}

	// Delegate all other messages to the inner model.
	updated, cmd := f.inner.Update(msg)
	f.inner = updated.(model)
	// Reset done flags so the model stays alive in form context.
	f.inner.submitted = false
	f.inner.quitting = false
	return f, cmd
}

// View renders the field.
func (f *Field) View() string {
	return f.inner.View().Content
}

// Focus focuses the field.
func (f *Field) Focus() tea.Cmd {
	f.Focused = true
	return nil
}

// Blur blurs the field.
func (f *Field) Blur() tea.Cmd {
	f.Focused = false
	f.syncValue()
	return nil
}

// KeyBinds returns the key bindings for this field.
func (f *Field) KeyBinds() []key.Binding {
	km := f.Keymap.Select
	return []key.Binding{km.Up, km.Down, km.Prev, km.Submit, km.Next}
}

// Run runs the field individually using huh's form runner.
func (f *Field) Run() error {
	return huh.Run(f)
}

// RunAccessible runs the field in accessible mode using plain text I/O.
func (f *Field) RunAccessible(w io.Writer, r io.Reader) error {
	_, _ = fmt.Fprintf(w, "%s\n", f.inner.list.Title)
	for i, item := range f.items {
		_, _ = fmt.Fprintf(w, "  %d. %s\n", i+1, item.Name)
	}
	_, _ = fmt.Fprint(w, "Enter choice: ")

	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		// Try number first.
		var idx int
		if _, err := fmt.Sscanf(input, "%d", &idx); err == nil && idx >= 1 && idx <= len(f.items) {
			if f.value != nil {
				*f.value = f.items[idx-1].Name
			}
			return nil
		}
		// Try name match.
		for _, item := range f.items {
			if strings.EqualFold(item.Name, input) {
				if f.value != nil {
					*f.value = item.Name
				}
				return nil
			}
		}
		return fmt.Errorf("invalid selection: %s", input)
	}
	return scanner.Err()
}

// WithTheme sets the theme on the field.
func (f *Field) WithTheme(theme huh.Theme) huh.Field {
	if f.Theme != nil {
		return f
	}
	f.Theme = theme
	return f
}

// WithKeyMap sets the keymap on the field.
func (f *Field) WithKeyMap(km *huh.KeyMap) huh.Field {
	f.Keymap = km
	return f
}

// WithWidth sets the width of the field.
func (f *Field) WithWidth(w int) huh.Field {
	f.Width = w
	f.inner.list.SetWidth(w)
	return f
}

// WithHeight sets the height of the field.
func (f *Field) WithHeight(h int) huh.Field {
	f.Height = h
	f.inner.list.SetHeight(h)
	return f
}

// WithPosition sets the position of the field within the form.
func (f *Field) WithPosition(p huh.FieldPosition) huh.Field {
	f.Position = p
	if f.Keymap != nil {
		f.Keymap.Select.Prev.SetEnabled(!p.IsFirst())
		f.Keymap.Select.Next.SetEnabled(!p.IsLast())
		f.Keymap.Select.Submit.SetEnabled(p.IsLast())
	}
	return f
}

// GetValue returns the current selected value.
func (f *Field) GetValue() any {
	return f.inner.choice
}

// syncValue writes the current selection to the external pointer.
func (f *Field) syncValue() {
	if f.value != nil && f.inner.choice != "" {
		*f.value = f.inner.choice
	} else if f.value != nil {
		// Use currently highlighted item if no explicit choice made.
		if i, ok := f.inner.list.SelectedItem().(Item); ok {
			*f.value = i.Name
		}
	}
}
