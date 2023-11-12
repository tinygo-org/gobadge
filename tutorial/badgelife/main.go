package main

import (
	"time"

	"tinygo.org/x/drivers/shifter"

	"github.com/acifani/vita/lib/game"
)

var (
	universe *game.Universe

	height     uint32 = 20
	width      uint32 = 20
	population        = 20
)

func main() {
	universe = game.NewUniverse(height, width)
	universe.Randomize(population)

	go startGame()

	buttons := shifter.NewButtons()
	buttons.Configure()

	for {
		buttons.ReadInput()

		if buttons.Pins[shifter.BUTTON_A].Get() {
			universe.Randomize(population)
		}

		if buttons.Pins[shifter.BUTTON_B].Get() {
			universe.Reset()
		}

		time.Sleep(time.Millisecond * 200)
	}
}
