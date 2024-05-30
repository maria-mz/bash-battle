// Footer Model component
package footer

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/config"
	"github.com/maria-mz/bash-battle/status"
	"github.com/maria-mz/bash-battle/tui/constants"
	"github.com/maria-mz/bash-battle/tui/messages"
)

type Footer struct {
	username   string
	connStatus status.ConnStatus
	gameStatus status.GameStatus
}

func New(conf config.Config) Footer {
	return Footer{
		username:   conf.Username,
		connStatus: status.Connecting,
		gameStatus: status.Initializing,
	}
}

func (m *Footer) Update(msg tea.Msg) {
	switch msg := msg.(type) {

	case messages.GameStatusMsg:
		m.gameStatus = msg.Status

	case messages.ConnStatusMsg:
		m.connStatus = msg.Status
	}
}

func (m Footer) View(width int) string {
	footerStyle := lipgloss.NewStyle().Width(width)
	var connStyle lipgloss.Style

	if m.connStatus == status.Disconnected {
		connStyle = styleRedText
	} else {
		connStyle = styleGreenText
	}

	footer := footerStyle.Render(
		fmt.Sprintf(
			"%s%s %s",
			styleUsername.Render(m.username),
			connStyle.Render(constants.ConnectionTextMap[m.connStatus]),
			constants.StatusTextMap[m.gameStatus],
		),
	)

	return footer
}
