package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

func main() {
	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds := ws2812.New(neo)
	ledColors := make([]color.RGBA, 5)

	rg := false
	for {
		for i := 0; i < 5; i++ {
			if rg {
				ledColors[i] = color.RGBA{255, 0, 0, 255}
			} else {
				ledColors[i] = color.RGBA{0, 255, 0, 255}
			}
			rg = !rg
		}
		leds.WriteColors(ledColors)
		time.Sleep(time.Millisecond * 300)
	}
}
