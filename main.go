package main

import (
	"fmt"
	"log"
	"os"

	// "tuisk/tui"
	"tuisk/btea"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	file, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)

	m := btea.ModelInit()
	p := tea.NewProgram(&m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
