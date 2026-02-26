package tui

import (
	"encoding/json"
	// "fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

type task struct {
	Id 			int			`json:"id"`
	Age	 		string		`json:"age"`
	Tags 		[]string	`json:"tags"`
	Due 		string		`json:"due"`  // this comes out to be a weird number gonna have to figure that one out
	Description	string		`json:"description"`
	Priority	string		`json:"priority"`
	Urgency 	float64		`json:"urgency"`
}

func (m *Model) GetTasks() {
	cmd := exec.Command("task", "status:pending", "export")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(out, &m.pending)
	if err != nil {
		log.Fatal(err)
	}

	columns := []table.Column{
		{Title: "ID", Width: 3},
		{Title: "Description", Width: 30},
		{Title: "Priority", Width: 8},
		{Title: "Urgency", Width: 7},
		{Title: "tags", Width: 50},
	}

	for _, task := range m.pending {
		var tags string
		for _, tag := range task.Tags {
			tags += tag + " "
		}
		m.rows = append(m.rows, table.Row{
			strconv.Itoa(task.Id),
			task.Description,
			task.Priority,
			strconv.FormatFloat(task.Urgency, 'f', 2, 64),
			tags,
		})	
	}

	m.table = table.New(
		table.WithColumns(columns),
		table.WithRows(m.rows),
		table.WithFocused(true),
		table.WithHeight(10),
		)

}

func (m *Model) MarkDone(id string) {
	cmd := exec.Command("task", id, "done")

	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}

// we need name due date* priority* tags* description*
func (m *Model) CreateTask() {

}

// gonna need the id along with what needs to be modified
func (m *Model) ModifyTask() {

}
