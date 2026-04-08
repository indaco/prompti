package confirm

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

// Field is a huh.Field adapter wrapping the prompti confirm component.
type Field struct {
	huhfield.Base
	inner model
	value *bool
}

// NewField creates a new confirm Field for use within a huh form.
// The result is written into the bool pointed to by value.
func NewField(cfg *Config, value *bool) *Field {
	c := *cfg
	c.setDefaults()
	return &Field{
		inner: c.initialModel(),
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
		km := f.Keymap.Confirm

		switch {
		case key.Matches(msg, km.Toggle):
			f.inner.confirmation = !f.inner.confirmation
			return f, nil
		case key.Matches(msg, km.Prev):
			f.syncValue()
			return f, huh.PrevField
		case key.Matches(msg, km.Accept):
			f.inner.confirmation = true
			f.syncValue()
			return f, huh.NextField
		case key.Matches(msg, km.Reject):
			f.inner.confirmation = false
			f.syncValue()
			return f, huh.NextField
		case key.Matches(msg, km.Next, km.Submit):
			f.syncValue()
			return f, huh.NextField
		}
	}

	return f, nil
}

// View renders the field.
func (f *Field) View() string {
	return f.inner.View().Content
}

// Focus focuses the field.
func (f *Field) Focus() tea.Cmd {
	f.Focused = true
	// Reset quitting so the view renders properly.
	f.inner.quitting = false
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
	km := f.Keymap.Confirm
	return []key.Binding{km.Toggle, km.Prev, km.Submit, km.Next, km.Accept, km.Reject}
}

// Run runs the field individually using huh's form runner.
func (f *Field) Run() error {
	return huh.Run(f)
}

// RunAccessible runs the field in accessible mode using plain text I/O.
func (f *Field) RunAccessible(w io.Writer, r io.Reader) error {
	prompt := f.inner.question
	if prompt == "" {
		prompt = f.inner.message
	}
	defaultHint := "[y/N]"
	if f.inner.confirmation {
		defaultHint = "[Y/n]"
	}
	_, _ = fmt.Fprintf(w, "%s %s ", prompt, defaultHint)

	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		input := strings.TrimSpace(strings.ToLower(scanner.Text()))
		switch input {
		case "y", "yes":
			if f.value != nil {
				*f.value = true
			}
		case "n", "no":
			if f.value != nil {
				*f.value = false
			}
		case "":
			// Use default (current confirmation state).
			if f.value != nil {
				*f.value = f.inner.confirmation
			}
		default:
			return fmt.Errorf("invalid input: %s", input)
		}
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
	return f
}

// WithHeight sets the height of the field.
func (f *Field) WithHeight(h int) huh.Field {
	f.Height = h
	return f
}

// WithPosition sets the position of the field within the form.
func (f *Field) WithPosition(p huh.FieldPosition) huh.Field {
	f.Position = p
	if f.Keymap != nil {
		f.Keymap.Confirm.Prev.SetEnabled(!p.IsFirst())
		f.Keymap.Confirm.Next.SetEnabled(!p.IsLast())
		f.Keymap.Confirm.Submit.SetEnabled(p.IsLast())
	}
	return f
}

// GetValue returns the current confirmation value.
func (f *Field) GetValue() any {
	return f.inner.confirmation
}

// syncValue writes the current confirmation to the external pointer.
func (f *Field) syncValue() {
	if f.value != nil {
		*f.value = f.inner.confirmation
	}
}
