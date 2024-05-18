package creategame

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	be "github.com/maria-mz/bash-battle/backend"
	"github.com/maria-mz/bash-battle/tui/commands"
	"github.com/maria-mz/bash-battle/utils"
)

type State int

const (
	ShowingForm State = iota
	PendingRequest
	ShowingCode
)

type Model struct {
	// elements
	title   string
	form    *huh.Form
	spinner spinner.Model

	// state
	State State

	// server-side
	backend *be.Backend

	// sizing
	width  int
	height int
}

func newSpinner() spinner.Model {
	spin := spinner.New()
	spin.Spinner = spinner.Points
	spin.Style = spinnerStyle
	return spin
}

func NewModel(backend *be.Backend) Model {
	return Model{
		title:   "▒▒▒▒ Create Game",
		form:    newForm(),
		spinner: newSpinner(),
		backend: backend,
	}
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)

	case commands.CreateGameMsg:
	}

	if m.State == ShowingForm {
		m, cmd = m.updateForm(msg)
		cmds = append(cmds, cmd)
	}

	if m.shouldCreateGame() {
		m.State = PendingRequest
		m.form = newForm()
		cmds = append(cmds, m.spinner.Tick, m.form.Init(), m.getRequestCmd())
	}

	return m, tea.Batch(cmds...)
}

func (m Model) updateForm(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	form, cmd := m.form.Update(msg)

	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) Reset() (Model, tea.Cmd) {
	m.form = newForm()
	m.State = ShowingForm
	return m, m.form.Init()
}

func (m Model) getRequestCmd() func() tea.Msg {
	rounds, _ := utils.StringToInt32(getRounds(m.form))
	minutes, _ := utils.StringToInt32(getRoundMinutes(m.form))

	return commands.CreateGameCmd(m.backend, rounds, minutes*60)
}

func (m Model) shouldCreateGame() bool {
	return m.State == ShowingForm &&
		m.form.State == huh.StateCompleted &&
		wantsToCreateGame(m.form)
}

func (m Model) View() string {
	if m.width == 0 {
		return ""
	}

	switch m.State {
	case ShowingForm:
		return m.getFormView()
	case PendingRequest:
		return m.getLoadingView()
	}

	return "nothing to show..."
}

func (m Model) getFormView() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(m.title),
		formStyle.Render(m.form.View()),
	)
}

func (m Model) getLoadingView() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(m.title),
		loadingStyle.Render(
			fmt.Sprint(m.spinner.View(), " Creating game, please wait..."),
		),
	)
}
