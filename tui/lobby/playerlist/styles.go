// Players Model component
package playerlist

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/tui/colors"
)

var (
	stylePagination   = list.DefaultStyles().PaginationStyle.Padding(0)
	styleSelectedItem = lipgloss.NewStyle().Foreground(colors.Cyan)
	styleList         = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 3)
	styleTitle        = lipgloss.NewStyle().Foreground(colors.Yellow)
)
