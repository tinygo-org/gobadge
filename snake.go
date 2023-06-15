package main

import (
	"image/color"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"tinygo.org/x/tinyfont/proggy"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const (
	GameSplash = iota
	GameStart
	GamePlay
	GameOver
	GameQuit
)

const (
	SnakeUp = iota
	SnakeDown
	SnakeLeft
	SnakeRight
)

const (
	WIDTHBLOCKS  = 16
	HEIGHTBLOCKS = 13
)

var (
	// Those variable are there for a more easy reading of the apple shape.
	re = colors[RED]   // red
	bk = colors[BLACK] // background
	gr = colors[SNAKE] // green

	// The array is split for a visual purpose too.
	appleBuf = []color.RGBA{
		bk, bk, bk, bk, bk, gr, gr, gr, bk, bk,
		bk, bk, bk, bk, gr, gr, gr, bk, bk, bk,
		bk, bk, bk, re, gr, gr, re, bk, bk, bk,
		bk, bk, re, re, re, re, re, re, bk, bk,
		bk, re, re, re, re, re, re, re, re, bk,
		bk, re, re, re, re, re, re, re, re, bk,
		bk, re, re, re, re, re, re, re, re, bk,
		bk, bk, re, re, re, re, re, re, bk, bk,
		bk, bk, bk, re, re, re, re, bk, bk, bk,
		bk, bk, bk, bk, bk, bk, bk, bk, bk, bk,
	}
)

type Snake struct {
	body      [208][2]int16
	length    int16
	direction int16
}

type SnakeGame struct {
	snake          Snake
	appleX, appleY int16
	status         uint8
	score          int
	frame, delay   int
}

var splashed = false
var scoreStr string

func NewSnakeGame() *SnakeGame {
	return &SnakeGame{
		snake: Snake{
			body: [208][2]int16{
				{0, 3},
				{0, 2},
				{0, 1},
			},
			length:    3,
			direction: SnakeLeft,
		},
		appleX: 5,
		appleY: 5,
		status: GameSplash,
		delay:  120,
	}
}

func (g *SnakeGame) Splash() {
	if !splashed {
		g.splash()
		splashed = true
		time.Sleep(1500 * time.Millisecond)
	}
}

func (g *SnakeGame) Start() {
	display.FillScreen(bk)

	g.initSnake()
	g.drawSnake()
	g.createApple()

	g.status = GamePlay
}

func (g *SnakeGame) Play(direction int) {
	if direction != -1 && ((g.snake.direction == SnakeUp && direction != SnakeDown) ||
		(g.snake.direction == SnakeDown && direction != SnakeUp) ||
		(g.snake.direction == SnakeLeft && direction != SnakeRight) ||
		(g.snake.direction == SnakeRight && direction != SnakeLeft)) {
		g.snake.direction = int16(direction)
	}

	g.moveSnake()
}

func (g *SnakeGame) Over() {
	display.FillScreen(bk)
	splashed = false

	g.status = GameOver
}

func (g *SnakeGame) splash() {
	display.FillScreen(bk)

	logo := `
   ____  _  __  ___    __ __  ____
  / __/ / |/ / / _ |  / //_/ / __/
 _\ \  /    / / __ | / ,<   / _/  
/___/ /_/|_/ /_/ |_|/_/|_| /___/  
                                     
    _____  ___    __  ___  ____
   / ___/ / _ |  /  |/  / / __/
  / (_ / / __ | / /|_/ / / _/  
  \___/ /_/ |_|/_/  /_/ /___/  
`
	for i, line := range strings.Split(strings.TrimSuffix(logo, "\n"), "\n") {
		tinyfont.WriteLine(&display, &tinyfont.Tiny3x3a2pt7b, 12, int16(10+i*5), line+"\n", gr)
	}

	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 30, 120, "Press START", colors[RED])

	if g.score > 0 {
		scoreStr = strconv.Itoa(g.score)
		tinyfont.WriteLineRotated(&display, &proggy.TinySZ8pt7b, 140, 100, "SCORE: "+scoreStr, colors[TEXT], tinyfont.ROTATION_270)
	}
}

func (g *SnakeGame) initSnake() {
	g.snake.body[0][0] = 0
	g.snake.body[0][1] = 3
	g.snake.body[1][0] = 0
	g.snake.body[1][1] = 2
	g.snake.body[2][0] = 0
	g.snake.body[2][1] = 1

	g.snake.length = 3
	g.snake.direction = SnakeRight
}

func (g *SnakeGame) collisionWithSnake(x, y int16) bool {
	for i := int16(0); i < g.snake.length; i++ {
		if x == g.snake.body[i][0] && y == g.snake.body[i][1] {
			return true
		}
	}
	return false
}

func (g *SnakeGame) createApple() {
	g.appleX = int16(rand.Int31n(WIDTHBLOCKS))
	g.appleY = int16(rand.Int31n(HEIGHTBLOCKS))
	for g.collisionWithSnake(g.appleX, g.appleY) {
		g.appleX = int16(rand.Int31n(WIDTHBLOCKS))
		g.appleY = int16(rand.Int31n(HEIGHTBLOCKS))
	}
	g.drawApple(g.appleX, g.appleY)
}

func (g *SnakeGame) moveSnake() {
	x := g.snake.body[0][0]
	y := g.snake.body[0][1]

	switch g.snake.direction {
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

	if g.collisionWithSnake(x, y) {
		g.score = int(g.snake.length - 3)
		g.Over()

		return
	}

	// draw head
	g.drawSnakePartial(x, y, colors[SNAKE])
	if x == g.appleX && y == g.appleY {
		g.snake.length++
		g.createApple()
	} else {
		// remove tail
		g.drawSnakePartial(g.snake.body[g.snake.length-1][0], g.snake.body[g.snake.length-1][1], colors[BLACK])
	}
	for i := g.snake.length - 1; i > 0; i-- {
		g.snake.body[i][0] = g.snake.body[i-1][0]
		g.snake.body[i][1] = g.snake.body[i-1][1]
	}
	g.snake.body[0][0] = x
	g.snake.body[0][1] = y
}

func (g *SnakeGame) drawApple(x, y int16) {
	display.FillRectangleWithBuffer(10*x, 10*y, 10, 10, appleBuf)
}

func (g *SnakeGame) drawSnake() {
	for i := int16(0); i < g.snake.length; i++ {
		g.drawSnakePartial(g.snake.body[i][0], g.snake.body[i][1], colors[SNAKE])
	}
}

func (g *SnakeGame) drawSnakePartial(x, y int16, c color.RGBA) {
	modY := int16(9)
	if y == 12 {
		modY = 8
	}
	display.FillRectangle(10*x, 10*y, 9, modY, c)
}

func (g *SnakeGame) Loop() {
	g.status = GameSplash
	splashed = false
	for {
		g.update()
		if g.status == GameQuit {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (g *SnakeGame) update() {
	buttons.ReadInput()
	switch g.status {
	case GameSplash:
		g.Splash()
		if buttons.Pins[shifter.BUTTON_START].Get() {
			g.Start()
		}
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
			g.status = GameOver
		}
		break

	case GamePlay:
		switch {
		case buttons.Pins[shifter.BUTTON_SELECT].Get():
			g.Over()
			break

		case buttons.Pins[shifter.BUTTON_RIGHT].Get():
			g.Play(SnakeRight)
			break

		case buttons.Pins[shifter.BUTTON_LEFT].Get():
			g.Play(SnakeLeft)
			break

		case buttons.Pins[shifter.BUTTON_DOWN].Get():
			g.Play(SnakeDown)
			break

		case buttons.Pins[shifter.BUTTON_UP].Get():
			g.Play(SnakeUp)
			break

		default:
			g.Play(-1)
			break
		}
		break
	case GameQuit:
	case GameOver:
		g.Splash()

		if buttons.Pins[shifter.BUTTON_START].Get() {
			g.Start()
		}
		if buttons.Pins[shifter.BUTTON_SELECT].Get() {
			g.status = GameQuit
		}
	}
}
