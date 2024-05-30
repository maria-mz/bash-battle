package configtable

import (
	"fmt"

	"github.com/maria-mz/bash-battle/config"
	"github.com/maria-mz/bash-battle/tui/constants"
	"github.com/maria-mz/bash-battle/utils"
)

const (
	// This number is the gap between the end of the longest
	// label and the start of the longest value
	absolutePadding = 4
)

type ConfigTable struct {
	serverAddr string
	rounds     string
	duration   string
	difficulty string
	fileSize   string

	view string // cached view
}

func New(conf config.Config) ConfigTable {
	m := ConfigTable{
		serverAddr: conf.ServerAddr,
		rounds:     utils.IntToString(conf.GameConfig.Rounds),
		duration:   utils.IntToString(conf.GameConfig.RoundDuration),
		difficulty: constants.DifficultyTextMap[conf.GameConfig.Difficulty],
		fileSize:   constants.FileSizeTextMap[conf.GameConfig.FileSize],
	}

	m.buildView()

	return m
}

func (m ConfigTable) View() string {
	return m.view
}

func (m *ConfigTable) buildView() {
	lengthOfLongestField := utils.Max(
		len(constants.ConfigLabelServerAddr),
		len(constants.ConfigLabelRounds),
		len(constants.ConfigLabelRoundDuration),
		len(constants.ConfigLabelDifficulty),
		len(constants.ConfigLabelFileSize),
	)

	lengthOfLongestValue := utils.Max(
		len(m.serverAddr),
		len(m.rounds),
		len(m.duration),
		len(m.difficulty),
		len(m.fileSize),
	)

	maxLength := lengthOfLongestField + lengthOfLongestValue

	table := m.renderTable(maxLength)

	m.view = m.putBorder(table)
}

func (m ConfigTable) putBorder(table string) string {
	return utils.BorderedBoxWithTitle(
		styleConfigTable,
		table,
		styleConfigTitle.Render(constants.ConfigTableTitle),
	)
}

func (m ConfigTable) renderTable(maxLength int) string {
	table := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s",
		m.renderLine(maxLength, constants.ConfigLabelServerAddr, m.serverAddr),
		m.renderLine(maxLength, constants.ConfigLabelRounds, m.rounds),
		m.renderLine(maxLength, constants.ConfigLabelRoundDuration, m.duration),
		m.renderLine(maxLength, constants.ConfigLabelDifficulty, m.difficulty),
		m.renderLine(maxLength, constants.ConfigLabelFileSize, m.fileSize),
	)
	return table
}

func (m ConfigTable) renderLine(maxLength int, label string, val string) string {
	padding := maxLength + absolutePadding - (len(label) + len(val))
	styledVal := styleConfigValue.Render(val)

	line := fmt.Sprintf("%s:%*s%s", label, padding, " ", styledVal)
	return line
}
