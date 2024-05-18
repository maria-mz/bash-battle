package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	be "github.com/maria-mz/bash-battle/backend"
	"github.com/maria-mz/bash-battle/tui/components/creategame"
	"github.com/maria-mz/bash-battle/tui/components/menu"
)

type StateView int

const (
	MenuView StateView = iota
	CreateGameView
)

type Tui struct {
	activeView     StateView
	skipViewUpdate bool

	menuModel       menu.Model
	createGameModel creategame.Model

	backend *be.Backend
}

func NewTui(backend *be.Backend) *Tui {
	return &Tui{
		activeView:      MenuView,
		menuModel:       menu.NewModel(),
		createGameModel: creategame.NewModel(backend),
		backend:         backend,
	}
}

func (m Tui) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Bash Battle"),
		m.menuModel.Init(),
		m.createGameModel.Init(),
	)
}

func (m Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("handling new message %+v", msg)
	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.skipViewUpdate = false

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m, cmd = m.handleWindowSizeMsg(msg)
		return m, cmd

	case tea.KeyMsg:
		m, cmd = m.handleKeyMsg(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	if !m.skipViewUpdate {
		m, cmd = m.updateActiveView(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Tui) handleWindowSizeMsg(msg tea.WindowSizeMsg) (Tui, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.menuModel, cmd = m.menuModel.Update(msg)
	cmds = append(cmds, cmd)

	m.createGameModel, cmd = m.createGameModel.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Tui) handleKeyMsg(msg tea.KeyMsg) (Tui, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {

	case "ctrl+c", "q":
		return m, tea.Quit

	case "enter":
		switch m.activeView {
		case MenuView:
			switch m.menuModel.Choice {
			case menu.CreateGameChoice:
				m.activeView = CreateGameView
				m.skipViewUpdate = true
			}
		}

	case "esc":
		switch m.activeView {
		case CreateGameView:
			m.activeView = MenuView
			m.createGameModel, cmd = m.createGameModel.Reset()
			return m, cmd
		}
	}

	return m, nil
}

func (m Tui) updateActiveView(msg tea.Msg) (Tui, tea.Cmd) {
	var cmd tea.Cmd

	switch m.activeView {
	case MenuView:
		m.menuModel, cmd = m.menuModel.Update(msg)

	case CreateGameView:
		m.createGameModel, cmd = m.createGameModel.Update(msg)
	}

	return m, cmd
}

func (m Tui) View() string {
	switch m.activeView {
	case MenuView:
		return m.menuModel.View()
	case CreateGameView:
		return m.createGameModel.View()
	default:
		return "nothing to show..."
	}
}
