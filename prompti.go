package prompti

import "errors"

var (
	// ErrCancelled is returned when the user cancels the prompt (ctrl+c, esc).
	ErrCancelled = errors.New("prompti: operation cancelled")

	// ErrEmpty is returned when the user submits empty input.
	ErrEmpty = errors.New("prompti: empty input")
)
