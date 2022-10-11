package main

import (
	"image/color"
	"machine"

	//"time"

	"tinygo.org/x/drivers/lis3dh"

	"tinygo.org/x/drivers/ws2812"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"
)

var display st7735.Device
var buttons shifter.Device
var leds ws2812.Device
var bzrPin machine.Pin
var accel lis3dh.Device
var snakeGame = Game{
	colors: []color.RGBA{
		color.RGBA{0, 0, 0, 255},
		color.RGBA{0, 200, 0, 255},
		color.RGBA{250, 0, 0, 255},
		color.RGBA{160, 160, 160, 255},
	},
	snake: Snake{
		body: [208][2]int16{
			{0, 3},
			{0, 2},
			{0, 1},
		},
		length:    3,
		direction: 3,
	},
	appleX: -1,
	appleY: -1,
	status: START,
}

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})
	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.SCL_PIN, SDA: machine.SDA_PIN})

	accel = lis3dh.New(machine.I2C0)
	accel.Address = lis3dh.Address0
	accel.Configure()

	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	buttons = shifter.NewButtons()
	buttons.Configure()

	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds = ws2812.New(neo)

	bzrPin = machine.A0
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	speaker := machine.SPEAKER_ENABLE
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker.High()

	for {
		Badge()
		switch menu() {
		case 0:
			Badge()
			break
		case 1:
			snakeGame.Start()
			break
		case 2:
			Leds()
			break
		case 3:
			Accel3D()
			break
		case 4:
			Music()
			break
		default:
			break
		}
		//time.Sleep(10 * time.Millisecond)
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
