package creategame

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	pb "github.com/maria-mz/bash-battle-proto/proto"
	"github.com/maria-mz/bash-battle/commands"
	"github.com/maria-mz/bash-battle/utils"
)

type State int

const (
	onForm State = iota
	awaitingResp
	onResults
)

type Model struct {
	// elements
	title   string
	form    *huh.Form
	spinner spinner.Model

	// state
	State State

	// service command builder
	cmdBuilder *commands.CmdBuilder

	// callbacks
	done func()

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

func NewModel(cmdBuilder *commands.CmdBuilder, done func()) Model {
	return Model{
		title:      "▒▒▒▒ Create Game",
		form:       newForm(),
		spinner:    newSpinner(),
		cmdBuilder: cmdBuilder,
		done:       done,
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

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.form = newForm()
			m.State = onForm
			m.done()
			return m, m.form.Init()
		}

	case spinner.TickMsg:
		if m.State == awaitingResp {
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}

	case pb.CreateGameResponse:
		m.State = onResults
		// todo: show results
	}

	if m.State == onForm {
		m, cmd = m.updateForm(msg)
		cmds = append(cmds, cmd)
	}

	if m.State == onForm && m.form.State == huh.StateCompleted {
		if wantsToCreateGame(m.form) {
			m.State = awaitingResp
			cmds = append(cmds, m.spinner.Tick, m.getCreateGameCmd())
		} else {
			m.form = newForm()
			m.State = onForm
			cmds = append(cmds, m.form.Init())
			m.done()
		}
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

func (m Model) getCreateGameCmd() func() tea.Msg {
	rounds, _ := utils.StringToInt32(getRounds(m.form))
	minutes, _ := utils.StringToInt32(getRoundMinutes(m.form))

	request := &pb.CreateGameRequest{
		GameConfig: &pb.GameConfig{
			Rounds:       rounds,
			RoundSeconds: minutes * 60,
		},
	}

	return m.cmdBuilder.NewCreateGameCmd(request)
}

func (m Model) View() string {
	if m.width == 0 {
		return ""
	}

	switch m.State {
	case onForm:
		return m.getFormView()
	case awaitingResp:
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
