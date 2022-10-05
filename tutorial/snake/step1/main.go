package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7735"
)

var display st7735.Device

func main() {
	// Setup the SPI connection of the GoBadge
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})

	// Create a new display device
	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	green := color.RGBA{0, 255, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	// fill the whole screen with black
	display.FillScreen(black)

	w, h := display.Size()
	// draw a green 10x10 rectangle at the middle of the screen
	display.FillRectangle((w-10)/2, (h-10)/2, 10, 10, green)
	for {
		time.Sleep(1 * time.Second)
	}
}
