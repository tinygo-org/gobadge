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
		"Snake",
		"Rainbow LEDs",
		"Accelerometer",
		"Game of Life",
		"Music!",
	}

	bgColor := color.RGBA{0, 40, 70, 255}
	display.FillScreen(bgColor)
	tinydraw.FilledTriangle(&display, 0, 128, 0, 45, 45, 0, color.RGBA{255, 255, 255, 255})
	tinydraw.FilledTriangle(&display, 45, 0, 0, 128, 145, 0, color.RGBA{255, 255, 255, 255})
	tinydraw.FilledTriangle(&display, 0, 128, 15, 128, 145, 0, color.RGBA{255, 255, 255, 255})
	for i := int16(0); i < 8; i++ {
		tinydraw.Line(&display, 0, 110+i, 110+i, 0, bgColor)
	}

	selected := int16(0)
	numOpts := int16(len(options))
	for i := int16(0); i < numOpts; i++ {
		tinydraw.Circle(&display, 32, 37+10*i, 4, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 39, 39+10*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 39, 40+10*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 39, 41+10*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40, 41+10*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41, 41+10*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41, 40+10*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 41, 39+10*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40, 39+10*i, options[i], color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 40, 40+10*i, options[i], color.RGBA{250, 250, 0, 255})
	}

	tinydraw.FilledCircle(&display, 32, 37, 2, color.RGBA{200, 200, 0, 255})

	released := true
	for {
		pressed, _ := buttons.ReadInput()

		if released && buttons.Pins[shifter.BUTTON_UP].Get(){
			prevSel := selected
			selected = (numOpts + selected -1) % numOpts
			tinydraw.FilledCircle(&display, 32, 37+10*selected, 2, color.RGBA{200, 200, 0, 255})
			tinydraw.FilledCircle(&display, 32, 37+10*prevSel, 2, color.RGBA{255, 255, 255, 255})
		}
		if released && buttons.Pins[shifter.BUTTON_DOWN].Get() {
			prevSel := selected
			selected = (selected +1 ) % numOpts
			tinydraw.FilledCircle(&display, 32, 37+10*selected, 2, color.RGBA{200, 200, 0, 255})
			tinydraw.FilledCircle(&display, 32, 37+10*prevSel, 2, color.RGBA{255, 255, 255, 255})
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
