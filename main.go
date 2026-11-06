package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).PaddingTop(0).
			PaddingBottom(0).
			PaddingLeft(2).
			PaddingRight(2)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			MarginTop(1)

	subStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Italic(true)
)

type model struct {
	cursor int
	items  []string
}

func initialModel() model {
	return model{
		items: []string{"Containers", "Images", "System Metrics", "Logs"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case "enter":
		}
	}

	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("Docktop - Terminal Docker Monitor") + "\n\n"
	s += "Select View:\n\n"

	for i, item := range m.items {
		cursor := " "
		itemText := item

		if m.cursor == i {
			cursor = "->"
			itemText = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#EE6FF8")).Render(item)
		}
		s += fmt.Sprintf(" %s %s\n", cursor, itemText)
	}
	s += infoStyle.Render(fmt.Sprintf("\nActive View:[%s]", m.items[m.cursor])) + "\n"
	s += subStyle.Render("\n(Press ↑/↓ or j/k to navigate • q to quit)\n")

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
