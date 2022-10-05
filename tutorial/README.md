# GoBadge Tutorial

## What you need

    - GoBadge aka Adafruit PyBadge
    - Personal computer with Go 1.18/1.19 and TinyGo installed, and a serial port.

## Installation

### Go

If somehow you have not installed Go on your computer already, you can download it here:

https://go.dev/dl/

Now you are ready to install TinyGo.

### TinyGo

Follow the instructions here for your operating system:

https://tinygo.org/getting-started/


## Putting code on your GoBadge

### Connecting the GoBadge to your computer

![Adafruit GoBadge](./images/pybadge_hello.jpg)

Plug the GoBadge into your computer using a USB cable. There may be one provided in your starter kit.

### Running the code

The TinyGo programs will run directly on the GoBadge's microcontoller. The procedure is basically:

- Edit your TinyGo program.
- Compile and flash it to your GoBadge.
- The program executes from the GoBadge. You can disconnect the GoBadge from your computer and plug it into a battery if you wish, the program executes directly on the microcontroller.

Let's get started!

## The Basics

There is a set of tutorial steps in the "basics" directory here that help show you, well, the basics!

This is a good place to get started with learning how to program your GoBadge.

## Snake Game

Once you have learned the basics, you can go thru our step by step on how to build the "Snake" game.

## My Name Is

The "My Name Is" example uses your GoBadge in a very traditional sense: as a name badge.

## Terminal

The Terminal example turns your GoBadge into a miniature text display that echoes whatever you type on your computer.

## Flightbadge

The Flightbadge example turns your GoBadge into a game-controller style keyboard interface that you can plug into your computer to control programs by sending the correct key commands.
