package progressbar

import (
	"math/rand/v2"
	"time"

	tea "charm.land/bubbletea/v2"
)

// Config represents the struct to configure the tui command.
type Config struct {
	Items         []string
	OnProgressCmd func(string) tea.Cmd
	// If true, the progressbar runs commands concurrently (tea.Batch) else in order (tea.Sequence).
	OnProgressMsg   string
	OnCompletedMsg  string
	RunConcurrently bool
	Styles          Styles
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults() {
	if cfg.OnCompletedMsg == "" {
		cfg.OnCompletedMsg = "Done!"
	}

	if cfg.OnProgressCmd == nil {
		cfg.OnProgressCmd = func(item string) tea.Cmd {
			// This is where you'd do i/o stuff to download and install packages. In
			// our case we're just pausing for a moment to simulate the process.
			d := time.Millisecond * time.Duration(rand.IntN(500)) //nolint:gosec // non-cryptographic use for demo delay
			return tea.Tick(d, func(t time.Time) tea.Msg {
				return IncrementMsg(item)
			})
		}
	}
}

func (cfg *Config) initialModel() model {
	return model{
		items:                cfg.Items,
		onProgressCmd:        cfg.OnProgressCmd,
		onProgressMsg:        cfg.OnProgressMsg,
		onCompletedMsg:       cfg.OnCompletedMsg,
		runConcurrently:      cfg.RunConcurrently,
		showLabel:            cfg.Styles.ShowLabel,
		currentItemNameStyle: cfg.Styles.CurrentItemStyle,
		progress:             newProgressBarModel(cfg),
	}
}

// Run displays an animated progress bar that iterates over the configured items.
func Run(cfg *Config) error {
	c := *cfg
	c.setDefaults()
	p := tea.NewProgram(c.initialModel())
	_, err := p.Run()
	return err
}
