package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maria-mz/bash-battle/tui"
)

func main() {
	tea.LogToFile("debug.log", "debug")

	p := tea.NewProgram(tui.NewTui(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error while running program: %v", err)
		os.Exit(1)
	}
}
