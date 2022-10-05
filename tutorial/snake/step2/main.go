package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"
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

	w, h := display.Size()
	x := (w - 10) / 2
	y := (h - 10) / 2
	for {

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

		// clear the display and paint everything black
		display.FillScreen(black)

		// draw our little snake-rectangle
		display.FillRectangle(x, y, 10, 10, green)

		time.Sleep(100 * time.Millisecond)
	}
}
