package context

import "github.com/maria-mz/bash-battle-proto/proto"

type GameConfig struct {
	MaxPlayers    int
	Rounds        int
	RoundDuration int
	Difficulty    string
	FileSize      string
}

type AppContext struct {
	ServerAddress    string
	Username         string
	PlayerNames      []string
	GameConfig       GameConfig
	ConnectionStatus ConnectionStatus
	GameStatus       GameStatus
}

// todo: think about where this logic should go...
func (ctx *AppContext) PopulateGameConfig(config *proto.GameConfig) {
	ctx.GameConfig.MaxPlayers = int(config.MaxPlayers)
	ctx.GameConfig.Rounds = int(config.Rounds)
	ctx.GameConfig.RoundDuration = int(config.RoundSeconds)
	ctx.GameConfig.Difficulty = difficultyValues[config.Difficulty]
	ctx.GameConfig.FileSize = fileSizeValues[config.FileSize]
}

func (ctx *AppContext) PopulatePlayerNames(players []*proto.Player) {
	ctx.PlayerNames = make([]string, 0, len(players))

	for _, player := range players {
		ctx.PlayerNames = append(ctx.PlayerNames, player.Username)
	}
}
