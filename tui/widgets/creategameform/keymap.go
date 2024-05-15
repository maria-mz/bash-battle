package creategameform

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Back key.Binding
	Next key.Binding
	Menu key.Binding
	Quit key.Binding
	Help key.Binding
}

var keys = keyMap{
	Back: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "next"),
	),
	Next: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("shift+tab", "back"),
	),
	Menu: key.NewBinding(
		key.WithKeys("-"),
		key.WithHelp("shift+m", "menu"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "more"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Next, k.Help}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Back, k.Next, k.Menu}, // first column
		{k.Help, k.Quit},         // second column
	}
}
