package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/config"
	"github.com/maria-mz/bash-battle/tui/constants"
	"github.com/maria-mz/bash-battle/tui/footer"
	"github.com/maria-mz/bash-battle/tui/lobby"
)

type StateView int

const (
	LobbyView StateView = iota
)

type Tui struct {
	activeView StateView

	lobbyModel  lobby.Lobby
	footerModel footer.Footer

	termWidth  int
	termHeight int

	running bool
}

func NewTui(conf config.Config) *Tui {
	return &Tui{
		activeView:  LobbyView,
		lobbyModel:  lobby.New(conf),
		footerModel: footer.New(conf),
	}
}

func (tui *Tui) Init() tea.Cmd {
	setRunning := func() tea.Cmd {
		tui.running = true
		return nil
	}

	return tea.Batch(
		tea.SetWindowTitle(constants.WindowTitle),
		tui.lobbyModel.Init(),
		setRunning(),
	)
}

func (tui *Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("handling new message %#v", msg)

	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		// Program sends this message on start up, I think it's the only way
		// to know when the program has started
		log.Printf("setting is running to true")
		tui.running = true
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
	tui.footerModel.Update(msg)

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
	mainView := "nothing to show..."

	switch m.activeView {
	case LobbyView:
		mainView = m.lobbyModel.View()
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		mainView,
		m.footerModel.View(m.termWidth),
	)
}

func (m *Tui) IsRunning() bool {
	log.Printf("running = %v", m.running)
	return m.running
}
