# GoBadge Snake Tutorial

In this tutorial we're going to build a _snake-like_ game from scratch.

It is best to have completed the "basics" tutorial before this one.

Let's get started!

## Code

### step0.go - Get a pixel on the screen

This step is to get familiar with the display and the drawing functions.

![Step 0](../images/snake/step0.jpg)

```
tinygo flash -target pybadge ./step0/
```

Once the PyBadge is flashed correctly, a green pixel will appear on the middle of the screen. Feel free to change the values of displaySetPixel and see where the pixel appears!

### step1.go - A pixel, but bigger!

Run the code.

```
tinygo flash -target pybadge ./step1/
```

Instead of a pixel, we are drawing a 10x10 green rectangle.

### step2.go - I like to move it, move it

Run the code.

```
tinygo flash -target pybadge ./step2/
```

We listen to the input buttons and move our rectangle across the display.

### step3.go - Run, snake, run

Run the code.

```
tinygo flash -target pybadge ./step3/
```

Have you noticed the snake at the previous step was kind of slow? That was because display.FillScreen draws the whole display and is a slow process, we could improve the speed if only re-draw the pixels that has 
changed. 

### step4.go - Welcome to the grid

Run the code.

```
tinygo flash -target pybadge ./step4/
```

In the previous step, the 10x10 _snake_ was moving by 1 pixel each time, we need to divide our display (160x128) in a 10x10 grid so the snake will move a whole block each time.

### step5.go - Long snake is lon

![Step 5](../images/snake/step5.jpg)

Run the code.

```
tinygo flash -target pybadge ./step5/
```

Our little snake grew from 1 block to 3 block length.

### step6.go - Time to grow up

![Step 6](../images/snake/step6.jpg)

Feed our snake some red apples so it can grow

```
tinygo flash -target pybadge ./step6/
```

### step7.go - Score & game mechanics

![Step 7](../images/snake/step7.jpg)

Run the code.

```
tinygo flash -target pybadge ./step7/
```

Add game mechanics such as collision (game over) and score.

### step8.go - Get wild!

There's no step 8, it's time to get creative and modify the game as you wish, try adding sounds or different colors.
