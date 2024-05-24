// Config Model component
package lobby

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/maria-mz/bash-battle/context"
	"github.com/maria-mz/bash-battle/tui/colors"
	"github.com/maria-mz/bash-battle/utils"
)

type configField string

const (
	serverAddrField    configField = "Server address"
	roundsField        configField = "Number of rounds"
	roundDurationField configField = "Round duration (seconds)"
	difficultyField    configField = "Difficulty"
	fileSizeField      configField = "File size"

	configTitle     = "Config"
	absolutePadding = 4
)

var (
	configStyle      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2)
	configTitleStyle = lipgloss.NewStyle().Foreground(colors.Yellow)
	configValueStyle = lipgloss.NewStyle().Foreground(colors.Green)
)

type configModel struct {
	values map[configField]string
	view   string
}

func newConfigModel(serverAddr string, config context.GameConfig) configModel {
	values := map[configField]string{
		serverAddrField:    serverAddr,
		roundsField:        utils.IntToString(config.Rounds),
		roundDurationField: utils.IntToString(config.RoundDuration),
		difficultyField:    config.Difficulty,
		fileSizeField:      config.FileSize,
	}

	m := configModel{values: values}
	m.buildView()

	return m
}

func (m configModel) View() string {
	return m.view
}

func (m *configModel) buildView() {
	lengthOfLongestField := utils.Max(
		len(serverAddrField),
		len(roundsField),
		len(roundDurationField),
		len(difficultyField),
		len(fileSizeField),
	)

	lengthOfLongestValue := utils.Max(
		len(m.values[serverAddrField]),
		len(m.values[roundsField]),
		len(m.values[roundDurationField]),
		len(m.values[difficultyField]),
		len(m.values[fileSizeField]),
	)

	maxLength := lengthOfLongestField + lengthOfLongestValue

	table := m.renderTable(maxLength)

	m.view = m.borderTable(table)
}

func (m configModel) borderTable(table string) string {
	return utils.BorderedBoxWithTitle(
		configStyle,
		table,
		configTitleStyle.Render(configTitle),
	)
}

func (m configModel) renderTable(maxLength int) string {
	table := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s",
		m.renderLine(maxLength, serverAddrField),
		m.renderLine(maxLength, roundsField),
		m.renderLine(maxLength, roundDurationField),
		m.renderLine(maxLength, difficultyField),
		m.renderLine(maxLength, fileSizeField),
	)
	return table
}

func (m configModel) renderLine(maxLength int, field configField) string {
	padding := maxLength + absolutePadding - (len(field) + len(m.values[field]))
	value := configValueStyle.Render(m.values[field])

	line := fmt.Sprintf("%s:%*s%s", field, padding, " ", value)
	return line
}
