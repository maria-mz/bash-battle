// Main Lobby Model!!!

package lobby

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/context"
	"github.com/maria-mz/bash-battle/tui/colors"
)

const (
	bashCube = `
                  +++++++         
		       +++*     *+++      
            +++             *++*  
		 *+*                   *++
		 +                    ++++
		 +                *+++++++
		 +             *++++++++++
		 +           +++++++++++++
		 +          ++++++++++++++
		 +          ++++++++++++++
		 +          +++++*++++++++
		 ++*        +++++*++*+++++
		    +++     +++++++++++*
			   ++++ ++++++++
			      +++++++ 
	`
	title = `
        __               __       __          __  __  __   
       / /_  ____ ______/ /_     / /_  ____ _/ /_/ /_/ /__ 
      / __ \/ __  / ___/ __ \   / __ \/ __  / __/ __/ / _ \
     / /_/ / /_/ (__  ) / / /  / /_/ / /_/ / /_/ /_/ /  __/
    /_.___/\__,_/____/_/ /_/  /_.___/\__,_/\__/\__/_/\___/ 
	`
	welcome = "echo \"Welcome to Bash Battle!\""
)

var (
	bashCubeStyle = lipgloss.NewStyle().Foreground(colors.Blue).PaddingBottom(2).PaddingRight(2)
	titleStyle    = lipgloss.NewStyle().Foreground(colors.Blue)
	welcomeStyle  = lipgloss.NewStyle().Foreground(colors.Blue).PaddingTop(1).PaddingBottom(2)
)

type LobbyModel struct {
	playersModel playersModel
	configModel  configModel
	helpModel    helpModel

	width  int
	height int
}

func New(ctx context.AppContext) LobbyModel {
	config := newConfigModel(ctx.ServerAddress, ctx.GameConfig)
	playersModel := newPlayersModel(ctx.PlayerNames, ctx.GameConfig.MaxPlayers)
	helpModel := newHelpModel()

	return LobbyModel{
		configModel:  config,
		playersModel: playersModel,
		helpModel:    helpModel,
	}
}

func (m LobbyModel) Init() tea.Cmd {
	return nil
}

func (m LobbyModel) Update(msg tea.Msg) (LobbyModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.helpModel.Keys.Quit):
			return m, tea.Quit
		}
	}

	cmd = m.playersModel.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m LobbyModel) View() string {
	if m.width == 0 {
		// Dimensions haven't been set yet, no view -> empty string
		return ""
	}

	return m.mainView()
}

func (m LobbyModel) mainView() string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.titleView(),
			m.tablesView(),
			m.helpModel.View(),
		),
	)
}

func (m LobbyModel) titleView() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			bashCubeStyle.Render(bashCube),
			lipgloss.JoinVertical(
				lipgloss.Center,
				titleStyle.Render(title),
				welcomeStyle.Render(welcome),
			),
		),
	)
}

func (m LobbyModel) tablesView() string {
	configTable := m.configModel.View()
	playersTable := m.playersModel.View()

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		lipgloss.Place(
			lipgloss.Width(configTable),
			lipgloss.Height(playersTable),
			lipgloss.Left,
			lipgloss.Top,
			configTable,
		),
		lipgloss.Place(
			lipgloss.Width(playersTable),
			lipgloss.Height(playersTable),
			lipgloss.Left,
			lipgloss.Top,
			playersTable,
		),
	)
}
