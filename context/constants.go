package context

import "github.com/maria-mz/bash-battle-proto/proto"

type (
	ConnectionStatus int
	GameStatus       int
)

const (
	IsNotConnected ConnectionStatus = iota
	IsConnected
)

const (
	WaitingForPlayers GameStatus = iota
	WaitingToStartGame
)

var difficultyValues = map[proto.Difficulty]string{
	proto.Difficulty_EasyDiff:   "Easy",
	proto.Difficulty_MediumDiff: "Medium",
	proto.Difficulty_HardDiff:   "Hard",
	proto.Difficulty_VariedDiff: "Varied",
}

var fileSizeValues = map[proto.FileSize]string{
	proto.FileSize_SmallSize:  "Small",
	proto.FileSize_MediumSize: "Medium",
	proto.FileSize_LargeSize:  "Large",
	proto.FileSize_VariedSize: "Varied",
}
