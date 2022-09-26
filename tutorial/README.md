# GoBadge Tutorial

## What you need

    - GoBadge aka Adafruit PyBadge
    - Personal computer with Go 1.18 and TinyGo installed, and a serial port.

## Installation

### Go

If somehow you have not installed Go on your computer already, you can download it here:

https://golang.org/dl/

Now you are ready to install TinyGo.

### TinyGo

Follow the instructions here for your operating system:

https://tinygo.org/getting-started/

Once you have finished installing TinyGo itself, install the drivers needed for the 

## Connecting the GoBadge to your computer

![Adafruit GoBadge](./images/pybadge_hello.jpg)

Plug the GoBadge into your computer using a USB cable. There may be one provided in your starter kit.

## Running the code

The TinyGo programs will run directly on the GoBadge's microcontoller. The procedure is basically:

- Edit your TinyGo program.
- Compile and flash it to your GoBadge.
- The program executes from the GoBadge. You can disconnect the GoBadge from your computer and plug it into a battery if you wish, the program executes directly on the microcontroller.

Let's get started!

## Code

### step0.go - Built-in LED

This tests that you can compile and flash your PyBadge with TinyGo code, by blinking the built-in LED (it's on the back).

![Adafruit PyBadge bootloader](./images/pybadge_bootloader.jpg)

```
tinygo flash -target pybadge ./step0/main.go
```

Once the PyBadge is flashed correctly, the built-in LED labeled "D13" (on the back) should start to turn on and off once per second. Now everything is setup correctly and you are ready to continue.

### step1.go - Built-in LED, START Button

Run the code.

```
tinygo flash -target pybadge ./step1/main.go
```

When you press the START button, the built-in LED should turn on.

### step2.go - Neopixels

Run the code.

```
tinygo flash -target pybadge ./step2/main.go
```

The 5 neopixels should light up green and red alternatively.

### step3.go - Neopixels, Buttons

Run the code.

```
tinygo flash -target pybadge ./step3/main.go
```

The 5 neopixels should light up in different colors depending on which button you press.

### step4.go - Light sensor, Neopixels

Step 4 has been temporarily removed from this workshop.

### step5.go - Display

![PyBadge](./images/pybadge_hello.jpg)

Run the code.

```
tinygo flash -target pybadge ./step5/main.go
```

The message "Hello Gophers!" should appear on the display.

### step6.go - Display, Buttons

![PyBadge](./images/pybadge_display_buttons.jpg)

Run the code.

```
tinygo flash -target pybadge ./step6/main.go
```

The display will show some blue circles. When a button is pressed a ring will be shown around its corresponding circle.

### step7.go - Display, Accelerometer

![PyBadge](./images/pybadge_accel.jpg)

Run the code.

```
tinygo flash -target pybadge ./step7/main.go
```

The display will show a bar for each X,Y,Z axis. Move the Pybadge to see it in action.

### step8.go - Buzzer, Buttons

Run the code.

```
tinygo flash -target pybadge ./step8/main.go
```

Press the buttons and create your melody.

### Snake Game

![PyBadge](./images/pybadge_snake.jpg)

Run the code.

```
tinygo flash -target pybadge ./snake
```

Play the famous Snake game on the pybadge.

### My Name Is

![PyBadge](./images/pybadge_mynameis.jpg)

Run the code.

```
tinygo flash -target pybadge ./mynameis
```

Configure your name and use the awesome TinyGo-powered badge!
