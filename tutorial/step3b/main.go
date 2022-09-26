package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"

	"tinygo.org/x/drivers/ws2812"
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

	var k uint8
	for {
		buttons.Read8Input()
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
