// Package progressbar provides an animated progress bar that iterates over a list of items.
package progressbar

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

const (
	padding  = 2
	maxWidth = 80
)

var (
	defaultCurrentItemNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle                   = lipgloss.NewStyle().Margin(1, 2)
	checkMark                   = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

// IncrementMsg is the msg received on progress.
type IncrementMsg string

// IncrementErrMsg is msg struct when error.
type IncrementErrMsg struct{ Err error }

// Error implements the error interface on the message.
func (e IncrementErrMsg) Error() string { return e.Err.Error() }

type model struct {
	items                []string
	index                int
	onProgressCmd        func(string) tea.Cmd
	onProgressMsg        string
	runConcurrently      bool
	showLabel            bool
	currentItemNameStyle lipgloss.Style
	onCompletedMsg       string
	progress             progress.Model
	done                 bool
	err                  error
}

func (m model) Init() tea.Cmd {
	if m.runConcurrently {
		return tea.Batch(m.onProgressCmd(m.items[m.index]))
	}

	return tea.Sequence(m.onProgressCmd(m.items[m.index]))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		w := min(msg.Width-padding*2-4, maxWidth)
		m.progress.SetWidth(w)
		return m, nil

	case IncrementErrMsg:
		m.err = msg.Err
		return m, tea.Quit

	case IncrementMsg:
		if m.index >= len(m.items)-1 {
			m.done = true
			return m, tea.Quit
		}

		// Update progressbar
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.items)-1))
		m.index++

		if m.runConcurrently {
			return m, tea.Batch(
				progressCmd,
				m.onProgressCmd(m.items[m.index]),
			)
		}

		return m, tea.Sequence(
			progressCmd,
			m.onProgressCmd(m.items[m.index]),
		)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		var cmd tea.Cmd
		m.progress, cmd = m.progress.Update(msg)
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() tea.View {
	if m.err != nil {
		return tea.NewView(fmt.Sprintf("Error: %s\n", m.err))
	}

	if m.done {
		return tea.NewView(doneStyle.Render(fmt.Sprintf("%s %s", checkMark.String(), m.onCompletedMsg)))
	}

	total := len(m.items)

	pad := strings.Repeat(" ", padding)
	w := lipgloss.Width(fmt.Sprintf("%d", total))
	counter := fmt.Sprintf(" %*d/%*d", w, m.index, w, total)
	progressBar := m.progress.View()

	itemName := formatItemName(&m, m.items[m.index])
	infoStr := formatMsg(&m, itemName)

	return tea.NewView("\n" +
		infoStr +
		pad + progressBar + counter)
}

//=============================================================================

func formatMsg(m *model, item string) string {
	if m.showLabel {
		return fmt.Sprintf("%s %s\n", m.onProgressMsg, lipgloss.NewStyle().Render(item))
	}
	return ""
}

func formatItemName(m *model, item string) string {
	if !theme.IsZeroStyle(m.currentItemNameStyle) {
		return m.currentItemNameStyle.Render(item)
	}
	return defaultCurrentItemNameStyle.Render(item)
}
