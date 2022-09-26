package main

import (
	"image/color"
	"math/rand"
	"time"

	"../fonts"
	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/tinyfont"
)

const (
	BCK = iota
	SNAKE
	APPLE
	TEXT
)

const (
	START = iota
	PLAY
	GAMEOVER
	QUIT
)

const (
	WIDTHBLOCKS  = 16
	HEIGHTBLOCKS = 13
)

type Snake struct {
	body      [208][2]int16
	length    int16
	direction int16
}

type Game struct {
	colors         []color.RGBA
	snake          Snake
	appleX, appleY int16
	status         uint8
}

func (game *Game) Start() {
	game.status = START
	scoreStr := []byte("SCORE: 123")
	display.FillScreen(game.colors[BCK])
	play := true
	for play {
		switch game.status {
		case START:
			display.FillScreen(game.colors[BCK])

			tinyfont.WriteLine(&display, &fonts.Bold24pt7b, 0, 50, []byte("SNAKE"), game.colors[TEXT])
			tinyfont.WriteLine(&display, &fonts.Regular12pt7b, 8, 100, []byte("Press START"), game.colors[TEXT])

			time.Sleep(2 * time.Second)
			for game.status == START {
				buttons.Read8Input()
				if buttons.Pins[BUTTON_START].Get() {
					game.status = PLAY
				}
			}
			break
		case GAMEOVER:
			display.FillScreen(game.colors[BCK])

			scoreStr[7] = 48 + uint8((game.snake.length-3)/100)
			scoreStr[8] = 48 + uint8(((game.snake.length-3)/10)%10)
			scoreStr[9] = 48 + uint8((game.snake.length-3)%10)

			tinyfont.WriteLine(&display, &fonts.Regular12pt7b, 8, 50, []byte("GAME OVER"), game.colors[TEXT])
			tinyfont.WriteLine(&display, &fonts.Regular12pt7b, 8, 100, []byte("Press START"), game.colors[TEXT])
			tinyfont.WriteLine(&display, &tinyfont.TomThumb, 50, 120, scoreStr, game.colors[TEXT])

			time.Sleep(2 * time.Second)
			for game.status == GAMEOVER {
				buttons.Read8Input()
				if buttons.Pins[BUTTON_START].Get() {
					game.status = PLAY
				}
			}
			break
		case PLAY:
			display.FillScreen(game.colors[BCK])
			game.snake.body = [208][2]int16{
				{0, 3},
				{0, 2},
				{0, 1},
			}
			game.snake.length = 3
			game.snake.direction = 3
			game.drawSnake()
			game.createApple()
			time.Sleep(2000 * time.Millisecond)
			for game.status == PLAY {

				buttons.ReadInput()
				if buttons.Pins[shifter.BUTTON_LEFT].Get() {
					if game.snake.direction != 3 {
						game.snake.direction = 0
					}
				}
				if buttons.Pins[shifter.BUTTON_UP].Get() {
					if game.snake.direction != 2 {
						game.snake.direction = 1
					}
				}
				if buttons.Pins[shifter.BUTTON_DOWN].Get() {
					if game.snake.direction != 1 {
						game.snake.direction = 2
					}
				}
				if buttons.Pins[shifter.BUTTON_RIGHT].Get() {
					if game.snake.direction != 0 {
						game.snake.direction = 3
					}
				}
				if buttons.Pins[shifter.BUTTON_SELECT].Get() {
					game.status = START
				}
				game.moveSnake()
				time.Sleep(100 * time.Millisecond)
			}

			break
		}
	}
}

func (g *Game) collisionWithSnake(x, y int16) bool {
	for i := int16(0); i < g.snake.length; i++ {
		if x == g.snake.body[i][0] && y == g.snake.body[i][1] {
			return true
		}
	}
	return false
}

func (g *Game) createApple() {
	g.appleX = int16(rand.Int31n(16))
	g.appleY = int16(rand.Int31n(13))
	for g.collisionWithSnake(g.appleX, g.appleY) {
		g.appleX = int16(rand.Int31n(16))
		g.appleY = int16(rand.Int31n(13))
	}
	g.drawSnakePartial(g.appleX, g.appleY, g.colors[APPLE])
}

func (g *Game) moveSnake() {
	x := g.snake.body[0][0]
	y := g.snake.body[0][1]

	switch g.snake.direction {
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
		g.status = GAMEOVER
	}

	// draw head
	g.drawSnakePartial(x, y, g.colors[SNAKE])
	if x == g.appleX && y == g.appleY {
		g.snake.length++
		g.createApple()
	} else {
		// remove tail
		g.drawSnakePartial(g.snake.body[g.snake.length-1][0], g.snake.body[g.snake.length-1][1], g.colors[BCK])
	}
	for i := g.snake.length - 1; i > 0; i-- {
		g.snake.body[i][0] = g.snake.body[i-1][0]
		g.snake.body[i][1] = g.snake.body[i-1][1]
	}
	g.snake.body[0][0] = x
	g.snake.body[0][1] = y
}

func (g *Game) drawSnake() {
	for i := int16(0); i < g.snake.length; i++ {
		g.drawSnakePartial(g.snake.body[i][0], g.snake.body[i][1], g.colors[SNAKE])
	}
}

func (g *Game) drawSnakePartial(x, y int16, c color.RGBA) {
	modY := int16(9)
	if y == 12 {
		modY = 8
	}
	display.FillRectangle(10*x, 10*y, 9, modY, c)
}
