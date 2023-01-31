package main

import (
	"time"

	"tinygo.org/x/drivers/shifter"
)

var (
	yourname string
	lastTime = time.Now()
)

func main() {
	if yourname == "" {
		yourname = "@tinygolang"
	}

	go handleDisplay()

	buttons := shifter.NewButtons()
	buttons.Configure()

	startLora()
	go loraRX()

	for {
		buttons.ReadInput()

		// yo
		if buttons.Pins[shifter.BUTTON_A].Get() {
			sendMessage("yo")
		}

		// ho
		if buttons.Pins[shifter.BUTTON_B].Get() {
			sendMessage("ho")
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func sendMessage(msg string) {
	// simple debounce
	if time.Since(lastTime) < 1000*time.Millisecond {
		return
	}

	lastTime = time.Now()

	err := loraTX([]byte(yourname + ":" + msg))
	if err != nil {
		showError(err)
	}

	showMessage([]byte(yourname + ":" + msg))
}
