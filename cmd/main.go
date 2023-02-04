package main

import (
	"flag"
	"fmt"

	"github.com/tinygo-org/gobadge/cmd/logos"
)

const (
	gopherconEU22Logo = "./cmd/assets/gopherconeu-2022.jpg"
	gopherconUK22Logo = "./cmd/assets/gopherconuk-2022.jpg"
	gopherconUS22Logo = "./cmd/assets/gopherconus-2022.jpg"
	fosdem23Logo      = "./cmd/assets/fosdem-2023.jpg"
	tinygoLogo        = "./cmd/assets/tinygo.jpg"
)

func main() {
	conf := flag.String("conf", tinygoLogo, "Choose the conference logo you want to (e.g.: tinygo, gceu22, gcuk22, gcus22)")
	flag.Parse()

	c := confs()
	logo, ok := c[*conf]
	if !ok {
		fmt.Println("I do not have yet this conf in my catalog.")
		return
	}

	logos.GenerateLogoRGBAFile(logo)
}

func confs() map[string]string {
	return map[string]string{
		"gceu22":   gopherconEU22Logo,
		"gcuk22":   gopherconUK22Logo,
		"gcus22":   gopherconUS22Logo,
		"fosdem23": fosdem23Logo,
		"tinygo":   tinygoLogo,
	}
}
