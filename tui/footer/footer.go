// Footer Model component
package footer

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/config"
	"github.com/maria-mz/bash-battle/messages"
	"github.com/maria-mz/bash-battle/status"
	"github.com/maria-mz/bash-battle/tui/constants"
)

type Footer struct {
	username   string
	connStatus status.ConnStatus
	gameStatus status.GameStatus

	spinner     spinner.Model
	showSpinner bool
}

func newSpinner() spinner.Model {
	spin := spinner.New()
	spin.Spinner = spinner.MiniDot
	spin.Style = styleSpinner
	return spin
}

func New(conf config.Config) Footer {
	return Footer{
		username:   conf.Username,
		connStatus: status.Connecting,
		gameStatus: status.Initializing,
		spinner:    newSpinner(),
	}
}

func (m *Footer) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case messages.GameStatusMsg:
		m.gameStatus = msg.Status

		if m.gameStatus == status.WaitingForPlayers {
			m.showSpinner = true
			cmd = m.spinner.Tick
		} else {
			m.showSpinner = false
		}

	case messages.ConnStatusMsg:
		m.connStatus = msg.Status

	case spinner.TickMsg:
		if m.showSpinner {
			m.spinner, cmd = m.spinner.Update(msg)
		}
	}

	return cmd
}

func (m Footer) View(width int) string {
	footerStyle := lipgloss.NewStyle().Width(width)

	footer := footerStyle.Render(
		fmt.Sprintf(
			"%s%s %s",
			styleUsername.Render(m.username),
			m.connStatusView(),
			m.gameStatusView(),
		),
	)

	return footer
}

func (m Footer) connStatusView() string {
	var connStyle lipgloss.Style

	if m.connStatus == status.Disconnected {
		connStyle = styleRedText
	} else {
		connStyle = styleGreenText
	}

	return connStyle.Render(constants.ConnStatusTextMap[m.connStatus])
}

func (m Footer) gameStatusView() string {
	if m.showSpinner {
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.spinner.View(),
			constants.GameStatusTextMap[m.gameStatus],
		)
	}
	return constants.GameStatusTextMap[m.gameStatus]
}
