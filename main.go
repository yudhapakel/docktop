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
	viewImages
	viewMetrics
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

	cardStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1, 2).
			MarginRight(2)

	metricLabelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#A0A0A0")).
				Bold(true)

	metricValStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EE6FF8")).
			Bold(true)
)

type model struct {
	state          sessionState
	menuCursor     int
	menuItems      []string
	containerTable table.Model
	imageTable     table.Model
}

func initialModel() model {
	menu := []string{"Containers", "Images", "System Metrics", "Logs"}

	containerCols := []table.Column{
		{Title: "ID", Width: 14},
		{Title: "NAME", Width: 20},
		{Title: "IMAGE", Width: 24},
		{Title: "STATUS", Width: 22},
		{Title: "PORTS", Width: 25},
	}

	containerRows := []table.Row{
		{"d4fe8ab80930", "web-laravel-app", "web-app:latest", "Up 4 hours (healthy)", "0.0.0.0:80->80/tcp"},
		{"a6aa0084b4c6", "mysql-database", "mysql:8.0", "Up 4 hours", "0.0.0.0:3306->3306/tcp"},
		{"f301bd707704", "redis-cache", "redis:7-alpine", "Up 2 hours", "0.0.0.0:6379->6379/tcp"},
		{"b0a0ccfb9ef9", "react-frontend", "node:20-alpine", "Exited (0) 1 hour ago", "-"},
		{"09f5dc4ef198", "minikube-dev", "kicbase:v0.0.50", "Exited (255) 2 days ago", "-"},
	}

	tContainers := table.New(
		table.WithColumns(containerCols),
		table.WithRows(containerRows),
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
	tContainers.SetStyles(s)

	imageCols := []table.Column{
		{Title: "REPOSITORY", Width: 22},
		{Title: "TAG", Width: 12},
		{Title: "IMAGE ID", Width: 16},
		{Title: "CREATED", Width: 18},
		{Title: "SIZE", Width: 12},
	}

	imageRows := []table.Row{
		{"web-app", "latest", "7b2c90a18f4e", "3 days ago", "184 MB"},
		{"mysql", "8.0", "c3d82f1b490a", "2 weeks ago", "544 MB"},
		{"redis", "7-alpine", "e10a29bd4851", "1 month ago", "36.2 MB"},
		{"node", "20-alpine", "a58b472e903f", "2 weeks ago", "178 MB"},
		{"kicbase", "v0.0.50", "9fa38b2910cd", "3 months ago", "1.15 GB"},
	}

	tImages := table.New(
		table.WithColumns(imageCols),
		table.WithRows(imageRows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	tImages.SetStyles(s)

	return model{
		state:          viewMenu,
		menuItems:      menu,
		containerTable: tContainers,
		imageTable:     tImages,
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
				switch m.menuCursor {
				case 0:
					m.state = viewContainers
				case 1:
					m.state = viewImages
				case 2:
					m.state = viewMetrics
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

		case viewImages:
			switch msg.String() {
			case "esc", "b":
				m.state = viewMenu
				return m, nil
			}
			m.imageTable, cmd = m.imageTable.Update(msg)
			return m, cmd

		case viewMetrics:
			switch msg.String() {
			case "esc", "b":
				m.state = viewMenu
				return m, nil
			}
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

	case viewImages:
		header := titleStyle.Render("Docktop • Images") + "\n\n"
		tableView := baseStyle.Render(m.imageTable.View()) + "\n"
		footer := subStyle.Render("\n(Press ↑/↓ or j/k to navigate • Esc/b to back • q to quit)\n")
		return header + tableView + footer

	case viewMetrics:
		header := titleStyle.Render("Docktop • System Metrics") + "\n\n"

		cpuCard := cardStyle.Render(
			metricLabelStyle.Render("CPU USAGE\n") +
				metricValStyle.Render("14.2% ") +
				subStyle.Render("(4 cores active)"),
		)

		memCard := cardStyle.Render(
			metricLabelStyle.Render("MEMORY USAGE\n") +
				metricValStyle.Render("2.35 GiB ") +
				subStyle.Render("/ 16.00 GiB (14.6%)"),
		)

		containerCard := cardStyle.Render(
			metricLabelStyle.Render("CONTAINER OVERVIEW\n") +
				metricValStyle.Render("3 Running ") +
				subStyle.Render("• 2 Stopped • 5 Total"),
		)

		diskCard := cardStyle.Render(
			metricLabelStyle.Render("DOCKER DISK USAGE\n") +
				metricValStyle.Render("2.09 GB ") +
				subStyle.Render("(Images: 2.05GB, Containers: 42MB)"),
		)

		row1 := lipgloss.JoinHorizontal(lipgloss.Top, cpuCard, memCard)
		row2 := lipgloss.JoinHorizontal(lipgloss.Top, containerCard, diskCard)

		body := lipgloss.JoinVertical(lipgloss.Left, row1, "\n", row2) + "\n"
		footer := subStyle.Render("\n(Press Esc/b to return to main menu • q to quit)\n")
		return header + body + footer

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
