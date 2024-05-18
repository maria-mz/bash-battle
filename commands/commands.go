package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maria-mz/bash-battle-proto/proto"
	be "github.com/maria-mz/bash-battle/backend"
)

type CmdBuilder struct {
	backend *be.Backend
}

func NewCmdBuilder(backend *be.Backend) *CmdBuilder {
	return &CmdBuilder{backend: backend}
}

func (b *CmdBuilder) NewLoginCmd(name string) func() tea.Msg {
	return func() tea.Msg {
		request := &proto.LoginRequest{
			PlayerName: name,
		}

		resp, err := b.backend.Login(request)

		if err != nil {
			return CriticalErrorMsg{Error: err}
		}

		var errCode ErrorCode

		switch resp.ErrorCode {
		case proto.LoginResponse_NAME_TAKEN_ERR:
			errCode = NameTaken
		}

		return &LoginMsg{ErrorCode: &errCode}
	}
}

func (b *CmdBuilder) NewCreateGameCmd(rounds int32, seconds int32) func() tea.Msg {
	return func() tea.Msg {
		request := &proto.CreateGameRequest{
			GameConfig: &proto.GameConfig{
				Rounds:       rounds,
				RoundSeconds: seconds,
			},
		}

		resp, err := b.backend.CreateGame(request)

		if err != nil {
			return CriticalErrorMsg{Error: err}
		}

		return CreateGameMsg{
			GameCode: resp.GameCode,
			GameID:   resp.GameID,
		}
	}
}

func (b *CmdBuilder) NewJoinGameCmd(id string, code string) func() tea.Msg {
	return func() tea.Msg {
		request := &proto.JoinGameRequest{
			GameID:   id,
			GameCode: code,
		}

		resp, err := b.backend.JoinGame(request)

		if err != nil {
			return CriticalErrorMsg{Error: err}
		}

		var errCode ErrorCode

		switch resp.ErrorCode {
		case proto.JoinGameResponse_GAME_NOT_FOUND_ERR:
			errCode = GameNotFound
		case proto.JoinGameResponse_INVALID_CODE_ERR:
			errCode = InvalidCode
		case proto.JoinGameResponse_GAME_LOBBY_CLOSED_ERR:
			errCode = GameLobbyClosed
		}

		return &JoinGameMsg{ErrorCode: &errCode}
	}
}
