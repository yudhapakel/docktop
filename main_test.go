package main

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestInitialModel(t *testing.T) {
	m := initialModel()
	if m.state != viewMenu {
		t.Fatalf("expected state viewMenu, got %v", m.state)
	}
	if len(m.menuItems) != 4 {
		t.Fatalf("expected 4 menu items, got %d", len(m.menuItems))
	}
	if len(m.containerTable.Rows()) != 5 {
		t.Fatalf("expected 5 container rows, got %d", len(m.containerTable.Rows()))
	}
	if len(m.imageTable.Rows()) != 5 {
		t.Fatalf("expected 5 image rows, got %d", len(m.imageTable.Rows()))
	}
}

func TestViewTransitions(t *testing.T) {
	m := initialModel()

	// Navigate to Images (index 1) and press enter
	m.menuCursor = 1
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = updated.(model)

	if m.state != viewImages {
		t.Fatalf("expected state viewImages, got %v", m.state)
	}

	viewOutput := m.View()
	if !strings.Contains(viewOutput, "Docktop • Images") {
		t.Fatalf("expected view to contain 'Docktop • Images'")
	}

	// Press Esc to return to menu
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	m = updated.(model)
	if m.state != viewMenu {
		t.Fatalf("expected state viewMenu after Esc, got %v", m.state)
	}

	// Navigate to System Metrics (index 2) and press enter
	m.menuCursor = 2
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = updated.(model)
	if m.state != viewMetrics {
		t.Fatalf("expected state viewMetrics, got %v", m.state)
	}

	viewOutput = m.View()
	if !strings.Contains(viewOutput, "CPU USAGE") || !strings.Contains(viewOutput, "MEMORY USAGE") {
		t.Fatalf("expected metrics view to contain CPU and MEMORY labels")
	}

	// Press 'b' to return to menu
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'b'}})
	m = updated.(model)
	if m.state != viewMenu {
		t.Fatalf("expected state viewMenu after 'b', got %v", m.state)
	}
}
