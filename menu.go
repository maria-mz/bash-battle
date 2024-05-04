package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// -- Menu Choices
const (
	CREATE_GAME_CHOICE = "Create Game"
	JOIN_GAME_CHOICE   = "Join Game"
	HELP_CHOICE        = "Help"
	QUIT_CHOICE        = "Quit"
)

// -- Menu Styles
var (
	activeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "28", Dark: "155"}) // Greens
	inactiveStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}) // Grays
)

type menuModel struct {
	title   string
	choices []string
	cursor  int

	width  int
	height int
}

func NewMenuModel() menuModel {
	title := `
    __                __               __       __          __  __  __   
    \ \              / /_  ____ ______/ /_     / /_  ____ _/ /_/ /_/ /__ 
     \ \            / __ \/ __  / ___/ __ \   / __ \/ __  / __/ __/ / _ \
     / /           / /_/ / /_/ (__  ) / / /  / /_/ / /_/ / /_/ /_/ /  __/
    /_/  ______   /_.___/\__,_/____/_/ /_/  /_.___/\__,_/\__/\__/_/\___/ 
        /_____/                                                          
    `

	return menuModel{
		title: title,
		choices: []string{
			CREATE_GAME_CHOICE,
			JOIN_GAME_CHOICE,
			HELP_CHOICE,
			QUIT_CHOICE,
		},
	}
}

func (m *menuModel) MoveCursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *menuModel) MoveCursorDown() {
	if m.cursor < len(m.choices)-1 {
		m.cursor++
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil // Don't need to do any I/O on start
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			m.MoveCursorUp()

		case "down":
			m.MoveCursorDown()

			// TODO: Handle enter
		}
	}

	return m, nil
}

func formatActiveChoice(choice string) string {
	withArrows := fmt.Sprintf(">> %s <<", choice)
	return fmt.Sprintf("%s\n", activeStyle.Render(withArrows))
}

func formatInactiveChoice(choice string) string {
	return fmt.Sprintf("%s\n", inactiveStyle.Render(choice))
}

func (m menuModel) View() string {
	if m.width == 0 {
		// Dimensions haven't been set yet, no view -> empty string
		return ""
	}

	var s strings.Builder

	s.WriteString(fmt.Sprintf("%s\n\n", m.title))

	for i, choice := range m.choices {
		if i == m.cursor {
			s.WriteString(formatActiveChoice(choice))
		} else {
			s.WriteString(formatInactiveChoice(choice))
		}
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		s.String(),
	)
}

func main() {
	p := tea.NewProgram(NewMenuModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error while running program: %v", err)
		os.Exit(1)
	}
}
