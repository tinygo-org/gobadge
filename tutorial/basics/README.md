# GoBadge Basic Tutorial

## Connecting the GoBadge to your computer

![Adafruit GoBadge](../images/gobadge.jpg)

Plug the GoBadge into your computer using a USB cable. There may be one provided in your starter kit.

## Running the code

The TinyGo programs will run directly on the GoBadge's microcontoller. The procedure is basically:

- Edit your TinyGo program.
- Compile and flash it to your GoBadge.
- The program executes from the GoBadge. You can disconnect the GoBadge from your computer and plug it into a battery if you wish, the program executes directly on the microcontroller.

Let's get started!

## Code

### step0.go - Built-in LED

![Step 0](../images/basics/step0.jpg)

This tests that you can compile and flash your PyBadge with TinyGo code, by blinking the built-in LED (it's on the back).

```
tinygo flash -target pybadge ./step0/
```

Once the PyBadge is flashed correctly, the built-in LED labeled "D13" (on the back) should start to turn on and off once per second. Now everything is setup correctly and you are ready to continue.

### step1.go - Built-in LED, START Button

![Step 1](../images/basics/step1.jpg)

Run the code.

```
tinygo flash -target pybadge ./step1/
```

When you press the START button, the built-in LED should turn on.

### step2.go - Neopixels

![Step 2](../images/basics/step2.jpg)

Run the code.

```
tinygo flash -target pybadge ./step2/
```

The 5 neopixels should light up green and red alternatively.

### step3.go - Neopixels, Buttons

![Step 3](../images/basics/step3.jpg)

Run the code.

```
tinygo flash -target pybadge ./step3/
```

The 5 neopixels should light up in different colors depending on which button you press.

### step4.go - Light sensor, Neopixels

![Step 4](../images/basics/step4.jpg)

Run the code.

```
tinygo flash -target pybadge ./step4/
```

The 5 neopixels should light up in different colors depending on which button you press.

### step5.go - Display

![Step 5](../images/basics/step5.jpg)

Run the code.

```
tinygo flash -target pybadge ./step5/
```

The message "Hello Gophers!" should appear on the display.

### step6.go - Display, Buttons

![Step 6](../images/basics/step6.jpg)

Run the code.

```
tinygo flash -target pybadge ./step6/
```

The display will show some blue circles. When a button is pressed a ring will be shown around its corresponding circle.

### step7.go - Display, Accelerometer

![Step 7](../images/basics/step7.jpg)

Run the code.

```
tinygo flash -target pybadge ./step7/
```

The display will show a bar for each X,Y,Z axis. Move the Pybadge to see it in action.

### step8.go - Buzzer, Buttons

![Step 8](../images/basics/step8.jpg)

Run the code.

```
tinygo flash -target pybadge ./step8/
```

Press the buttons and create your melody.

