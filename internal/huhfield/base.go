// Package huhfield provides a shared base for prompti huh.Field adapters.
package huhfield

import (
	"charm.land/huh/v2"
)

// Base holds common state shared by all prompti huh.Field adapters.
// It is not intended for direct use; embed it in each adapter struct.
type Base struct {
	Key      string
	Focused  bool
	Width    int
	Height   int
	Err      error
	Theme    huh.Theme
	Keymap   *huh.KeyMap
	Position huh.FieldPosition
}

// GetKey returns the field's key.
func (b *Base) GetKey() string {
	return b.Key
}

// Error returns the field's current error.
func (b *Base) Error() error {
	return b.Err
}

// Skip reports whether this field should be skipped.
func (b *Base) Skip() bool {
	return false
}

// Zoom reports whether this field should be zoomed.
func (b *Base) Zoom() bool {
	return false
}
