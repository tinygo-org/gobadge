package main

import (
	"image/color"
	"time"

	"github.com/tinygo-org/tinygo/src/machine"
	"tinygo.org/x/drivers/ws2812"
)

func main() {
	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	leds := ws2812.New(neo)
	ledColors := make([]color.RGBA, 5)

	machine.InitADC()
	light := machine.ADC{machine.LIGHTSENSOR}
	light.Configure()

	for {
		value := uint8(light.Get() / 256)
		c := color.RGBA{0, 255, 0, 255}
		if value < 25 {
			c = color.RGBA{255, 0, 0, 255}
		}
		for i := 0; i < 5; i++ {
			ledColors[i] = c
		}
		leds.WriteColors(ledColors)

		time.Sleep(100 * time.Millisecond)
	}
}
