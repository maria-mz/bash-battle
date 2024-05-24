// Help Model component
package lobby

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

var (
	helpStyle = lipgloss.NewStyle().Padding(1, 0)
)

type keyMap struct {
	Quit   key.Binding
	Up     key.Binding
	Down   key.Binding
	browse key.Binding
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "leave/quit"),
	),
	Up: key.NewBinding(
		key.WithKeys("↑"),
	),
	Down: key.NewBinding(
		key.WithKeys("↓"),
	),
	browse: key.NewBinding(
		key.WithKeys("↓", "↑"),
		key.WithHelp("↑/↓", "browse players"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.browse}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Quit, k.browse}}
}

type helpModel struct {
	Keys keyMap
	help help.Model
}

func newHelpModel() helpModel {
	return helpModel{Keys: keys, help: help.New()}
}

func (m helpModel) View() string {
	return helpStyle.Render(m.help.View(m.Keys))
}
