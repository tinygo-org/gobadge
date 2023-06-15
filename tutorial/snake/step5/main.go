package main

import (
	"image/color"
	"machine"
	"math/rand"
	"time"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"
)

const (
	WIDTHBLOCKS  = 16
	HEIGHTBLOCKS = 13
)

const (
	SnakeUp = iota
	SnakeDown
	SnakeLeft
	SnakeRight
)

var display st7735.Device

type Snake struct {
	body      [3][2]int16
	direction int16
}

var snake = Snake{
	body: [3][2]int16{
		{0, 3},
		{0, 2},
		{0, 1},
	},
	direction: 3,
}

var appleX = int16(-1)
var appleY = int16(-1)

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var black = color.RGBA{0, 0, 0, 255}

func main() {
	// Setup the SPI connection of the GoBadge
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})

	// Create a new display device
	display = st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	// Setup the buttons
	buttons := shifter.NewButtons()
	buttons.Configure()

	// fill the whole screen with black
	display.FillScreen(black)

	drawSnake()
	for {
		buttons.ReadInput()
		switch {
		// add some checks so the snake doesn't go backwards
		case buttons.Pins[shifter.BUTTON_LEFT].Get():
			if snake.direction != SnakeRight {
				snake.direction = SnakeLeft
			}
		case buttons.Pins[shifter.BUTTON_UP].Get():
			if snake.direction != SnakeDown {
				snake.direction = SnakeUp
			}
		case buttons.Pins[shifter.BUTTON_DOWN].Get():
			if snake.direction != SnakeUp {
				snake.direction = SnakeDown
			}
		case buttons.Pins[shifter.BUTTON_RIGHT].Get():
			if snake.direction != SnakeLeft {
				snake.direction = SnakeRight
			}
		}

		moveSnake()
		time.Sleep(100 * time.Millisecond)
	}
}

func moveSnake() {
	// get the coords of the head
	x := snake.body[0][0]
	y := snake.body[0][1]

	switch snake.direction {
	case SnakeLeft:
		x--
		break
	case SnakeUp:
		y--
		break
	case SnakeDown:
		y++
		break
	case SnakeRight:
		x++
		break
	}
	// check the bounds
	if x >= WIDTHBLOCKS {
		x = 0
	}
	if x < 0 {
		x = WIDTHBLOCKS - 1
	}
	if y >= HEIGHTBLOCKS {
		y = 0
	}
	if y < 0 {
		y = HEIGHTBLOCKS - 1
	}

	// draw head
	drawSnakePartial(x, y, green)

	// remove tail
	drawSnakePartial(snake.body[2][0], snake.body[2][1], black)

	// move each segment coords to the next one
	for i := 2; i > 0; i-- {
		snake.body[i][0] = snake.body[i-1][0]
		snake.body[i][1] = snake.body[i-1][1]
	}
	snake.body[0][0] = x
	snake.body[0][1] = y
}

func drawSnake() {
	for i := int16(0); i < 3; i++ {
		drawSnakePartial(snake.body[i][0], snake.body[i][1], green)
	}
}

func drawSnakePartial(x, y int16, c color.RGBA) {
	modY := int16(9)
	if y == 12 {
		modY = 8
	}
	// we changed the size of 10 to 9, so a black border is shown
	// around each segment of the snake
	display.FillRectangle(10*x, 10*y, 9, modY, c)
}

func createApple() {
	appleX = int16(rand.Int31n(WIDTHBLOCKS))
	appleY = int16(rand.Int31n(HEIGHTBLOCKS))
	drawSnakePartial(appleX, appleY, red)
}
