// Package toggle provides an inline yes/no toggle prompt built on bubbletea.
package toggle

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

type model struct {
	question          string
	okButtonLabel     string
	cancelButtonLabel string
	quitting          bool
	confirmation      bool
	// strings
	cursor  string
	divider string
	// styles
	styles Styles
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q", "n":
			m.confirmation = false
			m.quitting = true
			return m, tea.Quit
		case "left", "h", "ctrl+p", "tab",
			"right", "l", "ctrl+n", "shift+tab":
			m.confirmation = !m.confirmation
		case "enter":
			m.quitting = true
			return m, tea.Quit
		case "y":
			m.quitting = true
			m.confirmation = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	if m.quitting {
		return tea.NewView("")
	}

	var ok, cancel string

	if m.confirmation {
		ok = m.styles.ActiveButtonStyle.Render(m.okButtonLabel)
		cancel = m.styles.ButtonStyle.Render(m.cancelButtonLabel)
	} else {
		ok = m.styles.ButtonStyle.Render(m.okButtonLabel)
		cancel = m.styles.ActiveButtonStyle.Render(m.cancelButtonLabel)
	}

	question := m.styles.QuestionStyle.Render(prefixIconStyle(m.styles.PrefixIconColor).Render(m.styles.PrefixIcon) + m.question)
	prompt := lipgloss.NewStyle().Foreground(theme.Cyan).PaddingRight(1).Render(m.cursor)
	buttons := lipgloss.JoinHorizontal(0.5, ok, m.styles.DividerStyle.Render(m.divider), cancel)
	ui := m.styles.DialogStyle.Render(lipgloss.JoinHorizontal(lipgloss.Center, question, prompt, buttons))

	return tea.NewView(lipgloss.NewStyle().Render(ui))
}
