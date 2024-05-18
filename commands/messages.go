package commands

type ErrorCode int

const (
	NameTaken ErrorCode = iota
	GameNotFound
	InvalidCode
	GameLobbyClosed
)

type CriticalErrorMsg struct {
	Error error
}

type CreateGameMsg struct {
	ErrorCode *ErrorCode
	GameCode  string
	GameID    string
}

type JoinGameMsg struct {
	ErrorCode *ErrorCode
}

type LoginMsg struct {
	ErrorCode *ErrorCode
}
