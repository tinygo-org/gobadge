package main

import (
	"image/color"
	"machine"
	"time"
)

func Accel3D() {
	white := color.RGBA{255, 255, 255, 255}
	display.FillScreen(white)

	for {
		pressed, _ := buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}

		x, y, z := accel.ReadRawAcceleration()
		println("X:", x, "Y:", y, "Z:", z)
		time.Sleep(time.Millisecond * 100)
		time.Sleep(50 * time.Millisecond)
	}
}
