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
	doneCallback func(string, string)
	done         bool

	width int
}

func newSpinner() spinner.Model {
	spin := spinner.New()
	spin.Spinner = spinner.Points
	spin.Style = spinnerStyle
	return spin
}

func NewModel(menuCallback func(), doneCallback func(string, string)) *Model {
	return &Model{
		title:        "▒▒▒▒ Create Game",
		form:         newForm(),
		spinner:      newSpinner(),
		keys:         keys,
		help:         help.New(),
		menuCallback: menuCallback,
		doneCallback: doneCallback,
	}
}

func (m *Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil

		case key.Matches(msg, m.keys.Menu):
			m.clearForm()
			m.menuCallback()
			m.done = false // reset
			return m, m.form.Init()
		}

	case spinner.TickMsg:
		s, cmd := m.spinner.Update(msg)
		m.spinner = s
		return m, cmd
	}

	// Not any of those types, should be a message handled by the form
	var cmds []tea.Cmd
	m.updateForm(msg, &cmds)

	if !m.done && m.form.State == huh.StateCompleted {
		m.handleFormDone(&cmds)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	if m.width == 0 {
		// Dimensions haven't been set yet, nothing to show
		return ""
	}

	var view string

	switch m.form.State {
	case huh.StateCompleted:
		if wantsToCreateGame(m.form) {
			view = m.getLoadingView()
		} else {
			view = ""
		}
	default:
		view = m.getFormView()
	}

	return view
}

func (m *Model) updateForm(msg tea.Msg, cmds *[]tea.Cmd) {
	form, cmd := m.form.Update(msg)

	if f, ok := form.(*huh.Form); ok {
		m.form = f
		*cmds = append(*cmds, cmd)
	}
}

func (m *Model) handleFormDone(cmds *[]tea.Cmd) {
	if wantsToCreateGame(m.form) {
		m.doneCallback(getRounds(m.form), getRoundMinutes(m.form))
		m.done = true
		*cmds = append(*cmds, m.spinner.Tick)
	} else {
		m.clearForm()
		m.menuCallback()
		m.done = false // reset
		*cmds = append(*cmds, m.form.Init())
	}
}

func (m *Model) clearForm() {
	m.form = newForm()
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
