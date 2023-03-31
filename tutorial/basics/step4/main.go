package main

import (
	"machine"

	"image/color"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var (
	Red   = color.RGBA{255, 0, 0, 255}
	Green = color.RGBA{0, 255, 0, 255}
)

func main() {

	// get and configure neopixels
	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	leds := ws2812.New(neo)
	ledColors := make([]color.RGBA, 5)

	// Initialize the Analog to digital subsystem
	machine.InitADC()
	light := machine.ADC{machine.LIGHTSENSOR}
	light.Configure(machine.ADCConfig{})

	for {

		// clamp the light value from the sensor
		value := uint8(light.Get() / 256)
		c := Green
		if value < 25 {
			c = Red
		}
		for i := 0; i < 5; i++ {
			ledColors[i] = c
		}
		leds.WriteColors(ledColors)

		time.Sleep(100 * time.Millisecond)
	}
}
