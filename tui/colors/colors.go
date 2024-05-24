package colors

import (
	"github.com/charmbracelet/lipgloss"
)

// Palette from https://github.com/JonathanSpeek/palenight-iterm2
var (
	Red     = lipgloss.AdaptiveColor{Light: "#ff8b92", Dark: "#f07178"}
	Yellow  = lipgloss.AdaptiveColor{Light: "#ffe585", Dark: "#ffcb6b"}
	Blue    = lipgloss.AdaptiveColor{Light: "#9cc4ff", Dark: "#82aaff"}
	Green   = lipgloss.AdaptiveColor{Light: "#ddffa7", Dark: "#c3e88d"}
	Cyan    = lipgloss.AdaptiveColor{Light: "#a3f7ff", Dark: "#89ddff"}
	Magenta = lipgloss.AdaptiveColor{Light: "#e1acff", Dark: "#c792ea"}

	Gray = lipgloss.Color("243")
)
