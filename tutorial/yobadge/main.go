package main

import (
	"time"

	"tinygo.org/x/drivers/shifter"
)

var (
	yourname       string
	lastButtonPush time.Time
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

	lastButtonPush = time.Now()

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

		time.Sleep(500 * time.Millisecond)
	}
}

func sendMessage(msg string) {
	// simple debounce
	if time.Since(lastButtonPush) < 1*time.Second {
		return
	}

	lastButtonPush = time.Now()

	err := loraTX([]byte(yourname + ": " + msg + "!"))
	if err != nil {
		println(err.Error())
		showError(err)
	}

	showMessage([]byte(yourname + ": " + msg))
}
