package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

func Accel3D() {
	white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 255, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	black := color.RGBA{0, 0, 0, 255}

	display.FillScreen(white)
	tinydraw.Rectangle(&display, 26, 30, 132, 7, black)
	tinydraw.Rectangle(&display, 26, 40, 132, 7, black)
	tinydraw.Rectangle(&display, 26, 50, 132, 7, black)

	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 80, "Move the PyBadge to see", black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 14, 90, "the accelerometer in", black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 60, 100, "action.", black)

	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 8, 36, "X:", black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 8, 46, "Y:", black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 8, 56, "Z:", black)

	x, y, z := accel.ReadRawAcceleration()
	for {
		pressed, _ := buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}

		x, y, z = accel.ReadRawAcceleration()
		x = x / 500
		y = y / 500
		z = z / 500
		if x > 64 {
			x = 64
		}
		if y > 64 {
			y = 64
		}
		if z > 64 {
			z = 64
		}
		if x < -64 {
			x = -64
		}
		if y < -64 {
			y = -64
		}
		if z < -64 {
			z = -64
		}
		display.FillRectangle(28, 32, 128, 2, white)
		display.FillRectangle(28, 42, 128, 2, white)
		display.FillRectangle(28, 52, 128, 2, white)
		if x < 0 {
			display.FillRectangle(92+x, 32, -x, 2, red)
		} else {
			display.FillRectangle(92, 32, x, 2, red)
		}
		if y < 0 {
			display.FillRectangle(92+y, 42, -y, 2, green)
		} else {
			display.FillRectangle(92, 42, y, 2, green)
		}
		if z < 0 {
			display.FillRectangle(92+z, 52, -z, 2, blue)
		} else {
			display.FillRectangle(92, 52, z, 2, blue)
		}

		println("X:", x, "Y:", y, "Z:", z)
		time.Sleep(time.Millisecond * 100)
		time.Sleep(50 * time.Millisecond)
	}
}
