package choose

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/indaco/prompti"
)

// Config represents the struct to configure the tui command.
type Config struct {
	Title           string
	ListHeight      int
	DefaultWidth    int
	ErrorMsg        string
	ShowHelp        bool
	ShowStatusBar   bool
	EnableFiltering bool // ShowHelp must be true otherwise you will now see the help
	// styles
	Styles Styles
}

// setDefaults sets default values for Config if not present.
func (cfg *Config) setDefaults(numOfItems int) {
	cfg.ListHeight = setListHeight(cfg, numOfItems)
	cfg.Styles.setDefaults()
}

func (cfg *Config) initialModel(items []Item) model {
	listItems := toListItems(items)
	itemDelegate := itemDelegate{
		ItemIcon:          cfg.Styles.ItemIcon,
		ItemStyle:         cfg.Styles.ItemStyle,
		SelectedItemStyle: cfg.Styles.SelectedItemStyle,
	}
	l := list.New(listItems, itemDelegate, cfg.DefaultWidth, cfg.ListHeight)
	l.Title = titleMessage(cfg.Styles.PrefixIcon, cfg.Styles.PrefixIconColor, cfg.Styles.TitleStyle, cfg.Title)
	l.Styles.Title = cfg.Styles.TitleStyle
	l.Styles.TitleBar = cfg.Styles.TitleBarStyle
	l.SetShowHelp(cfg.ShowHelp)
	l.SetShowStatusBar(cfg.ShowStatusBar)
	l.SetFilteringEnabled(cfg.EnableFiltering)

	return model{
		list: l,
	}
}

// Run is used to prompt a list of available options to the user and retrieve the selection.
func Run(cfg *Config, items []Item) (string, error) {
	c := *cfg
	c.setDefaults(len(items))
	p := tea.NewProgram(c.initialModel(items))

	tm, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run choose: %w", err)
	}
	m := tm.(model)

	if m.quitting {
		return "", prompti.ErrCancelled
	}

	return m.list.SelectedItem().FilterValue(), nil
}

// =================================================================

func setListHeight(cfg *Config, numOfItems int) int {
	if cfg.ListHeight == 0 {
		if cfg.ShowHelp || cfg.ShowStatusBar || cfg.EnableFiltering {
			return numOfItems * 4
		}
		return numOfItems * 3
	}
	return cfg.ListHeight
}
