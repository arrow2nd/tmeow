package main

import (
	"github.com/arrow2nd/tmeow/view"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	view.SharedConfig.App = app

	view := view.NewView()
	view.Init()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
