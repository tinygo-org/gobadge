package main

import (
	"github.com/acifani/vita/lib/game"
	"image/color"
	"time"
	"tinygo.org/x/drivers/shifter"
)

var (
	gamebuffer []byte

	universe *game.Universe

	width      uint32 = 26
	height     uint32 = 21
	population        = 20
	cellSize   int16  = 6

	wh = colors[WHITE]

	cellBuf = []color.RGBA{
		wh, wh, wh, wh, wh, wh,
		wh, wh, bk, bk, wh, wh,
		wh, bk, bk, bk, bk, wh,
		wh, bk, bk, bk, bk, wh,
		wh, wh, bk, bk, wh, wh,
		wh, wh, wh, wh, wh, wh,
	}
)

func GameOfLife() {
	white := color.RGBA{255, 255, 255, 255}
	display.FillScreen(white)

	gamebuffer = make([]byte, height*width)
	universe = game.NewUniverse(height, width)
	universe.Randomize(population)
	universe.Read(gamebuffer)

	for {
		drawGrid()
		display.Display()
		universe.Read(gamebuffer)

		universe.Tick()

		buttons.ReadInput()
		println("LOOP")
		if buttons.Pins[shifter.BUTTON_B].Get() {
			println("RESET")
			universe.Reset()
			universe.Randomize(population)
		}
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
			break
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func drawGrid() {
	var rows, cols uint32

	for rows = 0; rows < height; rows++ {
		for cols = 0; cols < width; cols++ {
			idx := universe.GetIndex(rows, cols)

			switch {
			case universe.Cell(idx) == gamebuffer[idx]:
				// no change, so skip
				continue
			case universe.Cell(idx) == game.Alive:
				display.FillRectangleWithBuffer(2+cellSize*int16(cols), 1+cellSize*int16(rows), cellSize, cellSize, cellBuf)
			default: // game.Dead
				display.FillRectangle(2+cellSize*int16(cols), 1+cellSize*int16(rows), cellSize, cellSize, colors[WHITE])
			}

		}
	}
}
