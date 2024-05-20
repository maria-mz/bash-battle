package creategame

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/tui/constants"
)

var (
	titleStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingBottom(1).
			PaddingTop(1).
			Foreground(constants.GoldColor)
	formStyle        = lipgloss.NewStyle().PaddingLeft(2).PaddingBottom(1)
	spinnerStyle     = lipgloss.NewStyle().PaddingLeft(2).Foreground(constants.TextColor)
	loadingTextStyle = spinnerStyle.Italic(true)
)
