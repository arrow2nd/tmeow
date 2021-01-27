package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// InputBar 入力バー
type InputBar struct {
	inputField *tview.InputField
}

func newInputBar() *InputBar {
	ib := &InputBar{
		inputField: &tview.InputField{},
	}

	ib.inputField = tview.NewInputField().
		SetPlaceholder("いまどうしてる？").
		SetFieldTextColor(tcell.ColorDefault)

	return ib
}
