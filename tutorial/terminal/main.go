package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/tinyfont/proggy"
	"tinygo.org/x/tinyterm"
)

const (
	// CarriageReturn represents the character code for a new line.
	CarriageReturn = 10
	// NewLine represents the character code for a new line.
	NewLine = 13

	// bufSize represents the buffer size of the input before flushing memory
	bufSize = 128
)

var (
	// uart is the name of the serial port stream
	uart = machine.Serial
	// terminal represents the board terminal stream
	terminal = tinyterm.NewTerminal(&display)

	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)

	// all the main colors in RGBA code.
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}

	// font is the font used to display the text
	font = &proggy.TinySZ8pt7b
)

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})

	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	display.FillScreen(black)

	terminal.Configure(&tinyterm.Config{
		Font:              font,
		FontHeight:        10,
		FontOffset:        6,
		UseSoftwareScroll: true,
	})

	// read the input
	input := make([]byte, bufSize)
	i := 0
	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()

			// check the type of data
			switch data {
			// new line or new character, let's print the input and flush the memory
			case CarriageReturn, NewLine:
				writeInputAndFlush(input, i)
				i = 0
			// any other character, let's echo it
			default:
				// just echo the character
				uart.WriteByte(data)
				input[i] = data

				_ = terminal.WriteByte(data)
				i++

				// out of range, let's write the input and flush
				if i >= bufSize {
					writeInputAndFlush(input, i)

					i = 0
				}
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}

// writeInputAndFlush writes input on terminal and to serial port
func writeInputAndFlush(input []byte, i int) {
	// write on the board
	terminal.Write([]byte("\r\n"))
	terminal.Write([]byte("You typed: "))
	terminal.Write(input[:i])

	// write on the serial port
	uart.Write([]byte("\r\n"))
	uart.Write([]byte("You typed: "))
	uart.Write(input[:i])

	// write new lines in terminal and to serial port
	terminal.Write([]byte("\r\n"))
	uart.Write([]byte("\r\n"))
}
