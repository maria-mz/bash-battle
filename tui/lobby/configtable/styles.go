package configtable

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/tui/colors"
)

var (
	styleConfigTable = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2)
	styleConfigTitle = lipgloss.NewStyle()
	styleConfigValue = lipgloss.NewStyle().Foreground(colors.Green)
)
