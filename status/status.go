package status

type GameStatus int

const (
	Initializing GameStatus = iota
	WaitingForPlayers
)

type ConnStatus int

const (
	Connecting ConnStatus = iota
	Connected
	Disconnected
)
