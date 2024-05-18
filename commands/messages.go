package commands

type ErrorCode int

const (
	NoError ErrorCode = iota
	NameTaken
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
