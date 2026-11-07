package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sessionState int

const (
	viewMenu sessionState = iota
	viewContainers
)

var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 2)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			MarginTop(1)

	subStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Italic(true)
)

type model struct {
	state          sessionState
	menuCursor     int
	menuItems      []string
	containerTable table.Model
}

func initialModel() model {
	menu := []string{"Containers", "Images", "System Metrics", "Logs"}

	columns := []table.Column{
		{Title: "ID", Width: 14},
		{Title: "NAME", Width: 20},
		{Title: "IMAGE", Width: 24},
		{Title: "STATUS", Width: 22},
		{Title: "PORTS", Width: 25},
	}

	rows := []table.Row{
		{"d4fe8ab80930", "web-laravel-app", "web-app:latest", "Up 4 hours (healthy)", "0.0.0.0:80->80/tcp"},
		{"a6aa0084b4c6", "mysql-database", "mysql:8.0", "Up 4 hours", "0.0.0.0:3306->3306/tcp"},
		{"f301bd707704", "redis-cache", "redis:7-alpine", "Up 2 hours", "0.0.0.0:6379->6379/tcp"},
		{"b0a0ccfb9ef9", "react-frontend", "node:20-alpine", "Exited (0) 1 hour ago", "-"},
		{"09f5dc4ef198", "minikube-dev", "kicbase:v0.0.50", "Exited (255) 2 days ago", "-"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#EE6FF8")).
		Bold(true)
	t.SetStyles(s)

	return model{
		state:          viewMenu,
		menuItems:      menu,
		containerTable: t,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

		switch m.state {
		case viewMenu:
			switch msg.String() {
			case "up", "k":
				if m.menuCursor > 0 {
					m.menuCursor--
				}
			case "down", "j":
				if m.menuCursor < len(m.menuItems)-1 {
					m.menuCursor++
				}
			case "enter":
				if m.menuCursor == 0 {
					m.state = viewContainers
				}
			}

		case viewContainers:
			switch msg.String() {
			case "esc", "b":
				m.state = viewMenu
				return m, nil
			}
			m.containerTable, cmd = m.containerTable.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {
	switch m.state {
	case viewContainers:
		header := titleStyle.Render("Docktop • Containers") + "\n\n"
		tableView := baseStyle.Render(m.containerTable.View()) + "\n"
		footer := subStyle.Render("\n(Press ↑/↓ or j/k to navigate • Esc/b to back • q to quit)\n")
		return header + tableView + footer

	default:
		s := titleStyle.Render("Docktop - Terminal Docker Monitor") + "\n\n"
		s += "Select View:\n\n"

		for i, item := range m.menuItems {
			cursor := " "
			itemText := item

			if m.menuCursor == i {
				cursor = "->"
				itemText = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#EE6FF8")).Render(item)
			}
			s += fmt.Sprintf(" %s %s\n", cursor, itemText)
		}
		s += infoStyle.Render(fmt.Sprintf("\nActive View: [%s]", m.menuItems[m.menuCursor])) + "\n"
		s += subStyle.Render("\n(Press ↑/↓ or j/k to navigate • Enter to select • q to quit)\n")
		return s
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running docktop: %v\n", err)
		os.Exit(1)
	}
}
