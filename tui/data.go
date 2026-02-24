package tui

import (
	"encoding/json"
	// "fmt"
	"log"
	// "os"
	"os/exec"
)

type task struct {
	id 			int			`json:"id"`
	age	 		string		`json:"age"`
	tags 		[]string	`json:"tags"`
	due 		string		`json:"due"`  // this comes out to be a weird number gonna have to figure that one out
	description	string		`json:"description"`
	priority	string		`json:"priority"`
	urgency 	float32		`json:"urgency"`
}

func (m Model) GetTasks() {
	cmd := exec.Command("task", "status:pending", "export")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(out, m.pending)

	columns := []table.Column{
		{Title: "ID", Width: 3},
		{Title: "Description", Width: 30},
		{Title: "Pri", Width: 3},
		{Title: "Urg", Width: 3},
		{Title: "tags", Width: 50},
	}

	for task, i := range m.pending {
		
	}
}

func (m Model) MarkDone(id string) {

}

// we need name due date* priority* tags* description*
func (m Model) CreateTask() {

}

// gonna need the id along with what needs to be modified
func (m Model) ModifyTask() {

}
