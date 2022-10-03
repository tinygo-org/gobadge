package main

import (
	"time"

	"machine/usb/hid/keyboard"


	"tinygo.org/x/drivers/shifter"
)

var (
	shifted bool
	lastKey string
	lastTime time.Time
)

func main() {
	buttons := shifter.NewButtons()
	buttons.Configure()

	for {
		buttons.ReadInput()

		// takeoff
		if buttons.Pins[shifter.BUTTON_START].Get() {
			handleKey("[")
		}

		// land
		if buttons.Pins[shifter.BUTTON_A].Get() {
			handleKey("]")
		}

		// front flip
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
			handleKey("t")
		}

		// hold down button B to shift to access second set of arrow commands
		if buttons.Pins[shifter.BUTTON_B].Get() {
			shifted = true
		} else {
			shifted = false
		}

		if buttons.Pins[shifter.BUTTON_LEFT].Get() {
			handleShiftedKey("j", "a")
		}

		if buttons.Pins[shifter.BUTTON_UP].Get() {
			handleShiftedKey("i", "w")
		}

		if buttons.Pins[shifter.BUTTON_DOWN].Get() {
			handleShiftedKey("k", "s")
		}

		if buttons.Pins[shifter.BUTTON_RIGHT].Get() {
			handleShiftedKey("l", "d")
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func handleShiftedKey(key1, key2 string) {
	if shifted {
		handleKey(key1)
		return
	}
	handleKey(key2)
}

func handleKey(key string) {
	// simple debounce
	if key == lastKey && time.Since(lastTime) < 150 * time.Millisecond {
		return
	}

	kb := keyboard.New()
	kb.Write([]byte(key))

	lastKey, lastTime = key, time.Now()
}
