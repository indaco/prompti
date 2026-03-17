package toggle

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/indaco/prompti"
)

// Config represents the struct to configure the tui command.
type Config struct {
	Question          string
	OkButtonLabel     string
	CancelButtonLabel string
	Cursor            string
	Divider           string
	// styles
	Styles Styles
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults() {
	cfg.OkButtonLabel = setOkButtonLabel(cfg)
	cfg.CancelButtonLabel = setCancelButtonLabel(cfg)
	cfg.Cursor = setCursor(cfg)
	cfg.Divider = setDivider(cfg)
	cfg.Styles.setDefaults()
}

func (cfg *Config) initialModel() model {
	return model{
		question:          cfg.Question,
		okButtonLabel:     cfg.OkButtonLabel,
		cancelButtonLabel: cfg.CancelButtonLabel,
		confirmation:      true,
		cursor:            cfg.Cursor,
		divider:           cfg.Divider,
		styles:            cfg.Styles,
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
		return false, fmt.Errorf("unable to run toggle: %w", err)
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
		return divider
	}
	return cfg.Divider
}
