package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	pb "github.com/maria-mz/bash-battle-proto/proto"
	be "github.com/maria-mz/bash-battle/backend"
)

type CriticalErrMsg struct {
	Error error
}

type CmdBuilder struct {
	backend *be.Backend
}

func NewCmdBuilder(backend *be.Backend) *CmdBuilder {
	return &CmdBuilder{backend: backend}
}

func (b *CmdBuilder) NewLoginCmd(name string) func() tea.Msg {
	return func() tea.Msg {
		request := &pb.LoginRequest{
			PlayerName: name,
		}

		resp, err := b.backend.Login(request)

		if err != nil {
			return CriticalErrMsg{Error: err}
		}

		return resp
	}
}

func (b *CmdBuilder) NewCreateGameCmd(rounds int32, seconds int32) func() tea.Msg {
	return func() tea.Msg {
		request := &pb.CreateGameRequest{
			GameConfig: &pb.GameConfig{
				Rounds:       rounds,
				RoundSeconds: seconds,
			},
		}

		resp, err := b.backend.CreateGame(request)

		if err != nil {
			return CriticalErrMsg{Error: err}
		}

		return resp
	}
}

func (b *CmdBuilder) NewJoinGameCmd(id string, code string) func() tea.Msg {
	return func() tea.Msg {
		request := &pb.JoinGameRequest{
			GameID:   id,
			GameCode: code,
		}

		resp, err := b.backend.JoinGame(request)

		if err != nil {
			return CriticalErrMsg{Error: err}
		}

		return resp
	}
}
