package app

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maria-mz/bash-battle-proto/proto"
	be "github.com/maria-mz/bash-battle/backend"
	"github.com/maria-mz/bash-battle/config"
	"github.com/maria-mz/bash-battle/messages"
	"github.com/maria-mz/bash-battle/status"
	"github.com/maria-mz/bash-battle/tui"
)

type App struct {
	config *config.Config

	connStatus status.ConnStatus
	gameStatus status.GameStatus

	backend *be.Backend
	players []*proto.Player

	program *tea.Program
}

func New(host string, port uint16, username string) *App {
	app := &App{}
	app.config = &config.Config{
		ServerAddr: fmt.Sprintf("%s:%d", host, port),
		Username:   username,
	}

	return app
}

func (app *App) Run() error {
	if err := app.initBackend(); err != nil {
		return err
	}

	if err := app.connectUser(); err != nil {
		return err
	}

	if err := app.joinGame(); err != nil {
		return err
	}

	if err := app.populateGameConfig(); err != nil {
		return err
	}

	if err := app.populatePlayers(); err != nil {
		return err
	}

	app.program = tea.NewProgram(tui.NewTui(*app.config), tea.WithAltScreen())

	err := app.runTui() // blocking

	return err
}

func (app *App) Shutdown() {
	if app.program != nil {
		app.program.Quit()
	}

	if app.backend != nil {
		app.backend.Shutdown()
	}
}

func (app *App) initBackend() error {
	backend, err := be.NewBackend(app.config.ServerAddr)

	if err != nil {
		return err
	}

	app.backend = backend

	return nil
}

func (app *App) connectUser() error {
	req := &proto.ConnectRequest{
		Username: app.config.Username,
	}

	err := app.backend.Connect(req)

	if err != nil {
		return fmt.Errorf("failed to connect: %s", err)
	}

	app.connStatus = status.Connected

	return nil
}

func (app *App) joinGame() error {
	err := app.backend.JoinGame()

	if err != nil {
		return fmt.Errorf("failed to join game: %s", err)
	}

	app.gameStatus = status.WaitingForPlayers

	return nil
}

func (app *App) populateGameConfig() error {
	gameConfig, err := app.backend.GetGameConfig()

	if err != nil {
		return fmt.Errorf("failed to get game config: %s", err)
	}

	app.config.GameConfig.MaxPlayers = int(gameConfig.MaxPlayers)
	app.config.GameConfig.Rounds = int(gameConfig.Rounds)
	app.config.GameConfig.RoundDuration = int(gameConfig.RoundSeconds)
	app.config.GameConfig.Difficulty = config.GameDifficulty(gameConfig.Difficulty)
	app.config.GameConfig.FileSize = config.GameFileSize(gameConfig.FileSize)

	return nil
}

func (app *App) populatePlayers() error {
	players, err := app.backend.GetPlayers()

	if err != nil {
		return fmt.Errorf("failed to get players: %s", err)
	}

	app.players = players.GetPlayers()

	return nil
}

func (app *App) runTui() error {
	tea.LogToFile("debug.log", "debug")

	go func() {
		app.sendTuiConnStatusMsg()
		app.sendTuiGameStatusMsg()
		app.sendTuiPlayerJoinedMsg(app.config.Username)
		app.sendTuiPlayerNamesMsg()
	}()

	if _, err := app.program.Run(); err != nil {
		return err
	}

	return nil
}

func (app *App) sendTuiConnStatusMsg() {
	msg := messages.ConnStatusMsg{Status: app.connStatus}
	app.sendTuiMsg(msg)
}

func (app *App) sendTuiGameStatusMsg() {
	msg := messages.GameStatusMsg{Status: app.gameStatus}
	app.sendTuiMsg(msg)
}

func (app *App) sendTuiPlayerJoinedMsg(name string) {
	msg := messages.PlayerJoinedMsg{Name: name}
	app.sendTuiMsg(msg)
}

func (app *App) sendTuiPlayerNamesMsg() {
	msg := messages.UpdatedPlayerNamesMsg{Names: app.getPlayerNames()}
	app.sendTuiMsg(msg)
}

func (app *App) sendTuiMsg(msg tea.Msg) {
	if app.program == nil {
		return
	}

	log.Printf("sending tui a message %#v", msg)
	app.program.Send(msg)
}

func (app *App) getPlayerNames() []string {
	names := make([]string, 0, len(app.players))

	for _, name := range app.players {
		names = append(names, name.Username)
	}

	return names
}
