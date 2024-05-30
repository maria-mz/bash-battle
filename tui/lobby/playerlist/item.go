package playerlist

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type playerListItem string

func (p playerListItem) FilterValue() string {
	return ""
}

type itemDelegate struct{}

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
	i, ok := listItem.(playerListItem)
	if !ok {
		return
	}

	line := fmt.Sprintf("%d. %s", index+1, i)

	if index == m.Index() {
		fmt.Fprint(w, styleSelectedItem.Render(line))
	} else {
		fmt.Fprint(w, line)
	}
}
