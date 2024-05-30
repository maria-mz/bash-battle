// Players Model component
package playerlist

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maria-mz/bash-battle/config"
	"github.com/maria-mz/bash-battle/tui/constants"
	"github.com/maria-mz/bash-battle/tui/messages"
	"github.com/maria-mz/bash-battle/utils"
)

const (
	itemHeight       = 1  // Height of the list item
	itemSpacing      = 0  // Horizontal gap between list items
	defaultListWidth = 40 // TODO: think about this value, not really sure how it's working
	listHeight       = 5  // Aligns with Config box, fits 3 players at a time
)

type PlayerList struct {
	maxPlayers int
	numPlayers int
	list       list.Model
}

func NewPlayerList(conf config.Config) PlayerList {
	return PlayerList{
		maxPlayers: conf.GameConfig.MaxPlayers,
		list:       createEmptyList(),
	}
}

func createEmptyList() list.Model {
	l := list.New(
		make([]list.Item, 0),
		itemDelegate{},
		defaultListWidth,
		listHeight,
	)

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.SetShowTitle(false)

	l.Styles.PaginationStyle = stylePagination

	return l
}

func (m *PlayerList) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case messages.UpdatedPlayerNamesMsg:
		m.numPlayers = len(msg.Names)
		cmd = m.list.SetItems(convertNamesToListItems(msg.Names))

	default:
		m.list, cmd = m.list.Update(msg)
	}

	return cmd
}

func (m PlayerList) View() string {
	title := fmt.Sprintf(
		"%s (%d/%d)", constants.PlayersTableTitle, m.numPlayers, m.maxPlayers,
	)

	return utils.BorderedBoxWithTitle(
		styleList,
		m.list.View(),
		styleTitle.Render(title),
	)
}

func convertNamesToListItems(names []string) []list.Item {
	items := make([]list.Item, 0, len(names))

	for _, name := range names {
		items = append(items, playerListItem(name))
	}

	return items
}
