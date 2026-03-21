package main

import (
	"fmt"
	"os"
	// "tuisk/tui"
	"tuisk/btea"
	tea "github.com/charmbracelet/bubbletea"

)

func main() {
	m := btea.ModelInit()
	p := tea.NewProgram(&m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
