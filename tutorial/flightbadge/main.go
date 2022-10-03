package main

import (
	"image/color"
	"strings"
	"time"

	"machine"
	"machine/usb/hid/keyboard"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"

	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
)

var (
	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)

	terminal = tinyterm.NewTerminal(&display)

	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}

	font = &proggy.TinySZ8pt7b
)

var (
	shifted  bool
	lastKey  string
	lastTime time.Time
)

var logo = `
  ___ _ _      _   _      
 | __| (_)__ _| |_| |_    
 | _|| | / _\ | ' \  _|   
 |_|_|_|_\__, |_||_\__|   
 | _ ) __|___/| |__ _ ___ 
 | _ \/ _\ / _\ / _\ / -_)
 |___/\__,_\__,_\__, \___|
                |___/     
`

func main() {
	go handleDisplay()

	buttons := shifter.NewButtons()
	buttons.Configure()

	for {
		buttons.ReadInput()

		// takeoff
		if buttons.Pins[shifter.BUTTON_START].Get() {
			handleKey("[")
		}

		// land
		if buttons.Pins[shifter.BUTTON_A].Get() {
			handleKey("]")
		}

		// front flip
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
			handleKey("t")
		}

		// hold down button B to shift to access second set of arrow commands
		if buttons.Pins[shifter.BUTTON_B].Get() {
			shifted = true
		} else {
			shifted = false
		}

		if buttons.Pins[shifter.BUTTON_LEFT].Get() {
			handleShiftedKey("j", "a")
		}

		if buttons.Pins[shifter.BUTTON_UP].Get() {
			handleShiftedKey("i", "w")
		}

		if buttons.Pins[shifter.BUTTON_DOWN].Get() {
			handleShiftedKey("k", "s")
		}

		if buttons.Pins[shifter.BUTTON_RIGHT].Get() {
			handleShiftedKey("l", "d")
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func handleShiftedKey(key1, key2 string) {
	if shifted {
		handleKey(key1)
		return
	}
	handleKey(key2)
}

func handleKey(key string) {
	// simple debounce
	if key == lastKey && time.Since(lastTime) < 150*time.Millisecond {
		return
	}

	kb := keyboard.New()
	kb.Write([]byte(key))

	lastKey, lastTime = key, time.Now()
}

func handleDisplay() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})

	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	terminal.Configure(&tinyterm.Config{
		Font:              font,
		FontHeight:        10,
		FontOffset:        6,
		UseSoftwareScroll: true,
	})

	display.FillScreen(black)

	showSplash()

	input := make([]byte, 64)
	i := 0

	for {
		if machine.Serial.Buffered() > 0 {
			data, _ := machine.Serial.ReadByte()

			switch data {
			case 13:
				// return key
				terminal.Write([]byte("\r\n"))
				terminal.Write(input[:i])
				i = 0
			default:
				input[i] = data
				i++
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func showSplash() {
	for _, line := range strings.Split(strings.TrimSuffix(logo, "\n"), "\n") {
		terminal.Write([]byte(line))
	}
}
