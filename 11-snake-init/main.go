package main

import (
	"time"

	"github.com/just-purple/brownies/11-snake-init/app"
)

func main() {

	// TODO: press any key to start

	// MAYBE: this could return a score :)
	app := &app.App{
		PlayGround: app.Pair{X: 30, Y: 10},
		Snake:      app.Pair{X: 1, Y: 1},
		Direction:  app.Pair{X: 1, Y: 0},
		Delay:      100 * time.Millisecond,
	}

	app.Run()

	// MAYBE: make the user add a name and store the top 10 scores somewhere :)

	// TODO: press any key to restart
}
