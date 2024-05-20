package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	pb "github.com/maria-mz/bash-battle-proto/proto"
	be "github.com/maria-mz/bash-battle/backend"
)

// CriticalErrMsg represents a critical error message. If we receive this
// should probably shut down the program
type CriticalErrMsg struct {
	Error error
}

// CmdBuilder wraps the backend to prepare commands that integrate well with
// Bubble Tea framework
type CmdBuilder struct {
	backend *be.Backend
}

// NewCmdBuilder creates a new CmdBuilder
func NewCmdBuilder(backend *be.Backend) *CmdBuilder {
	return &CmdBuilder{backend: backend}
}

// NewLoginCmd returns a command that handles a login request
func (b *CmdBuilder) NewLoginCmd(req *pb.LoginRequest) func() tea.Msg {
	return func() tea.Msg {
		resp, err := b.backend.Login(req)

		if err != nil {
			return CriticalErrMsg{Error: err}
		}

		return resp
	}
}

// NewCreateGameCmd returns a command that handles a game creation request
func (b *CmdBuilder) NewCreateGameCmd(req *pb.CreateGameRequest) func() tea.Msg {
	return func() tea.Msg {
		resp, err := b.backend.CreateGame(req)

		if err != nil {
			return CriticalErrMsg{Error: err}
		}

		return resp
	}
}

// NewJoinGameCmd returns a command that handles a game join request
func (b *CmdBuilder) NewJoinGameCmd(req *pb.JoinGameRequest) func() tea.Msg {
	return func() tea.Msg {
		resp, err := b.backend.JoinGame(req)

		if err != nil {
			return CriticalErrMsg{Error: err}
		}

		return resp
	}
}
