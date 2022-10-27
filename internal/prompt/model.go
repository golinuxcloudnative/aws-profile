package prompt

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 2).
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4"))
	shellTextStyle = lipgloss.NewStyle().
			Margin(1, 0, 2, 2).
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4"))
)

type item struct {
	name, region, url, role, accountId string
}

func (i item) Title() string { return i.name }
func (i item) Description() string {
	return fmt.Sprintf("account_id: %s\trole: %s\tregion: %s\turl: %s", i.accountId, i.role, i.region, i.url)
}
func (i item) FilterValue() string { return i.name }

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i.name)
			}

			return m, tea.Quit

		}
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		msg := fmt.Sprintf("Starting a shell with AWS profile %s", m.choice)
		return shellTextStyle.Render(msg)
	}
	if m.quitting {
		return quitTextStyle.Render("Exiting AWS Profile Manager")
	}
	return "\n" + m.list.View()
}
