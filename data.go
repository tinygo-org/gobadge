package main

// Replace with your data by using -ldflags like this:
//
// tinygo flash -target pybadge -ldflags="-X main.YourName=@myontwitter -X main.YourTitle1='Amazing human' -X main.YourTitle2='also kind'"
//
// See Makefile for more info.
var (
	YourName, YourTitleA1, YourTitleA2, YourTitleB1, YourTitleB2     string
	YourMarqueeTop, YourMarqueeMiddle, YourMarqueeBottom, YourQRText string
)

const (
	DefaultName          = "@TinyGolang"
	DefaultTitleA1       = "Go Compiler"
	DefaultTitleA2       = "Small Places"
	DefaultMarqueeTop    = "This badge"
	DefaultMarqueeMiddle = "runs"
	DefaultMarqueeBottom = "TINYGO"
	DefaultQRText        = "https://tinygo.org"
	DefaultTitleB1       = "I enjoy"
	DefaultTitleB2       = "TINYGO"
)
