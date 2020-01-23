package main

import (
	"image/color"
	"machine"

	"github.com/conejoninja/pybadge/fonts"
	"github.com/conejoninja/tinydraw"
	"github.com/conejoninja/tinyfont"
)

const (
	WIDTH  = 160
	HEIGHT = 128
)

const (
	BLACK = iota
	WHITE
	RED
)

var colors = []color.RGBA{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{255, 255, 255, 255},
	color.RGBA{255, 0, 0, 255},
}

var rainbow []color.RGBA
var pressed uint8

func badge() {
	display.FillScreen(colors[BLACK])

	rainbow = make([]color.RGBA, 256)
	for i := 0; i < 256; i++ {
		rainbow[i] = getRainbowRGB(uint8(i))
	}

	blinkyRainbow("technologist", "FOR HIRE")
	myNameIsRainbow("@_conejo")

	/*for {
		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}*/
}

func myNameIs(name string) {
	display.FillScreen(colors[WHITE])

	var r int16 = 8

	// black corners detail
	display.FillRectangle(0, 0, r, r, colors[BLACK])
	display.FillRectangle(0, HEIGHT-r, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, 0, r, r, colors[BLACK])
	display.FillRectangle(WIDTH-r, HEIGHT-r, r, r, colors[BLACK])

	// round corners
	tinydraw.FilledCircle(&display, r, r, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, r, r, colors[RED])
	tinydraw.FilledCircle(&display, r, HEIGHT-r-1, r, colors[RED])
	tinydraw.FilledCircle(&display, WIDTH-r-1, HEIGHT-r-1, r, colors[RED])

	// top band
	tinydraw.FilledRectangle(&display, r, 0, WIDTH-2*r-1, r, colors[RED])
	tinydraw.FilledRectangle(&display, 0, r, WIDTH, 26, colors[RED])

	// bottom band
	tinydraw.FilledRectangle(&display, r, HEIGHT-r-1, WIDTH-2*r-1, r+1, colors[RED])
	tinydraw.FilledRectangle(&display, 0, HEIGHT-2*r-1, WIDTH, r, colors[RED])

	// top text : my NAME is
	w32, _ := tinyfont.LineWidth(&fonts.Regular12pt7b, []byte("my NAME is"))
	tinyfont.WriteLine(&display, &fonts.Regular12pt7b, (WIDTH-int16(w32))/2, 24, []byte("my NAME is"), colors[WHITE])

	// middle text
	w32, _ = tinyfont.LineWidth(&fonts.Bold9pt7b, []byte(name))
	tinyfont.WriteLine(&display, &fonts.Bold9pt7b, (WIDTH-int16(w32))/2, 70, []byte(name), colors[BLACK])

	tinyfont.WriteLineColors(&display, &fonts.Regular32pt, WIDTH-48, 110, []byte("BE"), []color.RGBA{getRainbowRGB(100), getRainbowRGB(200)})
}

func myNameIsRainbow(name string) {
	myNameIs(name)

	w32, _ := tinyfont.LineWidth(&fonts.Bold9pt7b, []byte(name))
	for i := 0; i < 230; i++ {
		tinyfont.WriteLineColors(&display, &fonts.Bold9pt7b, (WIDTH-int16(w32))/2, 70, []byte(name), rainbow[i:])
		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}
	}
}

func blinky(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, []byte(topline))
	w32bottom, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, []byte(bottomline))
	for i := int16(0); i < 10; i++ {
		// show black text
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32top))/2, 50, []byte(topline), colors[BLACK])
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 100, []byte(bottomline), colors[BLACK])

		// repeat the other way around
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32top))/2, 50, []byte(topline), colors[WHITE])
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 100, []byte(bottomline), colors[WHITE])

		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}
	}
}

func blinkyRainbow(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, []byte(topline))
	w32bottom, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, []byte(bottomline))
	for i := int16(0); i < 20; i++ {
		// show black text
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32top))/2, 50, []byte(topline), getRainbowRGB(uint8(i*12)))
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 100, []byte(bottomline), getRainbowRGB(uint8(i*12)))

		pressed, _ = buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}
	}
}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 85 {
		return color.RGBA{i * 3, 255 - i*3, 0, 255}
	} else if i < 170 {
		i -= 85
		return color.RGBA{255 - i*3, 0, i * 3, 255}
	}
	i -= 170
	return color.RGBA{0, i * 3, 255 - i*3, 255}
}
