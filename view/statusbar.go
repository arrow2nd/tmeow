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

func (sb *statusBar) init() {
	sb.setStatus("25件のツイートを読み込みました")
}

func (sb *statusBar) setStatus(msg string) {
	sb.textView.Clear()
	fmt.Fprintf(sb.textView, "[grey]%s[white] [#9c9c9c][ @%s ][white] ", msg, sc.Twitter.OwnUser.ScreenName)
}
