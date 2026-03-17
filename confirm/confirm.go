// Package confirm provides an interactive yes/no confirmation dialog with a styled message box.
package confirm

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	message           string
	question          string
	okButtonLabel     string
	cancelButtonLabel string
	quitting          bool
	confirmation      bool

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

	var aff, neg string

	if m.confirmation {
		aff = m.styles.ActiveButtonStyle.Render(m.okButtonLabel)
		neg = m.styles.ButtonStyle.Render(m.cancelButtonLabel)
	} else {
		aff = m.styles.ButtonStyle.Render(m.okButtonLabel)
		neg = m.styles.ActiveButtonStyle.Render(m.cancelButtonLabel)
	}

	message := lipgloss.NewStyle().Render(m.message)
	question := m.styles.QuestionStyle.Render(m.question)
	buttons := lipgloss.JoinHorizontal(lipgloss.Left, aff, neg)

	var ui string
	if message != "" {
		ui = m.styles.DialogStyle.Render(
			lipgloss.JoinVertical(lipgloss.Center, message, "\n", question, buttons))
	} else {
		ui = m.styles.DialogStyle.Render(
			lipgloss.JoinVertical(lipgloss.Center, question, buttons))
	}

	return tea.NewView(lipgloss.NewStyle().Render(ui))
}
