package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const (
	NumberOfLEDs = 5
)

var (
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
)

func main() {
	// get the LED strip pin for the NEOPixels
	neo := machine.NEOPIXELS
	// configure the pins for output
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// configure LED strip driver
	leds := ws2812.New(neo)
	ledColors := make([]color.RGBA, NumberOfLEDs)

	// rg represents if the LED is going to be red or green
	rg := false
	for {
		for i := 0; i < NumberOfLEDs; i++ {
			if rg {
				ledColors[i] = red
			} else {
				ledColors[i] = green
			}
			// swap color for the next LED
			rg = !rg
		}
		// set the color of the LEDs
		leds.WriteColors(ledColors)
		// sleep for 300 ms
		time.Sleep(time.Millisecond * 300)
	}
}
