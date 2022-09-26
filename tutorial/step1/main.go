package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttons := shifter.NewButtons()
	buttons.Configure()

	for {
		buttons.Read8Input()
		if buttons.Pins[shifter.BUTTON_START].Get() {
			led.High()
		} else {
			led.Low()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
