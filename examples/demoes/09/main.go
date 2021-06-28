package main

import (
	"log"

	"github.com/pkaftj/sciter-go"
	"github.com/pkaftj/sciter-go/window"
)

func main() {
	// Creating application window with default flags and rect size
	win, err := window.New(sciter.DefaultWindowCreateFlag, sciter.DefaultRect)
	if err != nil {
		log.Fatal(err)
	}

	// AddQuitMenu function adds Quit menu (in menubar) and support for
	// CMD+q to terminate / close application on MacOS
	//
	// Note : AddQuitMenu() only works for MacOS on other platform is
	//        does not have any effect.
	win.AddQuitMenu()

	win.Show()
	win.Run()
}
