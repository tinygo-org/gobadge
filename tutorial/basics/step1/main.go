package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
)

func main() {

	// get the pin that represents that LED on the back of the device D13
	led := machine.LED

	// configuring the LED pin mode for output
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// get the buttons on the device, and configure them
	buttons := shifter.NewButtons()
	buttons.Configure()

	for {
		// read the latest input from the badge for all the buttons
		buttons.ReadInput()
		// was the start button pressed?
		// if so, turn on the LED, otherwise turn it off
		// to get other button states see:
		// https://github.com/tinygo-org/drivers/blob/v0.23.0/shifter/pybadge.go#L8-L17
		if buttons.Pins[shifter.BUTTON_START].Get() {
			led.High()
		} else {
			led.Low()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
