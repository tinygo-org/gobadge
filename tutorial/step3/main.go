package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/ws2812"
)

var colors = [8]color.RGBA{
	color.RGBA{255, 0, 0, 255},
	color.RGBA{255, 255, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 255, 255, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{255, 0, 255, 255},
	color.RGBA{255, 255, 255, 255},
	color.RGBA{0, 0, 0, 255},
}

func main() {
	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	leds := ws2812.New(neo)
	ledColors := make([]color.RGBA, 5)

	buttons := shifter.NewButtons()
	buttons.Configure()

	c := 0
	for {
		buttons.ReadInput()
		if buttons.Pins[shifter.BUTTON_LEFT].Get() {
			c = 0
		}
		if buttons.Pins[shifter.BUTTON_UP].Get() {
			c = 1
		}
		if buttons.Pins[shifter.BUTTON_DOWN].Get() {
			c = 2
		}
		if buttons.Pins[shifter.BUTTON_RIGHT].Get() {
			c = 3
		}
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
			c = 6
		}
		if buttons.Pins[shifter.BUTTON_START].Get() {
			c = 7
		}
		if buttons.Pins[shifter.BUTTON_A].Get() {
			c = 4
		}
		if buttons.Pins[shifter.BUTTON_B].Get() {
			c = 5
		}

		for i := 0; i < 5; i++ {
			ledColors[i] = colors[c]
		}

		leds.WriteColors(ledColors)
		time.Sleep(time.Millisecond * 30)
	}
}
