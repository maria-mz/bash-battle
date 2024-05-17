package creategameform

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/tui/constants"
)

var (
	titleStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingBottom(1).
			PaddingTop(1).
			Foreground(constants.GreenColor)
	formStyle    = lipgloss.NewStyle().PaddingLeft(2).PaddingBottom(1)
	spinnerStyle = lipgloss.NewStyle().Foreground(constants.PurpleColor)
	helpStyle    = lipgloss.NewStyle().PaddingLeft(2)
	loadingStyle = lipgloss.NewStyle().PaddingLeft(2)
)
