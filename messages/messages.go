// These are messages that the App (backend) sends to the TUI (frontend)
package messages

import "github.com/maria-mz/bash-battle/status"

type PlayerJoinedMsg struct {
	Name string
}

type PlayerLeftMsg struct {
	Name string
}

type UpdatedPlayerNamesMsg struct {
	Names []string
}

type ConnStatusMsg struct {
	Status status.ConnStatus
}

type GameStatusMsg struct {
	Status status.GameStatus
}
