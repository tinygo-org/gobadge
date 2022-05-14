package main

import (
	"image/color"
	"machine"
	"strconv"
	"time"

	"github.com/tinygo-org/gobadge/fonts"
	"tinygo.org/x/tinyfont"

	"tinygo.org/x/drivers/scd4x"
)

var ledColors = make([]color.RGBA, 5)

func CO2Monitor() {
	display.EnableBacklight(true)
	display.FillScreen(colors[WHITE])

	sensor := scd4x.New(machine.I2C0)
	sensor.Configure()

	if err := sensor.StartPeriodicMeasurement(); err != nil {
		println(err)
	}

	w32, _ := tinyfont.LineWidth(&fonts.Bold12pt7b, "CO  Monitor")
	xPPM := (WIDTH - int16(w32)) / 2
	tinyfont.WriteLine(&display, &fonts.Bold12pt7b, xPPM, 30, "CO  Monitor", colors[BLACK])
	w32, _ = tinyfont.LineWidth(&fonts.Bold12pt7b, "CO")
	tinyfont.WriteLine(&display, &fonts.TinySZ8pt7b, xPPM+int16(w32), 33, "2", colors[BLACK])
	// move it one pixel to make it looks like BOLD font
	tinyfont.WriteLine(&display, &fonts.TinySZ8pt7b, xPPM+int16(w32)+1, 33, "2", colors[BLACK])

	w32, _ = tinyfont.LineWidth(&fonts.TinySZ8pt7b, "PPM")
	xPPM = WIDTH - int16(w32) - 4
	tinyfont.WriteLine(&display, &fonts.TinySZ8pt7b, xPPM, 64, "PPM", colors[BLACK])

	w32, _ = tinyfont.LineWidth(&fonts.TinySZ8pt7b, "TEMP.")
	tinyfont.WriteLine(&display, &fonts.TinySZ8pt7b, ((WIDTH/2)-int16(w32))/2, 94, "TMP", colors[BLACK])

	w32, _ = tinyfont.LineWidth(&fonts.TinySZ8pt7b, "RH")
	tinyfont.WriteLine(&display, &fonts.TinySZ8pt7b, (WIDTH/2)+((WIDTH/2)-int16(w32))/2, 94, "RH", colors[BLACK])

	w32, _ = tinyfont.LineWidth(&fonts.TinySZ8pt7b, "ºC")
	tinyfont.WriteLine(&display, &fonts.TinySZ8pt7b, (WIDTH/2)-int16(w32)-4, 112, "ºC", colors[BLACK])

	w32, _ = tinyfont.LineWidth(&fonts.TinySZ8pt7b, "%")
	tinyfont.WriteLine(&display, &fonts.TinySZ8pt7b, WIDTH-int16(w32)-4, 110, "%", colors[BLACK])

	oldCO2 := ""
	oldTemp := ""
	oldHumidity := ""
	var xCO2 int16
	var xTemp int16
	var xHumidity int16

	for {
		co2, err := sensor.ReadCO2()
		if err != nil {
			println(err)
		}

		ShowCO2Level(co2)

		// Clear old readings
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, xCO2, 77, oldCO2, colors[WHITE])
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, xTemp, 120, oldTemp, colors[WHITE])
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, xHumidity, 120, oldHumidity, colors[WHITE])

		// Display new readings
		oldCO2 = strconv.Itoa(int(co2))
		w32, _ = tinyfont.LineWidth(&fonts.Bold24pt7b, oldCO2)
		xCO2 = xPPM - int16(w32) - 10
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, xCO2, 77, oldCO2, colors[BLACK])

		value, _ := sensor.ReadTemperature()
		println("TEMP", value)
		oldTemp = strconv.FormatFloat(float64(value), 'f', 1, 64)
		w32, _ = tinyfont.LineWidth(&fonts.Bold12pt7b, oldTemp)
		xTemp = ((WIDTH/2)-int16(w32))/2 - 6
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, xTemp, 120, oldTemp, colors[BLACK])

		value, _ = sensor.ReadHumidity()
		println("HUMD", value)
		oldHumidity = strconv.FormatFloat(float64(value), 'f', 1, 64)
		w32, _ = tinyfont.LineWidth(&fonts.Bold12pt7b, oldHumidity)
		xHumidity = (WIDTH / 2) + ((WIDTH/2)-int16(w32))/2 - 6
		tinyfont.WriteLine(&display, &fonts.Bold12pt7b, xHumidity, 120, oldHumidity, colors[BLACK])

		pressed, _ := buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}

		time.Sleep(time.Second)
	}

	Clear()
	Clear()

	time.Sleep(50 * time.Millisecond)

	display.EnableBacklight(true)
}

// ShowCO2Level shows the current CO2 level on the LEDs.
func ShowCO2Level(co2 int32) {
	if co2 < 0 {
		co2 = 0
	}
	if co2 > 2000 {
		co2 = 2000
	}
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
	howmany := int(co2/400) + 1

	// clear old colors
	for i := 0; i < len(ledColors); i++ {
		if i < howmany {
			ledColors[i] = c
		} else {
			ledColors[i] = color.RGBA{0, 0, 0, 0}
		}
	}

	leds.WriteColors(ledColors)
}

func Clear() {
	for i := 0; i < len(ledColors); i++ {
		ledColors[i] = color.RGBA{0, 0, 0, 0}
	}
	leds.WriteColors(ledColors)

	time.Sleep(50 * time.Millisecond)
}
