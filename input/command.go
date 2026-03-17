package input

import (
	"fmt"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/indaco/prompti"
)

// Config represents the struct to configure the tui input.
type Config struct {
	Message      string
	Placeholder  string
	Initial      string
	ErrorMsg     string
	Password     bool
	ValidateFunc ValidateFunc

	Styles Styles
}

// setDefaults sets default values for Config struct.
func (cfg *Config) setDefaults() {
	cfg.Placeholder = setDefaultPlaceholderMsg(cfg)
	cfg.Styles.setDefaults()
}

func (cfg *Config) initialModel() model {
	ti := textinput.New()

	ti.Focus()
	ti.SetWidth(80)
	ti.Prompt = promptMessage(cfg.Message, cfg.Styles.PrefixIcon, cfg.Styles.PrefixIconColor)

	ti.Placeholder = cfg.Placeholder

	styles := ti.Styles()
	styles.Focused.Text = cfg.Styles.TextStyle
	styles.Focused.Prompt = prefixIconStyle(cfg.Styles.PrefixIconColor)
	styles.Focused.Placeholder = cfg.Styles.PlaceholderStyle
	ti.SetStyles(styles)

	if cfg.Password {
		ti.EchoMode = textinput.EchoPassword
		ti.EchoCharacter = '*'
	}

	return model{
		textInput:    ti,
		message:      cfg.Message,
		placeholder:  cfg.Placeholder,
		initial:      cfg.Initial,
		err:          nil,
		validateFunc: cfg.ValidateFunc,
	}
}

// Run is used to prompt an input the user and retrieve the value.
func Run(cfg *Config) (string, error) {
	c := *cfg
	c.setDefaults()
	p := tea.NewProgram(c.initialModel())

	tm, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run input: %w", err)
	}
	m := tm.(model)

	if m.quitting {
		return "", prompti.ErrCancelled
	}

	if m.textInput.Value() == "" {
		return "", prompti.ErrEmpty
	}

	return m.textInput.Value(), m.err
}

// =================================================================

func setDefaultPlaceholderMsg(cfg *Config) string {
	if cfg.Initial != "" {
		return fmt.Sprintf("%s (Default: %s)", cfg.Placeholder, cfg.Initial)
	}
	return cfg.Placeholder
}
