package btea

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Table table.Model
}

func ModelInit() Model {
	var m Model
	m.UpdateModel() 
	m.Glossy()

	return m
}

func (m Model) Glossy() {

}

func (m Model) UpdateModel() {

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
		}
	}

	var cmd tea.Cmd
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.Table.View() + "\n"
}
