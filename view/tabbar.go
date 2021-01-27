package view

import (
	"fmt"

	"github.com/rivo/tview"
)

// tabBar タブバー
type tabBar struct {
	textView *tview.TextView
}

func newTabBar(pg *tview.Pages) *tabBar {
	tb := &tabBar{
		textView: &tview.TextView{},
	}

	tb.textView = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetTextAlign(tview.AlignLeft).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pg.SwitchToPage(added[0])
		})

	return tb
}

func (tb *tabBar) setTab(tabname []string) {
	for i, name := range tabname {
		fmt.Fprintf(tb.textView, `["page_%d"][darkcyan] %s [white][""]`, i, name)
	}

	tb.textView.Highlight("page_0")
}

func (tb *tabBar) cursorLeft(pg *tview.Pages) {
	idx := getHighlightID(tb.textView)
	pageCount := pg.GetPageCount()
	if idx--; idx < 0 {
		idx = pageCount - 1
	}
	tb.textView.Highlight(fmt.Sprintf("page_%d", idx))
}

func (tb *tabBar) cursorRight(pg *tview.Pages) {
	idx := getHighlightID(tb.textView)
	pageCount := pg.GetPageCount()
	idx = (idx + 1) % pageCount
	tb.textView.Highlight(fmt.Sprintf("page_%d", idx))
}
