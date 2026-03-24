package btea

import (
	"strconv"
	"strings"
	"tuisk/data"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Table table.Model
	addForm form
	width int
	height int
}

func ModelInit() Model {
	var m Model
	m.UpdateModel() 
	m.GlossTable()

	m.addForm.initAddForm()

	return m
}

func (m *Model) UpdateModel() {

	tasks := data.GetTasks()

	columns := []table.Column{
		{Title: "ID", Width: 3},
		{Title: "Description", Width: 30},
		{Title: "Priority", Width: 8},
		{Title: "Urgency", Width: 7},
		{Title: "tags", Width: 50},
	}

	var rows []table.Row
	for _, task := range tasks {
		var tags strings.Builder
		for _, tag := range task.Tags {
			tags.WriteString(tag + " ")
		}
		rows = append(rows, table.Row{
			strconv.Itoa(task.Id),
			task.Description,
			task.Priority,
			strconv.FormatFloat(task.Urgency, 'f', 2, 64),
			tags.String(),
		})
	}


	m.Table = table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	if m.addForm.show {
		m.addForm, cmd = m.addForm.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "a":
			m.addForm.show = true
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		// m.height = msg.Height

		m.Table.SetWidth(msg.Width - 4)
		// m.Table.SetHeight(msg.Height - 6)
	}

	m.Table, cmd = m.Table.Update(msg)

	return m, cmd
}

func (m *Model) View() string {

	stack := lipgloss.JoinVertical(lipgloss.Left, m.Table.View(), m.addForm.View())
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		stack,
	)
}
