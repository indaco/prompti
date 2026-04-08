// Package confirm provides an interactive yes/no confirmation prompt with two
// visual modes: a bordered dialog box (ModeDialog, the default) and an inline
// single-line toggle (ModeInline).
package confirm

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

// Mode selects the visual presentation of the confirm prompt.
type Mode int

const (
	// ModeDialog renders a bordered dialog box with an optional message body.
	ModeDialog Mode = iota
	// ModeInline renders a compact, single-line toggle with cursor and divider.
	ModeInline
)

type model struct {
	mode              Mode
	message           string
	question          string
	okButtonLabel     string
	cancelButtonLabel string
	quitting          bool
	submitted         bool
	confirmation      bool
	// inline-mode strings
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
			m.submitted = true
			return m, tea.Quit
		case "y":
			m.quitting = true
			m.submitted = true
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

	if m.mode == ModeInline {
		return m.viewInline()
	}

	return m.viewDialog()
}

func (m model) viewDialog() tea.View {
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

func (m model) viewInline() tea.View {
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
