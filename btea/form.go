package btea

import (
	"strings"
	"tuisk/data"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type form struct {
	form []textinput.Model
	index int
	show bool
	width int
	formType string
}

func (f *form) View() string {

	if !f.show {
		return ""
	}

	var inputs []string
	for _, input := range f.form {
		inputs = append(inputs, input.View())
	}

	return lipgloss.JoinVertical(lipgloss.Left, inputs...)
}

//result := "+" + strings.ReplaceAll(s, " ", " +")
// result: "+hello +world +foo" this might work 

func (f form) Update(msg tea.Msg) (form, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			f.show = false
			return f, cmd
		case "ctrl+c":
			return f, tea.Quit
		case "tab":
			f.form[f.index].Blur()
			f.index = (f.index+1) % len(f.form)
			f.form[f.index].Focus()
			return f, cmd
		case "shift+tab":
			f.form[f.index].Blur()
			if f.index > 0 {
				f.index--
			} else {
				f.index = len(f.form)-1
			}
			f.form[f.index].Focus()
			return f, cmd
		}
	case tea.WindowSizeMsg:
		f.width = msg.Width
	}

	f.form[f.index], cmd = f.form[f.index].Update(msg)

	return f, cmd
}

//result := "+" + strings.ReplaceAll(s, " ", " +")
// result: "+hello +world +foo"  would of been cool be didn't work the way I wanted it to

func (f *form) submitForm() {
	switch f.formType {
	case "create":
		task := data.Task{
			Description: f.form[0].Value(),
			Priority: f.form[1].Value(),
			Due: f.form[2].Value(),
			Tags: strings.Split(f.form[2].Value(), " "),
		}
		data.CreateTask(task)

	case "modify":

	}

	for i := range f.form {
		f.form[i].SetValue("") 
	}

	f.show = false
}

func (f *form) initAddForm() {
	f.index = 0
	f.show = false
	f.formType = "create"

	desc := textinput.New()
	desc.Placeholder = "Description"
	desc.Focus()
	f.form = append(f.form, desc)

	prior := textinput.New()
	prior.Placeholder = "Priority"
	prior.Blur()
	f.form = append(f.form, prior)

	date := textinput.New()
	date.Placeholder = "Due Date"
	date.Blur()
	f.form = append(f.form, date)

	tags := textinput.New()
	tags.Placeholder = "Tags"
	tags.Blur()
	f.form = append(f.form, tags)

}
