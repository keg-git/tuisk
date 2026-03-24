package data

import (
	"log"
	"os/exec"
	"encoding/json"
	"strings"
)

type Task struct {
	Id 			int			`json:"id"`
	Age	 		string		`json:"age"`
	Tags 		[]string	`json:"tags"`
	Due 		string		`json:"due"`  // this comes out to be a weird number gonna have to figure that one out
	Description	string		`json:"description"`
	Priority	string		`json:"priority"`
	Urgency 	float64		`json:"urgency"`
}

func GetTasks() []Task {

	var tasks []Task
	cmd := exec.Command("task", "status:pending", "export")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(out, &tasks)
	if err != nil {
		log.Fatal(err)
	}

	return tasks
}

func MarkDone(id string) error {
	cmd := exec.Command("task", id, "done")

	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

// we need name due date* priority* tags* description*
func CreateTask(task Task) error {

	var due string
	var prior string

	if task.Due != "" {
		due = "due:" + task.Due
	}

	if task.Priority != "" {
		prior = "priority:" + task.Priority
	}

	tags := "+" + strings.Join(task.Tags, " +")

	cmd := exec.Command("task", "add", due, prior, tags, task.Description)

	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

// gonna need the id along with what needs to be modified
func ModifyTask(id string, category string, value string) error {

	mod := category + ":" + value
	cmd := exec.Command("task", id, "modify", mod)

	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
