package creategame

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Enter key.Binding
}

var keys = keyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter", "esc"),
		key.WithHelp("enter", "menu"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Enter}}
}
