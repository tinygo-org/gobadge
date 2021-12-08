# GoBadge

TinyGo powered badge using the Adafruit Pybadge hardware.

https://www.adafruit.com/product/4200

# How to install

- Install TinyGo using the instructions from https://tinygo.org

- Clone this repo

- Change directories into the directory with the repo

- Customize the file `data.go` with your own name and information

- Connect your Pybadge to your computer using a USB cable

- Run this command to compile and flash the code to your Pybadge:

```
tinygo flash -target pybadge .
```

Congratulations! It is now a GoBadge.
