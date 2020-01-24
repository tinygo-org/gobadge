package main

import (
	"image/color"
	"machine"
	"time"
)

func Leds() {
	display.EnableBacklight(false)
	display.FillScreen(color.RGBA{0,0,0,255})
	ledColors := make([]color.RGBA, 5)
	var i uint8
	for {
		ledColors[0] = getRainbowRGB(i)
		ledColors[1] = getRainbowRGB(i+10)
		ledColors[2] = getRainbowRGB(i+20)
		ledColors[3] = getRainbowRGB(i+30)
		ledColors[4] = getRainbowRGB(i+40)
		leds.WriteColors(ledColors)
		
		pressed, _ := buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}
		i+=2

		time.Sleep(50*time.Millisecond)
	}

	ledColors[0] = color.RGBA{0,0,0,255}
	ledColors[1] = color.RGBA{0,0,0,255}
	ledColors[2] = color.RGBA{0,0,0,255}
	ledColors[3] = color.RGBA{0,0,0,255}
	ledColors[4] = color.RGBA{0,0,0,255}
	leds.WriteColors(ledColors)
	time.Sleep(50*time.Millisecond)
	ledColors[0] = color.RGBA{0,0,0,255}
	ledColors[1] = color.RGBA{0,0,0,255}
	ledColors[2] = color.RGBA{0,0,0,255}
	ledColors[3] = color.RGBA{0,0,0,255}
	ledColors[4] = color.RGBA{0,0,0,255}
	leds.WriteColors(ledColors)
	time.Sleep(50*time.Millisecond)

	display.EnableBacklight(true)
}
