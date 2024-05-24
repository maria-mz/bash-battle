package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maria-mz/bash-battle-proto/proto"
	be "github.com/maria-mz/bash-battle/backend"
	"github.com/maria-mz/bash-battle/context"
	"github.com/maria-mz/bash-battle/tui"
)

type App struct {
	context context.AppContext
	backend *be.Backend
	program *tea.Program
}

func NewApp(host string, port uint16, username string) (*App, error) {
	app := &App{}
	app.context = context.AppContext{}

	if err := app.initBackend(host, port); err != nil {
		return app, err
	}

	if err := app.loginUser(username); err != nil {
		return app, err
	}

	app.program = tea.NewProgram(tui.NewTui(app.context), tea.WithAltScreen())

	return app, nil
}

func (app *App) RunTui() error {
	tea.LogToFile("debug.log", "debug")

	if _, err := app.program.Run(); err != nil {
		return err
	}

	return nil
}

func (app *App) Shutdown() {
	if app.program != nil {
		app.program.Quit()
	}

	if app.backend != nil {
		app.backend.Shutdown()
	}
}

func (app *App) initBackend(host string, port uint16) error {
	serverAddr := fmt.Sprintf("%s:%d", host, port)

	backend, err := be.NewBackend(serverAddr)

	if err != nil {
		return err
	}

	app.backend = backend
	app.context.ServerAddress = serverAddr

	return nil
}

func (app *App) loginUser(username string) error {
	loginRequest := &proto.LoginRequest{
		Username: username,
	}

	resp, err := app.backend.Login(loginRequest)

	if err != nil {
		return fmt.Errorf("failed to login: %s", err)
	}

	app.context.Username = username
	app.context.ConnectionStatus = context.IsConnected
	app.context.GameStatus = context.WaitingForPlayers
	app.context.PopulateGameConfig(resp.GameConfig)
	app.context.PopulatePlayerNames(resp.Players)

	return nil
}
