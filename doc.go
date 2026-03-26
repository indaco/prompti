// Package prompti provides a collection of interactive terminal UI prompts
// built on the charmbracelet bubbletea framework.
//
// It includes the following sub-packages:
//
//   - choose: single-select list prompt
//   - input: text input with optional validation
//   - confirm: yes/no confirmation dialog
//   - toggle: inline yes/no toggle
//   - detail: collapsible detail/summary section
//   - progressbar: animated progress bar
//
// Each sub-package exposes a Config struct and a Run function.
// Sentinel errors [ErrCancelled] and [ErrEmpty] allow callers to
// distinguish user cancellation from other failures.
package prompti
