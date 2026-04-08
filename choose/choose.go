// Package choose provides an interactive single-select list prompt built on bubbletea.
package choose

import (
	"fmt"
	"io"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/internal/theme"
)

type errMsg error

// Item represents an entry for choose (select) list.
type Item struct {
	Name string
	Desc string
}

// FilterValue returns the current value of the filter.
func (i Item) FilterValue() string { return i.Name }

// String returns the item name.
func (i *Item) String() string {
	return i.Name
}

// GetItemsKeys returns a slice of strings representing the item names.
func GetItemsKeys(items []Item) []string {
	res := []string{}
	for _, v := range items {
		res = append(res, v.Name)
	}
	return res
}

// ToItems converts a slice of strings into a slice of Item.
func ToItems(items []string) []Item {
	res := []Item{}
	for _, v := range items {
		res = append(res, Item{
			Name: v,
			Desc: v,
		})
	}
	return res
}

// toListItems converts a slice of Item to a slice of list.Item.
func toListItems(items []Item) []list.Item {
	res := make([]list.Item, len(items))
	for i, v := range items {
		res[i] = v
	}
	return res
}

type itemDelegate struct {
	ItemIcon          string
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := d.ItemIcon + theme.Whitespace + i.Desc

	fn := d.ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return d.SelectedItemStyle.Render(strings.Join(s, ""))
		}
	}

	_, _ = fmt.Fprint(w, fn(str))
}

type model struct {
	list      list.Model
	err       error
	choice    string
	quitting  bool
	submitted bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.choice = i.Name
			}
			m.submitted = true
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() tea.View {
	if m.choice != "" {
		return tea.NewView(m.choice)
	}
	return tea.NewView(m.list.View())
}
