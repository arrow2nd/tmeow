package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// inputBar 入力バー
type inputBar struct {
	inputField *tview.InputField
}

func newInputBar() *inputBar {
	ib := &inputBar{
		inputField: &tview.InputField{},
	}

	ib.inputField = tview.NewInputField().
		SetPlaceholder("いまどうしてる？").
		SetFieldTextColor(tcell.ColorDefault)

	return ib
}
