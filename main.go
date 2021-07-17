package main

import (
	_ "embed"
	"github.com/mimose/gcosy/lib"
	"github.com/wailsapp/wails"
	"mo-for-desktop/services/mo"
	"runtime"
)

const (
	programName = "mo-for-desktop"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func basic() string {
	return "Hello World!"
}

func main() {
	// complete program fileName
	var fileName = programName
	if runtime.GOOS == "windows" {
		fileName += ".exe"
	}
	// search the bin
	//moBin,err := lib.Search(fileName, programName)
	//if err != nil {
	//	lib.Abort("Fail to find '%s', error append: %s", fileName, err.Error())
	//}
	//if moBin == nil {
	//	lib.Abort("Fail to find '%s', please make sure it is on your path", fileName)
	//}
	// new a mo-program
	mo, err := mo.New(nil)
	if err != nil {
		lib.Abort(err.Error())
	}
	// wails create
	app := wails.CreateApp(&wails.AppConfig{
		Width:     380,
		Height:    600,
		Title:     "mo",
		JS:        js,
		CSS:       css,
		Colour:    "#fff",
		Resizable: true,
	})
	app.Bind(mo)
	app.Run()
}
