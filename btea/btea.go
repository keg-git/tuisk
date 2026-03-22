package btea

import (
	// "fmt"
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

func (m *Model) GlossTable() {

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
	var cmdTable tea.Cmd
	m.Table, cmdTable = m.Table.Update(msg)

	var cmdForm1 tea.Cmd
	m.addForm, cmdForm1 = m.addForm.Update(msg)

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
		m.height = msg.Height

		m.Table.SetWidth(msg.Width - 4)
		m.Table.SetHeight(msg.Height - 6)
	}

	return m, tea.Batch(cmdTable, cmdForm1)
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
