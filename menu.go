package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/shifter"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

// FilledTriangle draws a filled triangle given three points
// Code from tinydraw, but faster and may not work on all displays
func FilledTriangle(x0 int16, y0 int16, x1 int16, y1 int16, x2 int16, y2 int16, color color.RGBA) {
    if y0 > y1 {
        x0, y0, x1, y1 = x1, y1, x0, y0
    }
    if y1 > y2 {
        x1, y1, x2, y2 = x2, y2, x1, y1
    }
    if y0 > y1 {
        x0, y0, x1, y1 = x1, y1, x0, y0
    }

    if y0 == y2 { // y0 = y1 = y2 : it's a line
        a := x0
        b := x0
        if x1 < a {
            a = x1
        } else if x1 > b {
            b = x1
        }
        if x2 < a {
            a = x2
        } else if x2 > b {
            b = x2
        }
        // Line(display, a, y0, b, y0, color)
        display.DrawFastHLine(a, b, y0, color)
        return
    }

    dx01 := x1 - x0
    dy01 := y1 - y0
    dx02 := x2 - x0
    dy02 := y2 - y0
    dx12 := x2 - x1
    dy12 := y2 - y1

    sa := int16(0)
    sb := int16(0)
    a := int16(0)
    b := int16(0)

    last := y1 - 1
    if y1 == y2 {
        last = y1
    }

    y := y0
    for ; y <= last; y++ {
        a = x0 + sa/dy01
        b = x0 + sb/dy02
        sa += dx01
        sb += dx02
        display.DrawFastHLine(a, b, y, color)
    }

    sa = dx12 * (y - y1)
    sb = dx02 * (y - y0)

    for ; y <= y2; y++ {
        a = x1 + sa/dy12
        b = x0 + sb/dy02
        sa += dx12
        sb += dx02
        display.DrawFastHLine(a, b, y, color)
    }
}

func menu() int16 {
	display.FillScreen(color.RGBA{0, 0, 0, 255})
	options := []string{
		"Badge",
		"CO2 Monitor",
		"Rainbow LEDs",
		"Accelerometer",
		"Music!",
	}

	bgColor := color.RGBA{0, 40, 70, 255}
	display.FillScreen(bgColor)
	FilledTriangle(0, 128, 0, 45, 45, 0, color.RGBA{255, 255, 255, 255})
	FilledTriangle(45, 0, 0, 128, 145, 0, color.RGBA{255, 255, 255, 255})
	FilledTriangle(0, 128, 15, 128, 145, 0, color.RGBA{255, 255, 255, 255})

	FilledTriangle(0, 110, 110, 0, 117, 0, bgColor) 
	FilledTriangle(0, 110, 0, 117, 117, 0, bgColor) 

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

		if released && buttons.Pins[shifter.BUTTON_UP].Get() && selected > 0 {
			selected--
			tinydraw.FilledCircle(&display, 32, 37+10*selected, 2, color.RGBA{200, 200, 0, 255})
			tinydraw.FilledCircle(&display, 32, 37+10*(selected+1), 2, color.RGBA{255, 255, 255, 255})
		}
		if released && buttons.Pins[shifter.BUTTON_DOWN].Get() && selected < (numOpts-1) {
			selected++
			tinydraw.FilledCircle(&display, 32, 37+10*selected, 2, color.RGBA{200, 200, 0, 255})
			tinydraw.FilledCircle(&display, 32, 37+10*(selected-1), 2, color.RGBA{255, 255, 255, 255})
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
