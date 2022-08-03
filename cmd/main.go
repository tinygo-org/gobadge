package main

import (
	"github.com/tinygo-org/gobadge/cmd/logos"
)

const (
	gopherconUK22Logo = "./cmd/assets/gopherconuk-2022.jpg"
)

func main() {
	logos.GenerateLogoRGBAFile(gopherconUK22Logo)
}
