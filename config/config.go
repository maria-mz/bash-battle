package config

type GameDifficulty int

const (
	DiffVaried GameDifficulty = iota
	DiffMedium
	DiffHard
	DiffEasy
)

type GameFileSize int

const (
	FileVaried GameFileSize = iota
	FileMedium
	FileBig
	FileSmall
)

type GameConfig struct {
	Difficulty    GameDifficulty
	FileSize      GameFileSize
	Rounds        int
	RoundDuration int
	MaxPlayers    int
}

type Config struct {
	ServerAddr string
	Username   string
	GameConfig GameConfig
}
