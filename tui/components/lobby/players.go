// Players Model component
package lobby

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/tui/colors"
	"github.com/maria-mz/bash-battle/utils"
)

const (
	itemHeight       = 1  // Height of the list item
	itemSpacing      = 0  // Horizontal gap between list items
	defaultListWidth = 40 // TODO: think about this value, not really sure how it's working
	listHeight       = 5  // Aligns with Config box, fits 3 players at a time

	playersTitle = "Players"
)

var (
	paginationStyle   = list.DefaultStyles().PaginationStyle.Padding(0)
	selectedItemStyle = lipgloss.NewStyle().Foreground(colors.Cyan)
	listStyle         = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 3)
	listTitleStyle    = lipgloss.NewStyle().Foreground(colors.Yellow)
)

type item string

// -- Implement list.Item interface
func (p item) FilterValue() string {
	return ""
}

type itemDelegate struct{}

// -- Implement ItemDelegate interface
func (d itemDelegate) Height() int {
	return itemHeight
}

func (d itemDelegate) Spacing() int {
	return itemSpacing
}

func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	line := fmt.Sprintf("%d. %s", index+1, i)

	if index == m.Index() {
		fmt.Fprint(w, selectedItemStyle.Render(line))
	} else {
		fmt.Fprint(w, line)
	}
}

type playersModel struct {
	maxPlayers int
	numPlayers int
	list       list.Model
}

func newPlayersModel(playerNames []string, maxPlayers int) playersModel {
	return playersModel{
		maxPlayers: maxPlayers,
		numPlayers: len(playerNames),
		list:       createList(playerNames),
	}
}

func createList(playerNames []string) list.Model {
	l := list.New(
		getListItems(playerNames),
		itemDelegate{},
		defaultListWidth,
		listHeight,
	)

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.SetShowTitle(false)

	l.Styles.PaginationStyle = paginationStyle

	return l
}

func getListItems(playerNames []string) []list.Item {
	items := make([]list.Item, 0, len(playerNames))

	for _, name := range playerNames {
		items = append(items, item(name))
	}

	return items
}

func (m *playersModel) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return cmd
}

func (m *playersModel) UpdatePlayers(playerNames []string) tea.Cmd {
	cmd := m.list.SetItems(getListItems(playerNames))
	m.numPlayers = len(playerNames)
	return cmd
}

func (m playersModel) View() string {
	title := fmt.Sprintf("%s (%d/%d)", playersTitle, m.numPlayers, m.maxPlayers)

	return utils.BorderedBoxWithTitle(
		listStyle,
		m.list.View(),
		listTitleStyle.Render(title),
	)
}
