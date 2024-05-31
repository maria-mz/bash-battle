package main

import (
	"log"

	"github.com/maria-mz/bash-battle/app"
)

func main() {
	app := app.New("127.0.0.1", 5555, "maria :)")

	defer app.Shutdown()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
