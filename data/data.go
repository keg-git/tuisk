package data

import (
	"encoding/json"
	"log"
	"os/exec"
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

	// var due string
	// var prior string
	// var tags string
	//
	// desc := "\"" + task.Description + "\""
	//
	// if task.Due != "" {
	// 	due = "due:\"" + task.Due + "\""
	// }
	//
	// if task.Priority != "" {
	// 	prior = "priority:" + task.Priority
	// }
	//
	// if len(task.Tags) > 0 {
	// 	tags = "+" + strings.Join(task.Tags, " +")
	// }
	//
	// command := "task add " + desc + " " + tags + " " + prior + " " + due
	// log.Println(command)
	//
	// cmd := exec.Command("task", "add", desc, tags, prior, due)
	//
	// _, err := cmd.Output()
	// if err != nil {
	// 	return err
	// }

	desc := task.Description

	// Build args slice
	args := []string{"add", desc}

	// Add tags
	if len(task.Tags) > 0 {
		for _, tag := range task.Tags {
			args = append(args, tag)
		}
	}

	// Add priority
	if task.Priority != "" {
		args = append(args, "priority:"+task.Priority)
	}

	// Add due date
	if task.Due != "" {
		args = append(args, "due:"+task.Due)
	}

	// Log the full command for debugging
	log.Println("task", strings.Join(args, " "))

	// Execute
	cmd := exec.Command("task", args...)
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
