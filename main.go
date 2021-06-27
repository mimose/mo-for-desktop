package main

import (
	_ "embed"
	"mo-for-desktop/services/group"
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
	//mo, err := mo.New(nil)
	//if err != nil {
	//	lib.Abort(err.Error())
	//}
	// wails create
	//app := wails.CreateApp(&wails.AppConfig{
	//	Width:     1050,
	//	Height:    560,
	//	Title:     "mo",
	//	JS:        js,
	//	CSS:       css,
	//	Colour:    "#131313",
	//	Resizable: true,
	//})
	//app.Bind(mo)
	//app.Run()

	group.ListAll()
	//group.AddOne("测试")
}
