package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
)

var bzrPin machine.Pin

func main() {
	speaker := machine.SPEAKER_ENABLE
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker.High()

	bzrPin = machine.A0
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttons := shifter.NewButtons()
	buttons.Configure()

	for {
		buttons.ReadInput()
		if buttons.Pins[shifter.BUTTON_LEFT].Get() {
			tone(329)
		}
		if buttons.Pins[shifter.BUTTON_UP].Get() {
			tone(369)
		}
		if buttons.Pins[shifter.BUTTON_DOWN].Get() {
			tone(523)
		}
		if buttons.Pins[shifter.BUTTON_RIGHT].Get() {
			tone(739)
		}
		if buttons.Pins[shifter.BUTTON_A].Get() {
			tone(1046)
		}
		if buttons.Pins[shifter.BUTTON_B].Get() {
			tone(1975)
		}
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
			tone(2637)
		}
		if buttons.Pins[shifter.BUTTON_START].Get() {
			tone(5274)
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
