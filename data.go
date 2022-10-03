package main

// Replace with your data by using -ldflags like this:
//
// tinygo flash -target pybadge -ldflags="-X main.YourName=@myontwitter -X main.YourTitle1='Amazing human' -X main.YourTitle2='also kind'"
//
// See Makefile for more info.
//
var (
	YourName, YourTitle1, YourTitle2 string
)

const (
	DefaultName   = "@TinyGolang"
	DefaultTitle1 = "Go Compiler"
	DefaultTitle2 = "Small Places"
)
