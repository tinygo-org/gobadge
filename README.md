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

# Add an image

- Create an image with a 160x128 pixels size, copy it into `cmd/assets` folder.  
For the moment only jpeg images are supported.  
- In `cmd/main.go` replace the input with the path of your file:
```go
logos.GenerateLogoRGBAFile("cmd/assets/your-file.jpeg")
```

You can run:
```bash
make flash-gcuk
```

It will generate for you the image into a `[]color.RGBA` and store it in a variable.


👏 Congratulations! It is now a GoBadge.
