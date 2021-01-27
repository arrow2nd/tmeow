package main

import (
	"fmt"

	"github.com/rivo/tview"
)

// StatusBar ステータスバー
type StatusBar struct {
	textView *tview.TextView
}

func newStatusBar() *StatusBar {
	sb := &StatusBar{
		textView: &tview.TextView{},
	}

	sb.textView = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetTextAlign(tview.AlignRight)

	return sb
}

func (sb *StatusBar) setStatus(msg, username, screenname string) {
	sb.textView.Clear()
	if msg != "" {
		fmt.Fprintf(sb.textView, "[#9c9c9c]*%s*[white] ", msg)
	}
	fmt.Fprintf(sb.textView, "[#9c9c9c][%s / @%s][white] ", username, screenname)
}
