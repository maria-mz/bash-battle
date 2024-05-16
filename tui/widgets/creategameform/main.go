package creategameform

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	title string

	form    *huh.Form
	spinner spinner.Model

	keys keyMap
	help help.Model

	menuCallback func()

	width int
}

func newSpinner() spinner.Model {
	spin := spinner.New()
	spin.Spinner = spinner.Points
	spin.Style = spinnerStyle
	return spin
}

func NewModel(menuCallback func()) *Model {
	return &Model{
		title:        "▒▒▒▒ Create Game",
		form:         newForm(),
		spinner:      newSpinner(),
		keys:         keys,
		help:         help.New(),
		menuCallback: menuCallback,
	}
}

func (m *Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Menu):
			// Prepare for returning to menu
			m.form = newForm()
			m.menuCallback()
			return m, m.form.Init()
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	m.updateForm(msg, &cmds)

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	if m.width == 0 {
		// Dimensions haven't been set yet, nothing to show
		return ""
	}

	switch m.form.State {
	case huh.StateCompleted:
		if !m.isYesSelected() {
			// User has selected 'No', nothing to show
			return ""
		}
		return m.getLoadingView()
	default:
		return m.getFormView()
	}
}

func (m *Model) updateForm(msg tea.Msg, cmds *[]tea.Cmd) {
	form, cmd := m.form.Update(msg)

	if f, ok := form.(*huh.Form); ok {
		m.form = f
		*cmds = append(*cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		if !m.isYesSelected() {
			// User has selected 'No', return to menu
			m.form = newForm()
			m.menuCallback()
			*cmds = append(*cmds, m.form.Init())
		} else {
			// Start the ticker!
			*cmds = append(*cmds, m.spinner.Tick)
		}
	}
}

func (m *Model) isYesSelected() bool {
	return m.form.GetBool(CONFIRM_KEY)
}

func (m *Model) getFormView() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(m.title),
		formStyle.Render(m.form.View()),
		helpStyle.Render(m.help.View(m.keys)),
	)
}

func (m *Model) getLoadingView() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(m.title),
		loadingStyle.Render(fmt.Sprint(m.spinner.View(), " Creating game, please wait...")),
	)
}
