package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/tinydraw"
)

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		MOSI:      machine.SPI1_MOSI_PIN,
		MISO:      machine.SPI1_MISO_PIN,
		Frequency: 8000000,
	})

	display := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	buttons := shifter.NewButtons()
	buttons.Configure()

	display.FillScreen(color.RGBA{255, 255, 255, 255})

	circle := color.RGBA{0, 100, 250, 255}
	white := color.RGBA{255, 255, 255, 255}
	ring := color.RGBA{200, 0, 0, 255}
	display.FillScreen(white)
	tinydraw.FilledCircle(&display, 25, 74, 8, circle) // LEFT
	tinydraw.FilledCircle(&display, 55, 74, 8, circle) // RIGHT
	tinydraw.FilledCircle(&display, 40, 59, 8, circle) // UP
	tinydraw.FilledCircle(&display, 40, 89, 8, circle) // DOWN

	tinydraw.FilledCircle(&display, 45, 30, 8, circle)  // SELECT
	tinydraw.FilledCircle(&display, 120, 30, 8, circle) // START

	tinydraw.FilledCircle(&display, 120, 80, 8, circle) // B
	tinydraw.FilledCircle(&display, 135, 65, 8, circle) // A

	for {
		buttons.ReadInput()
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
			tinydraw.Circle(&display, 45, 30, 10, ring)
		} else {
			tinydraw.Circle(&display, 45, 30, 10, white)
		}
		if buttons.Pins[shifter.BUTTON_START].Get() {
			tinydraw.Circle(&display, 120, 30, 10, ring)
		} else {
			tinydraw.Circle(&display, 120, 30, 10, white)
		}
		if buttons.Pins[shifter.BUTTON_A].Get() {
			tinydraw.Circle(&display, 135, 65, 10, ring)
		} else {
			tinydraw.Circle(&display, 135, 65, 10, white)
		}
		if buttons.Pins[shifter.BUTTON_B].Get() {
			tinydraw.Circle(&display, 120, 80, 10, ring)
		} else {
			tinydraw.Circle(&display, 120, 80, 10, white)
		}
		if buttons.Pins[shifter.BUTTON_LEFT].Get() {
			tinydraw.Circle(&display, 25, 74, 10, ring)
		} else {
			tinydraw.Circle(&display, 25, 74, 10, white)
		}
		if buttons.Pins[shifter.BUTTON_RIGHT].Get() {
			tinydraw.Circle(&display, 55, 74, 10, ring)
		} else {
			tinydraw.Circle(&display, 55, 74, 10, white)
		}
		if buttons.Pins[shifter.BUTTON_UP].Get() {
			tinydraw.Circle(&display, 40, 59, 10, ring)
		} else {
			tinydraw.Circle(&display, 40, 59, 10, white)
		}
		if buttons.Pins[shifter.BUTTON_DOWN].Get() {
			tinydraw.Circle(&display, 40, 89, 10, ring)
		} else {
			tinydraw.Circle(&display, 40, 89, 10, white)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
