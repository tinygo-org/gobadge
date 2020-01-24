package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/conejoninja/tinydraw"
	"github.com/conejoninja/tinyfont"
	"github.com/conejoninja/tinyfont/proggy"
)

func menu() int16 {
	display.FillScreen(color.RGBA{0, 0, 0, 255})
	options := [][]byte{
		[]byte("Badge"),
		[]byte("Snake Game"),
		[]byte("Rainbow LEDs"),
		[]byte("Accel 3D"),
		//[]byte("Music!"),
	}

	selected := int16(0)
	numOpts := int16(len(options))
	for i := int16(0); i < numOpts; i++ {
		tinydraw.Circle(&display, 32, 37+10*i, 4, color.RGBA{255, 255, 255, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40, 40+10*i, options[i], color.RGBA{255, 255, 255, 255})
	}

	tinydraw.FilledCircle(&display, 32, 37, 2, color.RGBA{155, 155, 255, 255})

	var pressed uint8
	var oldPressed uint8
	for {
		pressed, _ = buttons.Read8Input()
		if pressed != oldPressed && pressed&machine.BUTTON_UP_MASK > 0 && selected > 0 {
			selected--
			tinydraw.FilledCircle(&display, 32, 37+10*selected, 2, color.RGBA{155, 155, 255, 255})
			tinydraw.FilledCircle(&display, 32, 37+10*(selected+1), 2, color.RGBA{0, 0, 0, 255})
		}
		if pressed != oldPressed && pressed&machine.BUTTON_DOWN_MASK > 0 && selected < (numOpts-1) {
			selected++
			tinydraw.FilledCircle(&display, 32, 37+10*selected, 2, color.RGBA{155, 155, 255, 255})
			tinydraw.FilledCircle(&display, 32, 37+10*(selected-1), 2, color.RGBA{0, 0, 0, 255})
		}
		if pressed != oldPressed && pressed&machine.BUTTON_START_MASK > 0 {
			break
		}
		oldPressed = pressed
		time.Sleep(200 * time.Millisecond)
	}
	return selected
}
