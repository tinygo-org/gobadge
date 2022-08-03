package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("logo.png")
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	contents := "package main\n\nimport \"image/color\"\n\nvar logoBuffer = []color.RGBA{\n"
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()
			contents += "{" + strconv.Itoa(int(r>>8)) + ", " + strconv.Itoa(int(g>>8)) + ", " + strconv.Itoa(int(b>>8)) + ", 255}"
			if y != img.Bounds().Max.Y-1 || x != img.Bounds().Max.X-1 {
				contents += ", "
			}

		}
	}
	contents += "\n}"
	fmt.Println(contents)
}
