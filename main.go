package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1000,
		Height:    800,
		Title:     "mo",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
		Resizable: true,
	})
	app.Bind(basic)
	app.Run()
}
