package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"

	"tinygo.org/x/drivers/ws2812"
)

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

	var k uint8
	for {
		buttons.ReadInput()
		if buttons.Pins[shifter.BUTTON_LEFT].Get() {
			k++
		}
		if buttons.Pins[shifter.BUTTON_RIGHT].Get() {
			k--
		}

		ledColors[0] = getRainbowRGB(k)
		ledColors[1] = getRainbowRGB(k + 10)
		ledColors[2] = getRainbowRGB(k + 20)
		ledColors[3] = getRainbowRGB(k + 30)
		ledColors[4] = getRainbowRGB(k + 40)
		leds.WriteColors(ledColors)

		time.Sleep(10 * time.Millisecond)
	}
}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 85 {
		return color.RGBA{i * 3, 255 - i*3, 0, 255}
	} else if i < 170 {
		i -= 85
		return color.RGBA{255 - i*3, 0, i * 3, 255}
	}
	i -= 170
	return color.RGBA{0, i * 3, 255 - i*3, 255}
}
