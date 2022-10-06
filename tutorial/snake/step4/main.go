package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"
)

const (
	WIDTHBLOCKS  = 16
	HEIGHTBLOCKS = 13
)

var display st7735.Device

func main() {
	// Setup the SPI connection of the GoBadge
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})

	// Create a new display device
	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	// Setup the buttons
	buttons := shifter.NewButtons()
	buttons.Configure()

	green := color.RGBA{0, 255, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	// fill the whole screen with black
	display.FillScreen(black)

	x := int16(8)
	y := int16(6)
	modY := int16(10)
	for {

		// "clear" our previous snake position
		if y == 12 { // since the screen is 128 pixels in height, the last row of the grid should be 8 pixels instead of 10
			modY = 8
		} else {
			modY = 10
		}
		display.FillRectangle(x*10, y*10, 10, modY, black)

		buttons.ReadInput()
		switch {
		case buttons.Pins[shifter.BUTTON_LEFT].Get():
			x--
		case buttons.Pins[shifter.BUTTON_UP].Get():
			y--
		case buttons.Pins[shifter.BUTTON_DOWN].Get():
			y++
		case buttons.Pins[shifter.BUTTON_RIGHT].Get():
			x++
		}

		// control it doesn't get out of bounds
		if x >= WIDTHBLOCKS {
			x = 0
		}
		if x < 0 {
			x = WIDTHBLOCKS - 1
		}
		if y >= HEIGHTBLOCKS {
			y = 0
		}
		if y < 0 {
			y = HEIGHTBLOCKS - 1
		}

		// draw our little snake-rectangle in their new position
		if y == 12 { // since the screen is 128 pixels in height, the last row of the grid should be 8 pixels instead of 10
			modY = 8
		} else {
			modY = 10
		}
		display.FillRectangle(x*10, y*10, 10, modY, green)

		time.Sleep(100 * time.Millisecond)
	}
}
