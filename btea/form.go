package btea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// tea "github.com/charmbracelet/bubbletea"
)

type form struct {
	form []textinput.Model
	index int
	show bool
	width int
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


func (f form) Update(msg tea.Msg) (form, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			f.show = false
			fmt.Print("escape press")
			return f, cmd
		case "ctrl+c":
			return f, tea.Quit
		}
	case tea.WindowSizeMsg:
		f.width = msg.Width
	}

	return f, cmd
}

func (f *form) initAddForm() {
	f.index = 0
	f.show = false

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
