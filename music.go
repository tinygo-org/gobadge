package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/tinygo-org/gobadge/fonts"
	"tinygo.org/x/tinyfont"
)

func Music() {
	white := color.RGBA{255, 255, 255, 255}
	display.FillScreen(white)

	tinyfont.WriteLine(&display, &fonts.Bold24pt7b, 0, 50, "MUSIC", color.RGBA{0, 100, 250, 255})
	tinyfont.WriteLine(&display, &fonts.Bold9pt7b, 20, 100, "Press any key", color.RGBA{200, 0, 0, 255})

	for {
		pressed, _ := buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}

		if pressed&machine.BUTTON_START_MASK > 0 {
			tone(5274)
		}
		if pressed&machine.BUTTON_A_MASK > 0 {
			tone(1046)
		}
		if pressed&machine.BUTTON_B_MASK > 0 {
			tone(1975)
		}

		if pressed&machine.BUTTON_LEFT_MASK > 0 {
			tone(329)
		}
		if pressed&machine.BUTTON_RIGHT_MASK > 0 {
			tone(739)
		}
		if pressed&machine.BUTTON_UP_MASK > 0 {
			tone(369)
		}
		if pressed&machine.BUTTON_DOWN_MASK > 0 {
			tone(523)
		}
	}
}

func tone(tone int) {
	for i := 0; i < 10; i++ {
		bzrPin.High()
		time.Sleep(time.Duration(tone) * time.Microsecond)

		bzrPin.Low()
		time.Sleep(time.Duration(tone) * time.Microsecond)
	}
}
