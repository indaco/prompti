package confirm

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/indaco/prompti"
)

const (
	okLabel     = "Yes"
	cancelLabel = "No"
	cursorLabel = ">"
	dividerChar = "/"
)

// Config represents the struct to configure the tui command.
type Config struct {
	// Mode selects the visual presentation. ModeDialog (default) renders a
	// bordered dialog box; ModeInline renders a compact single-line toggle.
	Mode Mode
	// Message is an optional body text displayed inside the dialog border.
	// Only used in ModeDialog.
	Message string
	// Question is the prompt text shown to the user.
	Question string
	// OkButtonLabel is the label for the affirmative option (default "Yes").
	OkButtonLabel string
	// CancelButtonLabel is the label for the negative option (default "No").
	CancelButtonLabel string
	// Cursor is the cursor character shown before the toggle options.
	// Only used in ModeInline (default ">").
	Cursor string
	// Divider is the separator between the two toggle options.
	// Only used in ModeInline (default "/").
	Divider string
	// Styles configures the visual appearance.
	Styles Styles
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults() {
	cfg.OkButtonLabel = setOkButtonLabel(cfg)
	cfg.CancelButtonLabel = setCancelButtonLabel(cfg)
	if cfg.Mode == ModeInline {
		cfg.Cursor = setCursor(cfg)
		cfg.Divider = setDivider(cfg)
	}
	cfg.Styles.setDefaults(cfg.Mode)
}

func (cfg *Config) initialModel() model {
	return model{
		mode:              cfg.Mode,
		message:           cfg.Message,
		question:          cfg.Question,
		okButtonLabel:     cfg.OkButtonLabel,
		cancelButtonLabel: cfg.CancelButtonLabel,
		confirmation:      true,
		cursor:            cfg.Cursor,
		divider:           cfg.Divider,
		// styles
		styles: cfg.Styles,
	}
}

// Run provides a shell script interface for prompting a user to confirm an
// action with an affirmative or negative answer.
func Run(cfg *Config) (bool, error) {
	c := *cfg
	c.setDefaults()
	p := tea.NewProgram(c.initialModel())
	m, err := p.Run()

	if err != nil {
		return false, fmt.Errorf("unable to run confirm: %w", err)
	}

	if m.(model).confirmation {
		return true, nil
	}

	return false, prompti.ErrCancelled
}

// =================================================================

func setOkButtonLabel(cfg *Config) string {
	if cfg.OkButtonLabel == "" {
		return okLabel
	}
	return cfg.OkButtonLabel
}

func setCancelButtonLabel(cfg *Config) string {
	if cfg.CancelButtonLabel == "" {
		return cancelLabel
	}
	return cfg.CancelButtonLabel
}

func setCursor(cfg *Config) string {
	if cfg.Cursor == "" {
		return cursorLabel
	}
	return cfg.Cursor
}

func setDivider(cfg *Config) string {
	if cfg.Divider == "" {
		return dividerChar
	}
	return cfg.Divider
}
