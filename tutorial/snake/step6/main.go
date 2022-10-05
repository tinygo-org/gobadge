package main

import (
	"image/color"
	"machine"
	"math/rand"
	"time"

	"tinygo.org/x/drivers/st7735"
)

const (
	WIDTHBLOCKS  = 16
	HEIGHTBLOCKS = 13
)

var display st7735.Device

type Snake struct {
	body      [208][2]int16
	length    int16
	direction int16
}

var snake = Snake{
	body: [208][2]int16{
		{0, 3},
		{0, 2},
		{0, 1},
	},
	length:    3,
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
	createApple()
	for {
		buttons.ReadInput()
		switch {
		// add some checks so the snake doesn't go backwards
		case buttons.Pins[shifter.BUTTON_LEFT].Get():
			if snake.direction != 3 {
				snake.direction = 0
			}
		case buttons.Pins[shifter.BUTTON_UP].Get():
			if snake.direction != 2 {
				snake.direction = 1
			}
		case buttons.Pins[shifter.BUTTON_DOWN].Get():
			if snake.direction != 1 {
				snake.direction = 2
			}
		case buttons.Pins[shifter.BUTTON_RIGHT].Get():
			if snake.direction != 0 {
				snake.direction = 3
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
	case 0:
		x--
		break
	case 1:
		y--
		break
	case 2:
		y++
		break
	case 3:
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

	if x == appleX && y == appleY {
		// grow our snake if we eat the apple
		snake.length++
		// create a new apple
		createApple()
	} else {
		// remove tail in case we do not eat the apple
		drawSnakePartial(snake.body[snake.length-1][0], snake.body[snake.length-1][1], black)
	}
	for i := snake.length - 1; i > 0; i-- {
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
