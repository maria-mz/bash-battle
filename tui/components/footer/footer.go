// Footer Model component
package footer

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/context"
	"github.com/maria-mz/bash-battle/tui/colors"
)

var gameStatusText = map[context.GameStatus]string{
	context.WaitingForPlayers: "Waiting for players...",
}

var connectionStatusText = map[context.ConnectionStatus]string{
	context.IsConnected:    "✓ Connected",
	context.IsNotConnected: "Disconnected",
}

type Status int

const (
	WaitingForPlayers Status = iota
	WaitingForServer
	PlayerLoggedIn
	PlayerLoggedOut
)

var (
	usernameStyle       = lipgloss.NewStyle().Foreground(colors.Cyan).Bold(true).Padding(0, 1)
	isConnectedStyle    = lipgloss.NewStyle().Foreground(colors.Green).Bold(true).Padding(0, 1)
	isNotConnectedStyle = lipgloss.NewStyle().Foreground(colors.Red).Bold(true).Padding(0, 1)
)

type FooterModel struct {
	Username         string
	ConnectionStatus context.ConnectionStatus
	GameStatus       context.GameStatus
}

func (m FooterModel) View(width int) string {
	connectionText := connectionStatusText[m.ConnectionStatus]
	connectionStyle := m.getConnectionStyle()
	statusText := gameStatusText[m.GameStatus]

	footerStyle := lipgloss.NewStyle().Width(width)

	footer := footerStyle.Render(
		fmt.Sprintf(
			"%s%s %s",
			usernameStyle.Render(m.Username),
			connectionStyle.Render(connectionText),
			statusText,
		),
	)

	return footer
}

func (m FooterModel) getConnectionStyle() lipgloss.Style {
	if m.ConnectionStatus == context.IsConnected {
		return isConnectedStyle
	} else {
		return isNotConnectedStyle
	}
}
