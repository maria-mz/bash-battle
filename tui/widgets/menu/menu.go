package menu

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/tui/constants"
)

// -- Menu Choices
const (
	CREATE_GAME_CHOICE = "Create Game"
	JOIN_GAME_CHOICE   = "Join Game"
)

var (
	activeStyle = lipgloss.NewStyle().Foreground(constants.AquaColor)
)

type Model struct {
	title   string
	choices []string
	cursor  int

	keys keyMap
	help help.Model

	onCreateGameCallback func()
	onJoinGameCallback   func()

	width  int
	height int
}

func NewModel(createGameCallback func(), joinGameCallback func()) *Model {
	title := `
    __                __               __       __          __  __  __   
    \ \              / /_  ____ ______/ /_     / /_  ____ _/ /_/ /_/ /__ 
     \ \            / __ \/ __  / ___/ __ \   / __ \/ __  / __/ __/ / _ \
     / /           / /_/ / /_/ (__  ) / / /  / /_/ / /_/ / /_/ /_/ /  __/
    /_/  ______   /_.___/\__,_/____/_/ /_/  /_.___/\__,_/\__/\__/_/\___/ 
        /_____/                                                          
    `

	return &Model{
		title: title,
		choices: []string{
			CREATE_GAME_CHOICE,
			JOIN_GAME_CHOICE,
		},
		onCreateGameCallback: createGameCallback,
		onJoinGameCallback:   joinGameCallback,
		help:                 help.New(),
		keys:                 keys,
	}
}

func (m *Model) moveCursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *Model) moveCursorDown() {
	if m.cursor < len(m.choices)-1 {
		m.cursor++
	}
}

func (m *Model) Init() tea.Cmd {
	return nil // Don't need to do any I/O on start
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Up):
			m.moveCursorUp()

		case key.Matches(msg, m.keys.Down):
			m.moveCursorDown()

		case key.Matches(msg, m.keys.Enter):
			// TODO: better way to do this?
			if m.choices[m.cursor] == CREATE_GAME_CHOICE {
				m.onCreateGameCallback()
			} else if m.choices[m.cursor] == JOIN_GAME_CHOICE {
				m.onJoinGameCallback()
			}
		}
	}

	return m, nil
}

func formatActiveChoice(choice string) string {
	withArrows := fmt.Sprintf("> %s <", choice)
	return fmt.Sprintf("%s\n", activeStyle.Render(withArrows))
}

func formatInactiveChoice(choice string) string {
	return fmt.Sprintf("%s\n", choice)
}

func (m *Model) View() string {
	if m.width == 0 {
		// Dimensions haven't been set yet, no view -> empty string
		return ""
	}

	var s strings.Builder

	title := lipgloss.NewStyle().Render(m.title)

	s.WriteString("echo \"Welcome to Bash Battle!\"\n\n")

	for i, choice := range m.choices {
		if i == m.cursor {
			s.WriteString(formatActiveChoice(choice))
		} else {
			s.WriteString(formatInactiveChoice(choice))
		}
	}

	menu := lipgloss.Place(
		lipgloss.Width(title),
		m.height-lipgloss.Width(title),
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			s.String(),
			m.help.View(m.keys),
		),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		menu,
	)
}
