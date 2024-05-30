// Main Lobby Model!!!

package lobby

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/maria-mz/bash-battle/config"
	"github.com/maria-mz/bash-battle/tui/constants"

	"github.com/maria-mz/bash-battle/tui/lobby/configtable"
	"github.com/maria-mz/bash-battle/tui/lobby/help"
	"github.com/maria-mz/bash-battle/tui/lobby/playerlist"
)

type Lobby struct {
	playerList  playerlist.PlayerList
	configTable configtable.ConfigTable
	help        help.Help

	width  int
	height int
}

func New(conf config.Config) Lobby {
	return Lobby{
		configTable: configtable.NewConfigTable(conf),
		playerList:  playerlist.NewPlayerList(conf),
		help:        help.NewHelp(),
	}
}

func (m Lobby) Init() tea.Cmd {
	return nil
}

func (m Lobby) Update(msg tea.Msg) (Lobby, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.help.Keys.Quit):
			return m, tea.Quit
		}
	}

	cmd := m.updateComponents(msg)

	return m, cmd
}

func (m *Lobby) updateComponents(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.playerList.Update(msg))
	return tea.Batch(cmds...)
}

func (m Lobby) View() string {
	if m.width == 0 {
		return ""
	}

	return m.mainView()
}

func (m Lobby) mainView() string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.titleView(),
			m.tablesView(),
			m.help.View(),
		),
	)
}

func (m Lobby) titleView() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			styleBashCube.Render(constants.BashCubeASCII),
			lipgloss.JoinVertical(
				lipgloss.Center,
				styleBashBattleTitle.Render(constants.BashBattleTitle),
				styleBashBattleWelcome.Render(constants.BashBattleWelcome),
			),
		),
	)
}

func (m Lobby) tablesView() string {
	configTable := m.configTable.View()
	playersTable := m.playerList.View()

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		lipgloss.Place(
			lipgloss.Width(configTable),
			lipgloss.Height(playersTable),
			lipgloss.Left,
			lipgloss.Top,
			lipgloss.NewStyle().PaddingRight(1).Render(configTable),
		),
		lipgloss.Place(
			lipgloss.Width(playersTable),
			lipgloss.Height(playersTable),
			lipgloss.Left,
			lipgloss.Top,
			playersTable,
		),
	)
}
