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

const (
	CreateGameChoice string = "Create Game"
	JoinGameChoice   string = "Join Game"
)

var (
	titleStyle  = lipgloss.NewStyle().Foreground(constants.BlueColor)
	activeStyle = lipgloss.NewStyle().Foreground(constants.GoldColor)
)

type Model struct {
	// elements
	title   string
	choices []string
	cursor  int
	keys    keyMap
	help    help.Model

	// callbacks
	callbacks *MenuChoiceCallbacks

	// sizing
	width  int
	height int
}

type MenuChoiceCallbacks struct {
	CreateGameChoice func()
	JoinGameChoice   func()
}

func NewModel(callbacks *MenuChoiceCallbacks) Model {
	title := `
    __                __               __       __          __  __  __   
    \ \              / /_  ____ ______/ /_     / /_  ____ _/ /_/ /_/ /__ 
     \ \            / __ \/ __  / ___/ __ \   / __ \/ __  / __/ __/ / _ \
     / /           / /_/ / /_/ (__  ) / / /  / /_/ / /_/ / /_/ /_/ /  __/
    /_/  ______   /_.___/\__,_/____/_/ /_/  /_.___/\__,_/\__/\__/_/\___/ 
        /_____/                                                          
    `

	return Model{
		title:     title,
		choices:   []string{CreateGameChoice, JoinGameChoice},
		help:      help.New(),
		keys:      keys,
		callbacks: callbacks,
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

func (m Model) Init() tea.Cmd {
	return nil // Don't need to do any I/O on start
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	currentOption := m.choices[m.cursor]

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
			if currentOption == CreateGameChoice {
				m.callbacks.CreateGameChoice()
			} else if currentOption == JoinGameChoice {
				m.callbacks.JoinGameChoice()
			}
		}
	}

	return m, nil
}

func formatActiveChoice(choice string) string {
	withArrows := fmt.Sprintf(">> %s <<", choice)
	return fmt.Sprintf("%s\n", activeStyle.Render(withArrows))
}

func formatInactiveChoice(choice string) string {
	return fmt.Sprintf("%s\n", choice)
}

func (m Model) View() string {
	if m.width == 0 {
		// Dimensions haven't been set yet, no view -> empty string
		return ""
	}

	var s strings.Builder

	title := titleStyle.Render(m.title)

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
