package creategame

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/tui/constants"
)

type formStyles struct {
	form lipgloss.Style
}

type loadingStyles struct {
	spinner lipgloss.Style
	text    lipgloss.Style
}
type resultsStyles struct {
	successMsg lipgloss.Style
	id         lipgloss.Style
	help       lipgloss.Style
}

type styles struct {
	title         lipgloss.Style
	formStyles    formStyles
	loadingStyles loadingStyles
	resultsStyles resultsStyles
}

func newFormStyles() formStyles {
	return formStyles{
		form: lipgloss.NewStyle().PaddingLeft(2).PaddingBottom(1),
	}
}

func newLoadingStyles() loadingStyles {
	return loadingStyles{
		spinner: lipgloss.NewStyle().PaddingLeft(2).Foreground(constants.TextColor),
		text:    lipgloss.NewStyle().Foreground(constants.TextColor).Italic(true),
	}
}

func newResultsStyles() resultsStyles {
	return resultsStyles{
		successMsg: lipgloss.NewStyle().PaddingLeft(2).Foreground(constants.BlueColor),
		id:         lipgloss.NewStyle().Foreground(constants.BlueColor),
		help:       lipgloss.NewStyle().PaddingLeft(2).PaddingTop(1),
	}
}

func newStyles() styles {
	return styles{
		title: lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingBottom(1).
			PaddingTop(1).
			Foreground(constants.GoldColor),
		formStyles:    newFormStyles(),
		loadingStyles: newLoadingStyles(),
		resultsStyles: newResultsStyles(),
	}
}
