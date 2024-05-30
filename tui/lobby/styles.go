package lobby

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/tui/colors"
)

var (
	styleBashCube          = lipgloss.NewStyle().Foreground(colors.Blue).PaddingBottom(2).PaddingRight(2)
	styleBashBattleTitle   = lipgloss.NewStyle().Foreground(colors.Blue)
	styleBashBattleWelcome = lipgloss.NewStyle().Foreground(colors.Blue).PaddingTop(1).PaddingBottom(2)
)
