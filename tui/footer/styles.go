package footer

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/tui/colors"
)

var (
	styleUsername  = lipgloss.NewStyle().Foreground(colors.Cyan).Bold(true).Padding(0, 1)
	styleGreenText = lipgloss.NewStyle().Foreground(colors.Green).Bold(true).Padding(0, 1)
	styleRedText   = lipgloss.NewStyle().Foreground(colors.Red).Bold(true).Padding(0, 1)
	styleSpinner   = lipgloss.NewStyle().PaddingRight(1)
)
