package main

import (
	"fmt"

	"github.com/rivo/tview"
)

// TabBar タブバー
type TabBar struct {
	textView *tview.TextView
}

func newTabBar(pg *tview.Pages) *TabBar {
	tb := &TabBar{
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

func (tb *TabBar) setTab(tabname []string) {
	for i, name := range tabname {
		fmt.Fprintf(tb.textView, `["page_%d"][darkcyan] %s [white][""]`, i, name)
	}

	tb.textView.Highlight("page_0")
}

func (tb *TabBar) cursorLeft(pg *tview.Pages) {
	idx := getHighlightID(tb.textView)
	pageCount := pg.GetPageCount()
	if idx--; idx < 0 {
		idx = pageCount - 1
	}
	tb.textView.Highlight(fmt.Sprintf("page_%d", idx))
}

func (tb *TabBar) cursorRight(pg *tview.Pages) {
	idx := getHighlightID(tb.textView)
	pageCount := pg.GetPageCount()
	idx = (idx + 1) % pageCount
	tb.textView.Highlight(fmt.Sprintf("page_%d", idx))
}
