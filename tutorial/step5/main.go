package main

import (
	"image/color"
	"machine"

	"../fonts"
	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/tinyfont"
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

	display.FillScreen(color.RGBA{0, 0, 0, 255})

	tinyfont.WriteLine(&display, &fonts.Bold12pt7b, 10, 50, []byte("Hello"), color.RGBA{R: 255, G: 255, B: 0, A: 255})
	tinyfont.WriteLine(&display, &fonts.Bold12pt7b, 40, 80, []byte("Gophers!"), color.RGBA{R: 255, G: 0, B: 255, A: 255})
}
