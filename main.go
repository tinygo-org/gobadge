package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
	"tinygo.org/x/drivers/lis3dh"

	"tinygo.org/x/drivers/ws2812"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"
)

var display st7735.Device
var buttons shifter.Device
var leds ws2812.Device
var bzr buzzer.Device
var accel lis3dh.Device
var snakeGame Game

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		MOSI:      machine.SPI1_MOSI_PIN,
		MISO:      machine.SPI1_MISO_PIN,
		Frequency: 8000000,
	})
	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.SCL_PIN, SDA: machine.SDA_PIN})

	accel = lis3dh.New(machine.I2C0)
	accel.Address = lis3dh.Address0
	accel.Configure()
	println(accel.Connected())

	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	buttons = shifter.New(shifter.EIGHT_BITS, machine.BUTTON_LATCH, machine.BUTTON_CLK, machine.BUTTON_OUT)
	buttons.Configure()

	neo := machine.NEOPIXELS
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds = ws2812.New(neo)

	speaker := machine.SPEAKER_ENABLE
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker.High()

	bzrPin := machine.A0
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bzr = buzzer.New(bzrPin)

	Music()

	snakeGame = Game{
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

	for {
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
		default:
			break
		}
		println("LOOP")
		time.Sleep(1 * time.Second)
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
