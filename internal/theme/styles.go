package theme

import (
	"reflect"

	"charm.land/lipgloss/v2"
)

// IsZeroStyle reports whether s is a zero-value lipgloss.Style.
func IsZeroStyle(s lipgloss.Style) bool {
	return reflect.DeepEqual(s, lipgloss.Style{})
}

// IsZeroBorder reports whether b is a zero-value lipgloss.Border.
func IsZeroBorder(b lipgloss.Border) bool {
	return reflect.DeepEqual(b, lipgloss.Border{})
}
