package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maria-mz/bash-battle-proto/proto"
	be "github.com/maria-mz/bash-battle/backend"
)

type ErrMsg struct {
	Err error
}

type CreateGameMsg struct {
	ErrMsg
	GameCode string
	GameID   string
}

func CreateGameCmd(backend *be.Backend, rounds int32, minutes int32) func() tea.Msg {
	return func() tea.Msg {
		request := &proto.CreateGameRequest{
			GameConfig: &proto.GameConfig{
				Rounds:       rounds,
				RoundSeconds: minutes,
			},
		}
		response, err := backend.CreateGame(request)

		if err != nil {
			return CreateGameMsg{ErrMsg: ErrMsg{Err: err}}
		}

		return CreateGameMsg{
			GameCode: response.GameCode,
			GameID:   response.GameID,
		}
	}
}
