package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/context"
	"github.com/maria-mz/bash-battle/tui/components/footer"
	"github.com/maria-mz/bash-battle/tui/components/lobby"
)

type StateView int

const (
	LobbyView StateView = iota
)

type Tui struct {
	activeView StateView

	lobbyModel  lobby.LobbyModel
	footerModel footer.FooterModel

	termWidth  int
	termHeight int
}

func NewTui(ctx context.AppContext) *Tui {
	tui := &Tui{
		activeView: LobbyView,
		lobbyModel: lobby.New(ctx),
		footerModel: footer.FooterModel{
			Username:         ctx.Username,
			ConnectionStatus: ctx.ConnectionStatus,
			GameStatus:       ctx.GameStatus,
		},
	}

	return tui
}

func (tui *Tui) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Bash Battle"),
		tui.lobbyModel.Init(),
	)
}

func (tui *Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// log.Printf("handling new message %+v", msg)

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		tui.termWidth = msg.Width
		tui.termHeight = msg.Height

		cmd = tui.updateAllViews(msg)
		return tui, cmd
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return tui, tea.Quit
		}
	}

	cmd = tui.updateActiveView(msg)

	return tui, cmd
}

func (tui *Tui) updateAllViews(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	tui.lobbyModel, cmd = tui.lobbyModel.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (tui *Tui) updateActiveView(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	switch tui.activeView {
	case LobbyView:
		tui.lobbyModel, cmd = tui.lobbyModel.Update(msg)
	}
	return cmd
}

func (m *Tui) View() string {
	activeView := "nothing to show..."

	switch m.activeView {
	case LobbyView:
		activeView = m.lobbyModel.View()
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		activeView,
		m.footerModel.View(m.termWidth),
	)
}
