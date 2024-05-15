package creategameform

import (
	"errors"
	"unicode"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/tui/constants"
)

var (
	titleStyle = lipgloss.NewStyle().PaddingLeft(2).PaddingBottom(1).PaddingTop(1).Foreground(constants.Color_MintMain)
	formStyle  = lipgloss.NewStyle().PaddingLeft(2).PaddingBottom(1)
	helpStyle  = lipgloss.NewStyle().PaddingLeft(2)
)

type Model struct {
	title  string
	form   *huh.Form
	keys   keyMap
	help   help.Model
	width  int
	height int
}

func isInputEmpty(s string) bool {
	return len(s) == 0
}

func isInputNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func getFormTheme() *huh.Theme {
	t := huh.ThemeBase()

	t.Focused.Title = t.Focused.Title.Foreground(constants.Color_BlueSecondary)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(constants.Color_RedWarn)
	t.Focused.Description = t.Focused.Description.Foreground(constants.Color_Gray)

	t.Blurred.ErrorMessage = t.Focused.ErrorMessage
	t.Blurred.Description = t.Focused.Description

	return t
}

func initForm() *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("rounds").
				Title("Number of rounds").
				Placeholder("Enter a number").
				Validate(func(s string) error {
					if isInputEmpty(s) {
						return errors.New("field cannot be empty")
					}
					if !isInputNumeric(s) {
						return errors.New("not a valid number")
					}
					return nil
				}).
				CharLimit(2),
			huh.NewInput().
				Key("roundDuration").
				Title("Round duration (minutes)").
				Placeholder("Enter a number").
				Validate(func(s string) error {
					if isInputEmpty(s) {
						return errors.New("field cannot be empty")
					}
					if !isInputNumeric(s) {
						return errors.New("not a valid number")
					}
					return nil
				}).
				CharLimit(2),
			huh.NewConfirm().
				Key("done").
				Title("Create game?").
				Description(
					"Entering 'Yes' should create a new game on the server. \n"+
						"You should shortly receive a code that others can use "+
						"to join the game.",
				).
				Affirmative("Yes").
				Negative("No"),
		).
			WithShowHelp(false),
	)

	form.WithTheme(getFormTheme())

	return form
}

func NewModel() *Model {
	m := &Model{}

	m.title = "▒▒▒▒ Create Game"

	m.form = initForm()
	m.help = help.New()
	m.keys = keys

	return m
}

func (m *Model) Init() tea.Cmd {
	return m.form.Init()
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
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			// TODO: return to menu case
		}
		// Next and Back cases are handled by the form
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	if m.width == 0 {
		// Dimensions haven't been set yet, no view -> empty string
		return ""
	}

	switch m.form.State {
	case huh.StateCompleted:
		// TODO: do things
		return "form complete"
	default:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(m.title),
			formStyle.Render(m.form.View()),
			helpStyle.Render(m.help.View(m.keys)),
		)
	}
}
