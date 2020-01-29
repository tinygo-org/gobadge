package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/drivers/buzzer"
)

func Music() {
	circle := color.RGBA{0, 100, 250, 255}
	white := color.RGBA{255, 255, 255, 255}
	ring := color.RGBA{200, 0, 0, 255}
	display.FillScreen(white)
	tinydraw.FilledCircle(&display, 25, 64, 8, circle) // LEFT
	tinydraw.FilledCircle(&display, 55, 64, 8, circle) // RIGHT
	tinydraw.FilledCircle(&display, 40, 49, 8, circle) // UP
	tinydraw.FilledCircle(&display, 40, 79, 8, circle) // DOWN

	tinydraw.FilledCircle(&display, 120, 20, 8, circle) // START

	tinydraw.FilledCircle(&display, 120, 70, 8, circle) // B
	tinydraw.FilledCircle(&display, 135, 55, 8, circle) // A

	for {
		pressed, _ := buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}
		if pressed&machine.BUTTON_START_MASK > 0 {
			tinydraw.Circle(&display, 120, 20, 10, ring)
			bzr.Tone(buzzer.G3, buzzer.Eighth)
			time.Sleep(100 * time.Millisecond)
		} else {
			tinydraw.Circle(&display, 120, 20, 10, white)
		}
		if pressed&machine.BUTTON_A_MASK > 0 {
			tinydraw.Circle(&display, 135, 55, 10, ring)
			bzr.Tone(buzzer.F3, buzzer.Eighth)
			time.Sleep(100 * time.Millisecond)
		} else {
			tinydraw.Circle(&display, 135, 55, 10, white)
		}
		if pressed&machine.BUTTON_B_MASK > 0 {
			tinydraw.Circle(&display, 120, 70, 10, ring)
			bzr.Tone(buzzer.E3, buzzer.Eighth)
			time.Sleep(100 * time.Millisecond)
		} else {
			tinydraw.Circle(&display, 120, 70, 10, white)
		}

		if pressed&machine.BUTTON_LEFT_MASK > 0 {
			tinydraw.Circle(&display, 25, 64, 10, ring)
			bzr.Tone(buzzer.A3, buzzer.Eighth)
			time.Sleep(100 * time.Millisecond)
		} else {
			tinydraw.Circle(&display, 25, 64, 10, white)
		}
		if pressed&machine.BUTTON_RIGHT_MASK > 0 {
			tinydraw.Circle(&display, 55, 64, 10, ring)
			bzr.Tone(buzzer.C3, buzzer.Eighth)
			time.Sleep(100 * time.Millisecond)
		} else {
			tinydraw.Circle(&display, 55, 64, 10, white)
		}
		if pressed&machine.BUTTON_UP_MASK > 0 {
			tinydraw.Circle(&display, 40, 49, 10, ring)
			bzr.Tone(buzzer.B3, buzzer.Eighth)
			time.Sleep(100 * time.Millisecond)
		} else {
			tinydraw.Circle(&display, 40, 49, 10, white)
		}
		if pressed&machine.BUTTON_DOWN_MASK > 0 {
			tinydraw.Circle(&display, 40, 79, 10, ring)
			bzr.Tone(buzzer.D3, buzzer.Eighth)
			time.Sleep(100 * time.Millisecond)
		} else {
			tinydraw.Circle(&display, 40, 79, 10, white)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
