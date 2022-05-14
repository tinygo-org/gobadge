//go:build co2
// +build co2

package main

import (
	"image/color"
	"machine"
	"strconv"
	"time"

	"tinygo.org/x/drivers/scd4x"

	"github.com/tinygo-org/gobadge/fonts"
	"tinygo.org/x/tinyfont"
)

var ledColors = make([]color.RGBA, 5)

func CO2Monitor() {
	display.EnableBacklight(true)
	display.FillScreen(color.RGBA{0,0,0,255})
	
	sensor := scd4x.New(machine.I2C0)
	sensor.Configure()

	if err := sensor.StartPeriodicMeasurement(); err != nil {
		println(err)
	}

	for {
		co2, err := sensor.ReadCO2()
		if err != nil {
			println(err)
		}

		DisplayCO2("CO2", strconv.Itoa(int(co2)))
		ShowCO2Level(int32(co2))

		pressed, _ := buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}

		time.Sleep(time.Second)
	}

	Clear()
	Clear()

	time.Sleep(50*time.Millisecond)

	display.EnableBacklight(true)
}

func ShowCO2Level(co2 int32)  {
	// color
	var c color.RGBA
	switch {
	case co2 < 800:
		c = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
	case co2 < 1500:
		c = color.RGBA{R: 0xff, G: 0xff, B: 0x00}
	default:
		c = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
	}

	// how many to light up
	howmany := int(Rescale(co2, 0, 1600, 0, int32(len(ledColors))))

	for i := 0; i < howmany; i++ {
		ledColors[i] = c
	}

	leds.WriteColors(ledColors)	
}

func Clear() {
	for i := 0; i < len(ledColors); i++ {
		ledColors[i] = color.RGBA{0,0,0,0}
	}
	leds.WriteColors(ledColors)
	time.Sleep(50*time.Millisecond)
}

// Rescale performs a direct linear rescaling of an integer from one scale to another.
//
// For example:
//
//		val := gopherbot.Rescale(25, 0, 100, 0, 10)
//
// This re-scales the number 25 from a scale that ranges from 0-100,
// to a scale that ranges from 0-10.
func Rescale(input, fromMin, fromMax, toMin, toMax int32) int32 {
	return (input-fromMin)*(toMax-toMin)/(fromMax-fromMin) + toMin
}

func DisplayCO2(topline, bottomline string) {
	display.FillScreen(colors[WHITE])

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, topline)
	w32bottom, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, bottomline)
	//for i := int16(0); i < 10; i++ {
		// show black text
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32top))/2, 50, topline, colors[BLACK])
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 100, bottomline, colors[BLACK])

		// // repeat the other way around
		// tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32top))/2, 50, topline, colors[WHITE])
		// tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (WIDTH-int16(w32bottom))/2, 100, bottomline, colors[WHITE])

		// pressed, _ = buttons.Read8Input()
		// if pressed&machine.BUTTON_SELECT_MASK > 0 {
		// 	quit = true
		// 	break
		// }
	//}
}
