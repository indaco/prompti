// Package input provides an interactive text input prompt with optional validation.
package input

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

type model struct {
	textInput    textinput.Model
	message      string
	placeholder  string
	initial      string
	err          error
	quitting     bool
	validateFunc ValidateFunc
}

func (m model) Init() tea.Cmd { return nil }
func (m model) View() tea.View {
	if m.textInput.Value() == "" {
		return tea.NewView(m.textInput.View())
	}

	if m.err != nil {
		return tea.NewView(lipgloss.NewStyle().Render(
			lipgloss.JoinVertical(lipgloss.Left,
				m.textInput.View(),
				errorMessage(m.err.Error()))))
	}

	return tea.NewView(m.textInput.View())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textInput.SetWidth(msg.Width)
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.initial != "" && m.textInput.Value() == "" {
				if m.validateFunc != nil {
					m.err = m.validateFunc(m.initial)
				}
				m.textInput.SetValue(m.initial)
			}

			if m.err == nil || m.validateFunc == nil || m.validateFunc(m.textInput.Value()) == nil {
				m.err = nil
				return m, tea.Quit
			}
		}
		m.textInput, cmd = m.textInput.Update(msg)
		if m.validateFunc != nil {
			m.err = m.validateFunc(m.textInput.Value())
		}
	case error:
		m.err = msg
		return m, nil
	}
	return m, cmd
}

var errorMessage = func(s string) string {
	return errorStyle.Render(
		lipgloss.JoinHorizontal(lipgloss.Center,
			cancelMarkStyle.Render(theme.CancelMark), s))
}
