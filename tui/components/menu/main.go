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

type Choice int

const (
	CreateGameChoice Choice = 0
	JoinGameChoice   Choice = 1
)

var (
	activeStyle = lipgloss.NewStyle().Foreground(constants.GreenColor)
	titleStyle  = lipgloss.NewStyle().Foreground(constants.BlueColor)
)

type Model struct {
	// elements
	title   string
	choices []string
	Choice  Choice
	keys    keyMap
	help    help.Model

	// sizing
	width  int
	height int
}

func NewModel() Model {
	title := `

    Y88b         888                        888           888               888    888    888          
	 Y88b        888                        888           888               888    888    888          
	  Y88b       888                        888           888               888    888    888          
	   Y88b      88888b.   8888b.  .d8888b  88888b.       88888b.   8888b.  888888 888888 888  .d88b.  
	   d88P      888 "88b     "88b 88K      888 "88b      888 "88b     "88b 888    888    888 d8P  Y8b 
	  d88P       888  888 .d888888 "Y8888b. 888  888      888  888 .d888888 888    888    888 88888888 
	 d88P        888 d88P 888  888      X88 888  888      888 d88P 888  888 Y88b.  Y88b.  888 Y8b.     
    d88P         88888P"  "Y888888  88888P' 888  888      88888P"  "Y888888  "Y888  "Y888 888  "Y8888  

	`

	return Model{
		title:   title,
		choices: []string{"Create Game", "Join Game"},
		help:    help.New(),
		keys:    keys,
	}
}

func (m *Model) moveCursorUp() {
	if m.Choice > 0 {
		m.Choice--
	}
}

func (m *Model) moveCursorDown() {
	if int(m.Choice) < len(m.choices)-1 {
		m.Choice++
	}
}

func (m Model) Init() tea.Cmd {
	return nil // Don't need to do any I/O on start
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
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
		}
	}

	return m, nil
}

func (m Model) GetActiveChoice() Choice {
	return m.Choice
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
		if i == int(m.Choice) {
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
