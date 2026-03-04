package tui

import (
	// "fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Table   table.Model
	pending []task
	rows    []table.Row
}

func InitialModel() Model {
	var m Model
	m.GetTasks()

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	m.Table.SetStyles(s)

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

			// case "up", "k":
			// 	if m.table.Cursor() > 0 {
			// 		m.table.SetCursor(m.table.Cursor()-1)
			// 	}
			//
			// case "down", "j":
			// 	if m.table.Cursor() < len(m.rows)-1 {
			// 		m.table.SetCursor(m.table.Cursor()+1)
			// 	}
		}
	}

	var cmd tea.Cmd
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.Table.View() + "\n"
}
