package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maria-mz/bash-battle/tui/widgets/creategameform"
	"github.com/maria-mz/bash-battle/tui/widgets/menu"
)

type Widget int

const (
	MENU             Widget = 0
	CREATE_GAME_FORM Widget = 1
)

type Tui struct {
	currentWidget  *Widget
	menu           tea.Model
	createGameForm tea.Model
}

func NewTui() *Tui {
	var currentWidget = new(Widget)

	m := &Tui{}

	m.menu = menu.NewModel(m.onMenuCreateGame, m.onMenuJoinGame)
	m.createGameForm = creategameform.NewModel()
	m.currentWidget = currentWidget

	return m
}

func (m *Tui) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tea.SetWindowTitle("Bash Battle"),
		m.menu.Init(),
		m.createGameForm.Init(),
	)
}

func (m *Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// This update should go to all widgets, difficult to send
		// this one after switch, so update as we go.
		cmds = append(cmds, m.updateWidget(&m.menu, msg))
		cmds = append(cmds, m.updateWidget(&m.createGameForm, msg))
	default:
		// Other updates should just target the current widget.
		switch *m.currentWidget {
		case MENU:
			cmds = append(cmds, m.updateWidget(&m.menu, msg))
		case CREATE_GAME_FORM:
			cmds = append(cmds, m.updateWidget(&m.createGameForm, msg))
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *Tui) updateWidget(widget *tea.Model, msg tea.Msg) tea.Cmd {
	model, cmd := (*widget).Update(msg)
	*widget = model
	return cmd
}

func (m *Tui) View() string {
	switch *m.currentWidget {
	case MENU:
		return m.menu.View()
	case CREATE_GAME_FORM:
		return m.createGameForm.View()
	default:
		// This case shouldn't happen
		return "nothing to show..."
	}
}

func (m *Tui) onMenuCreateGame() {
	*m.currentWidget = CREATE_GAME_FORM
}

func (m *Tui) onMenuJoinGame() {

}
