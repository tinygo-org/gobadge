package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/shifter"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

func menu() int16 {
	display.FillScreen(color.RGBA{0, 0, 0, 255})
	options := []string{
		"Badge",
		"CO2 Monitor",
		"Rainbow LEDs",
		"Accelerometer",
		"Music!",
	}

	selected := int16(0)
	numOpts := int16(len(options))
	for i := int16(0); i < numOpts; i++ {
		tinydraw.Circle(&display, 32, 37+10*i, 4, color.RGBA{255, 255, 255, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40, 40+10*i, options[i], color.RGBA{255, 255, 255, 255})
	}

	tinydraw.FilledCircle(&display, 32, 37, 2, color.RGBA{155, 155, 255, 255})

	released := true
	for {
		pressed, _ := buttons.ReadInput()

		if released && buttons.Pins[shifter.BUTTON_UP].Get() && selected > 0 {
			selected--
			tinydraw.FilledCircle(&display, 32, 37+10*selected, 2, color.RGBA{155, 155, 255, 255})
			tinydraw.FilledCircle(&display, 32, 37+10*(selected+1), 2, color.RGBA{0, 0, 0, 255})
		}
		if released && buttons.Pins[shifter.BUTTON_DOWN].Get() && selected < (numOpts-1) {
			selected++
			tinydraw.FilledCircle(&display, 32, 37+10*selected, 2, color.RGBA{155, 155, 255, 255})
			tinydraw.FilledCircle(&display, 32, 37+10*(selected-1), 2, color.RGBA{0, 0, 0, 255})
		}
		if released && buttons.Pins[shifter.BUTTON_START].Get() {
			break
		}
		if pressed == 0 {
			released = true
		} else {
			released = false
		}
		time.Sleep(200 * time.Millisecond)
	}
	return selected
}
