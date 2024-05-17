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
			m.clearForm()
			m.menuCallback()
			m.done = false // reset
			return m, m.form.Init()
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

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

	switch m.form.State {
	case huh.StateCompleted:
		if wantsToCreateGame(m.form) {
			return m.getLoadingView()
		}
		return ""
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
}

func (m *Model) handleFormDone(cmds *[]tea.Cmd) {
	if wantsToCreateGame(m.form) {
		m.doneCallback(getRounds(m.form), getRoundMinutes(m.form))
		m.done = true
		*cmds = append(*cmds, m.spinner.Tick)
		return
	}

	m.clearForm()
	m.menuCallback()
	m.done = false // reset
	*cmds = append(*cmds, m.form.Init())
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
