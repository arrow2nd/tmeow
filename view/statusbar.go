package view

import (
	"fmt"

	"github.com/rivo/tview"
)

// statusBar ステータスバー
type statusBar struct {
	textView *tview.TextView
}

func newStatusBar() *statusBar {
	sb := &statusBar{
		textView: &tview.TextView{},
	}

	sb.textView = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetTextAlign(tview.AlignRight)

	return sb
}

func (sb *statusBar) setStatus(msg, username, screenname string) {
	sb.textView.Clear()
	if msg != "" {
		fmt.Fprintf(sb.textView, "[#9c9c9c]*%s*[white] ", msg)
	}
	fmt.Fprintf(sb.textView, "[#9c9c9c][%s / @%s][white] ", username, screenname)
}
