package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"
)

const (
	BUTTON_LEFT = iota
	BUTTON_UP
	BUTTON_DOWN
	BUTTON_RIGHT
	BUTTON_SELECT
	BUTTON_START
	BUTTON_A
	BUTTON_B
)

var display st7735.Device
var buttons shifter.Device

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		MOSI:      machine.SPI1_MOSI_PIN,
		MISO:      machine.SPI1_MISO_PIN,
		Frequency: 8000000,
	})
	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	buttons = shifter.NewButtons()
	buttons.Configure()

	snakeGame := Game{
		colors: []color.RGBA{
			color.RGBA{0, 0, 0, 255},
			color.RGBA{0, 200, 0, 255},
			color.RGBA{250, 0, 0, 255},
			color.RGBA{160, 160, 160, 255},
		},
		snake: Snake{
			body: [208][2]int16{
				{0, 3},
				{0, 2},
				{0, 1},
			},
			length:    3,
			direction: 3,
		},
		appleX: -1,
		appleY: -1,
		status: START,
	}

	for {
		snakeGame.Start()
		time.Sleep(1 * time.Second)
	}
}
