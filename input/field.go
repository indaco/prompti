package input

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

// Field is a huh.Field adapter wrapping the prompti input component.
type Field struct {
	huhfield.Base
	inner model
	value *string
}

// NewField creates a new input Field for use within a huh form.
// The result is written into the string pointed to by value.
func NewField(cfg *Config, value *string) *Field {
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
		km := f.Keymap.Input

		switch {
		case key.Matches(msg, km.Prev):
			f.syncValue()
			return f, huh.PrevField
		case key.Matches(msg, km.Next, km.Submit):
			// Delegate enter to inner model for validation.
			updated, _ := f.inner.Update(msg)
			f.inner = updated.(model)
			if f.inner.isDone() && !f.inner.isCancelled() {
				f.syncValue()
				return f, huh.NextField
			}
			// Validation failed, stay on field.
			f.Err = f.inner.err
			return f, nil
		}
	}

	// Delegate all other messages to the inner model.
	updated, cmd := f.inner.Update(msg)
	f.inner = updated.(model)
	// Reset done/quitting flags set by inner model so it stays alive in form context.
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
	km := f.Keymap.Input
	return []key.Binding{km.Prev, km.Submit, km.Next}
}

// Run runs the field individually using huh's form runner.
func (f *Field) Run() error {
	return huh.Run(f)
}

// RunAccessible runs the field in accessible mode using plain text I/O.
func (f *Field) RunAccessible(w io.Writer, r io.Reader) error {
	prompt := f.inner.message
	if f.inner.placeholder != "" {
		prompt += " (" + f.inner.placeholder + ")"
	}
	prompt += ": "
	_, _ = fmt.Fprint(w, prompt)

	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		val := strings.TrimSpace(scanner.Text())
		if val == "" && f.inner.initial != "" {
			val = f.inner.initial
		}
		if f.inner.validateFunc != nil {
			if err := f.inner.validateFunc(val); err != nil {
				return err
			}
		}
		if f.value != nil {
			*f.value = val
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
	f.inner.textInput.SetWidth(w)
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
		f.Keymap.Input.Prev.SetEnabled(!p.IsFirst())
		f.Keymap.Input.Next.SetEnabled(!p.IsLast())
		f.Keymap.Input.Submit.SetEnabled(p.IsLast())
	}
	return f
}

// GetValue returns the current value of the field.
func (f *Field) GetValue() any {
	return f.inner.textInput.Value()
}

// syncValue writes the current input value to the external pointer.
func (f *Field) syncValue() {
	if f.value != nil {
		*f.value = f.inner.textInput.Value()
	}
}
