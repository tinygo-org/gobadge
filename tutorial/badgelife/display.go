package main

import (
	"machine"

	"image/color"
	"time"

	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/tinydraw"

	"github.com/acifani/vita/lib/game"
)

var (
	display st7735.Device

	gamebuffer []byte

	startx int16 = 24
	starty int16 = 8
	radius int16 = 2
	space  int16 = 2

	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
)

func startGame() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})

	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	display.FillScreen(black)

	gamebuffer = make([]byte, height*width)
	universe.Read(gamebuffer)

	for {
		drawGrid()
		display.Display()
		universe.Read(gamebuffer)

		universe.Tick()

		time.Sleep(10 * time.Millisecond)
	}
}

func drawGrid() {
	var rows, cols uint32
	c := black

	for rows = 0; rows < height; rows++ {
		for cols = 0; cols < width; cols++ {
			idx := universe.GetIndex(rows, cols)

			switch {
			case universe.Cell(idx) == gamebuffer[idx]:
				// no change, so skip
				continue
			case universe.Cell(idx) == game.Alive:
				c = white
			default: // game.Dead
				c = black
			}

			x := startx + int16(cols)*radius*2 - radius + int16(cols)*space
			y := starty + int16(rows)*radius*2 - radius + int16(rows)*space
			tinydraw.FilledCircle(&display, x, y, radius, c)
		}
	}
}
