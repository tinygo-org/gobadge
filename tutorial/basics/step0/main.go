package main

import (
	"machine"
	"time"
)

func main() {

	// get the pin that represents that LED on the back of the device D13
	led := machine.LED

	// configuring the LED pin mode for output
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		// turn off the LED
		led.Low()

		// wait 500 ms
		time.Sleep(time.Millisecond * 500)

		// turn on the LED
		led.High()

		// wait 500 ms
		time.Sleep(time.Millisecond * 500)
	}
}
