package main

import (
	"log"

	"github.com/maria-mz/bash-battle/app"
)

func main() {
	app, err := app.NewApp("127.0.0.1", 5555, "maria")

	defer app.Shutdown()

	if err != nil {
		log.Fatalf("Hmm, something went wrong: %s", err)
	}

	if err = app.RunTui(); err != nil {
		log.Fatalf("Failed to run TUI %s", err)
	}
}
