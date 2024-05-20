package creategame

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
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
	OnForm State = iota
	OnLoading
	OnResults
)

type Model struct {
	// elements
	title   string
	form    *huh.Form // form has its own keys & help
	spinner spinner.Model
	keys    keyMap     // keys for loading & results view
	help    help.Model // help for loading & results view

	// styles
	styles styles

	// state
	State  State
	gameID string

	cmdBuilder *commands.CmdBuilder

	// callbacks
	doneCallback func()

	width  int
	height int
}

func newSpinner(style lipgloss.Style) spinner.Model {
	spin := spinner.New()
	spin.Spinner = spinner.Line
	spin.Style = style
	return spin
}

func NewModel(cmdBuilder *commands.CmdBuilder, done func()) Model {
	styles := newStyles()

	return Model{
		title:        "▒▒▒▒ Create Game",
		form:         newForm(),
		spinner:      newSpinner(styles.loadingStyles.spinner),
		keys:         keys,
		help:         help.New(),
		styles:       styles,
		cmdBuilder:   cmdBuilder,
		doneCallback: done,
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
			m.done()
			return m, m.form.Init()
		case "enter":
			if m.State == OnResults {
				m.done()
				return m, m.form.Init()
			}
			// if OnForm, form will handle enter event (e.g. move to next field)
		}

	case spinner.TickMsg:
		if m.State == OnLoading {
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}

	case *pb.CreateGameResponse:
		m.State = OnResults
		m.gameID = msg.GameID
	}

	if m.State == OnForm {
		m, cmd = m.updateForm(msg)
		cmds = append(cmds, cmd)
	}

	if m.State == OnForm && m.form.State == huh.StateCompleted {
		if wantsToCreateGame(m.form) {
			m.State = OnLoading
			cmds = append(cmds, m.spinner.Tick, m.getCreateGameCmd())
		} else {
			m.done()
			cmds = append(cmds, m.form.Init())
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

func (m *Model) done() {
	m.form = newForm()
	m.State = OnForm
	m.doneCallback()
}

func (m Model) View() string {
	if m.width == 0 {
		return ""
	}

	switch m.State {
	case OnForm:
		return m.formView()
	case OnLoading:
		return m.loadingView()
	case OnResults:
		return m.resultsView()
	}

	return "nothing to show..."
}

func (m Model) formView() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.styles.title.Render(m.title),
		m.styles.formStyles.form.Render(m.form.View()),
	)
}

func (m Model) loadingView() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.styles.title.Render(m.title),
		fmt.Sprint(
			m.spinner.View(),
			m.styles.loadingStyles.text.Render(" Creating game, please hold...")),
	)
}

func (m Model) resultsView() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.styles.title.Render(m.title),
		fmt.Sprintf(
			"%s Use the following ID to join the game: %s.",
			m.styles.resultsStyles.successMsg.Render("✓ Game Created!"),
			m.styles.resultsStyles.id.Render(m.gameID),
		),
		m.styles.resultsStyles.help.Render(m.help.View(m.keys)),
	)
}
