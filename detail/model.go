// Package detail provides a collapsible detail/summary prompt built on bubbletea.
package detail

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

type model struct {
	summary  string
	content  string
	expanded bool
	quitting bool
	// strings
	collapsedIndicator string
	expandedIndicator  string
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
		case "ctrl+c", "esc", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter", "space":
			m.expanded = !m.expanded
		case "left", "h", "up", "k":
			m.expanded = false
		case "right", "l", "down", "j":
			m.expanded = true
		case "y":
			m.quitting = true
			m.expanded = true
			return m, tea.Quit
		case "n":
			m.quitting = true
			m.expanded = false
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	if m.quitting {
		return tea.NewView("")
	}

	var indicator string
	if m.expanded {
		indicator = m.expandedIndicator
	} else {
		indicator = m.collapsedIndicator
	}

	header := summaryMessage(
		m.styles.PrefixIcon,
		m.styles.PrefixIconColor,
		m.styles.SummaryStyle,
		m.styles.IndicatorStyle,
		indicator,
		m.summary,
	)

	var ui string
	if m.expanded {
		content := m.styles.ContentStyle.Render(m.content)
		ui = m.styles.DialogStyle.Render(
			lipgloss.JoinVertical(lipgloss.Left, header, content))
	} else {
		ui = m.styles.DialogStyle.Render(header)
	}

	hint := m.styles.HintStyle.Render(hintMessage)
	output := lipgloss.JoinVertical(lipgloss.Left, ui, hint)

	return tea.NewView(lipgloss.NewStyle().Render(output))
}

var hintMessage = theme.Whitespace + "enter/space: toggle" + theme.Whitespace + theme.PromptMark + theme.Whitespace + "q/esc: quit"
