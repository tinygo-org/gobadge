package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/ws2812"
)

const (
	Red = iota
	MellonGreen
	Green
	Cyan
	Blue
	Purple
	White
	Off
)

var colors = [...]color.RGBA{
	color.RGBA{255, 0, 0, 255},     // RED
	color.RGBA{255, 255, 0, 255},   // MELLON_GREEN
	color.RGBA{0, 255, 0, 255},     // GREEN
	color.RGBA{0, 255, 255, 255},   // CYAN
	color.RGBA{0, 0, 255, 255},     // BLUE
	color.RGBA{255, 0, 255, 255},   // PURPLE
	color.RGBA{255, 255, 255, 255}, // WHITE
	color.RGBA{0, 0, 0, 255},       // OFF
}

func main() {

	// get and configure neopixels
	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	leds := ws2812.New(neo)
	// a color for each led on the board
	ledColors := make([]color.RGBA, 5)

	// get and configure buttons on the board
	buttons := shifter.NewButtons()
	buttons.Configure()

	c := 0
	for {
		// read buttons states
		buttons.ReadInput()
		switch {
		case buttons.Pins[shifter.BUTTON_LEFT].Get():
			c = Red
		case buttons.Pins[shifter.BUTTON_UP].Get():
			c = MellonGreen
		case buttons.Pins[shifter.BUTTON_DOWN].Get():
			c = Green
		case buttons.Pins[shifter.BUTTON_RIGHT].Get():
			c = Cyan
		case buttons.Pins[shifter.BUTTON_SELECT].Get():
			c = White
		case buttons.Pins[shifter.BUTTON_START].Get():
			c = Off
		case buttons.Pins[shifter.BUTTON_A].Get():
			c = Blue
		case buttons.Pins[shifter.BUTTON_B].Get():
			c = Purple
		}

		// set color for LEDs
		for i := range ledColors {
			ledColors[i] = colors[c]
		}

		leds.WriteColors(ledColors)

		time.Sleep(30 * time.Millisecond)
	}
}
