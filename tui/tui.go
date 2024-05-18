package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maria-mz/bash-battle/commands"
	"github.com/maria-mz/bash-battle/tui/components/creategame"
	"github.com/maria-mz/bash-battle/tui/components/menu"
)

type StateView int

const (
	MenuView StateView = iota
	CreateGameView
)

type Tui struct {
	activeView StateView

	menuModel       menu.Model
	createGameModel creategame.Model

	cmdBuilder *commands.CmdBuilder
}

func NewTui(cmdBuilder *commands.CmdBuilder) *Tui {
	tui := &Tui{}

	tui.activeView = MenuView

	menuCallbacks := &menu.MenuChoiceCallbacks{
		CreateGameChoice: tui.onChoosesCreateGame,
		JoinGameChoice:   tui.onChoosesJoinGame,
	}

	tui.menuModel = menu.NewModel(menuCallbacks)
	tui.createGameModel = creategame.NewModel(cmdBuilder, tui.onCreateGameViewDone)
	tui.cmdBuilder = cmdBuilder

	return tui
}

func (tui *Tui) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Bash Battle"),
		tui.menuModel.Init(),
		tui.createGameModel.Init(),
	)
}

func (tui *Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("handling new message %+v", msg)

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cmd = tui.updateAllViews(msg)
		return tui, cmd
	}

	cmd = tui.updateActiveView(msg)

	return tui, cmd
}

func (tui *Tui) updateAllViews(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	tui.menuModel, cmd = tui.menuModel.Update(msg)
	cmds = append(cmds, cmd)

	tui.createGameModel, cmd = tui.createGameModel.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (tui *Tui) updateActiveView(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	switch tui.activeView {
	case MenuView:
		tui.menuModel, cmd = tui.menuModel.Update(msg)

	case CreateGameView:
		tui.createGameModel, cmd = tui.createGameModel.Update(msg)
	}

	return cmd
}

func (tui *Tui) onChoosesCreateGame() {
	tui.activeView = CreateGameView
}

func (tui *Tui) onChoosesJoinGame() {
	// todo
}

func (tui *Tui) onCreateGameViewDone() {
	tui.activeView = MenuView
}

func (m *Tui) View() string {
	switch m.activeView {
	case MenuView:
		return m.menuModel.View()
	case CreateGameView:
		return m.createGameModel.View()
	default:
		return "nothing to show..."
	}
}
