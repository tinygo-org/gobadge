package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func main() {
	// Setup the screen pins
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})
	// Setup the display
	display := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	// Clear the screen to black
	display.FillScreen(color.RGBA{0, 0, 0, 255})

	// Write "Hello" 10 pixels from the right and 50 pixels from the top, in mellon green
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, 50, "Hello", color.RGBA{R: 255, G: 255, B: 0, A: 255})
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 40, 80, "Gophers!", color.RGBA{R: 255, G: 0, B: 255, A: 255})
}
