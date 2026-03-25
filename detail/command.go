package detail

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/indaco/prompti"
)

// Config represents the struct to configure the tui command.
type Config struct {
	Summary            string
	Content            string
	CollapsedIndicator string
	ExpandedIndicator  string
	// styles
	Styles Styles
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults() {
	cfg.CollapsedIndicator = setCollapsedIndicator(cfg)
	cfg.ExpandedIndicator = setExpandedIndicator(cfg)
	cfg.Styles.setDefaults()
}

func (cfg *Config) initialModel() model {
	return model{
		summary:            cfg.Summary,
		content:            cfg.Content,
		expanded:           false,
		collapsedIndicator: cfg.CollapsedIndicator,
		expandedIndicator:  cfg.ExpandedIndicator,
		styles:             cfg.Styles,
	}
}

// Run provides a shell script interface for prompting a user with a
// collapsible detail/summary section. It returns whether the detail
// was expanded when the user quit.
func Run(cfg *Config) (bool, error) {
	c := *cfg
	c.setDefaults()
	p := tea.NewProgram(c.initialModel())
	m, err := p.Run()

	if err != nil {
		return false, fmt.Errorf("unable to run detail: %w", err)
	}

	result := m.(model)
	if result.quitting {
		return result.expanded, prompti.ErrCancelled
	}

	return result.expanded, nil
}

// =================================================================

func setCollapsedIndicator(cfg *Config) string {
	if cfg.CollapsedIndicator == "" {
		return collapsedIndicatorDefault
	}
	return cfg.CollapsedIndicator
}

func setExpandedIndicator(cfg *Config) string {
	if cfg.ExpandedIndicator == "" {
		return expandedIndicatorDefault
	}
	return cfg.ExpandedIndicator
}
