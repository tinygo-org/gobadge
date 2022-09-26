package main

import (
	"image/color"
	"machine"

	"../fonts"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"

	"tinygo.org/x/drivers/st7735"
)

const (
	WIDTH  = 160
	HEIGHT = 128
	name   = "@TinyGolang"
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

	display.FillScreen(color.RGBA{255, 255, 255, 255})

	var r int16 = 8

	// black corners detail
	display.FillRectangle(0, 0, r, r, color.RGBA{0, 0, 0, 255})
	display.FillRectangle(0, HEIGHT-r, r, r, color.RGBA{0, 0, 0, 255})
	display.FillRectangle(WIDTH-r, 0, r, r, color.RGBA{0, 0, 0, 255})
	display.FillRectangle(WIDTH-r, HEIGHT-r, r, r, color.RGBA{0, 0, 0, 255})

	// round corners
	tinydraw.FilledCircle(&display, r, r, r, color.RGBA{255, 0, 0, 255})
	tinydraw.FilledCircle(&display, WIDTH-r-1, r, r, color.RGBA{255, 0, 0, 255})
	tinydraw.FilledCircle(&display, r, HEIGHT-r-1, r, color.RGBA{255, 0, 0, 255})
	tinydraw.FilledCircle(&display, WIDTH-r-1, HEIGHT-r-1, r, color.RGBA{255, 0, 0, 255})

	// top band
	display.FillRectangle(r, 0, WIDTH-2*r-1, r, color.RGBA{255, 0, 0, 255})
	display.FillRectangle(0, r, WIDTH, 26, color.RGBA{255, 0, 0, 255})

	// bottom band
	display.FillRectangle(r, HEIGHT-r-1, WIDTH-2*r-1, r+1, color.RGBA{255, 0, 0, 255})
	display.FillRectangle(0, HEIGHT-2*r-1, WIDTH, r, color.RGBA{255, 0, 0, 255})

	// top text : my NAME is
	w32, _ := tinyfont.LineWidth(&fonts.Regular12pt7b, []byte("my NAME is"))
	tinyfont.WriteLine(&display, &fonts.Regular12pt7b, (WIDTH-int16(w32))/2, 24, []byte("my NAME is"), color.RGBA{255, 255, 255, 255})

	// middle text
	w32, _ = tinyfont.LineWidth(&fonts.Bold9pt7b, []byte(name))
	tinyfont.WriteLine(&display, &fonts.Bold9pt7b, (WIDTH-int16(w32))/2, 72, []byte(name), color.RGBA{0, 0, 0, 255})

	// gophers fonts
	tinyfont.WriteLine(&display, &fonts.Regular32pt, WIDTH-48, 110, []byte("BE"), color.RGBA{0, 0, 0, 255})
}
