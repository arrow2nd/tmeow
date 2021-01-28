package view

import (
	"fmt"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/rivo/tview"
)

// tweets ツイート
type tweets struct {
	textView *tview.TextView
	contents *[]anaconda.Tweet
}

func newtweets() *tweets {
	tw := &tweets{
		textView: tview.NewTextView(),
		contents: &[]anaconda.Tweet{},
	}

	tw.textView.SetDynamicColors(true).
		SetScrollable(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added []string, removed []string, remaining []string) {
			tw.textView.ScrollToHighlight()
		})

	return tw
}

func (tw *tweets) draw() {
	for i := 0; i < 15; i++ {
		text := fmt.Sprintf("[#ffffff]ユーザー名 [#9c9c9c](@screen_name)[#ce99de] 2021/01/01 00:00:00 [#e887b9]1fav\n")
		text += fmt.Sprintf("[default]@hogehoge ツイート文ここ (%d)", i)
		rg := fmt.Sprintf(`[white]["tweet_%d"] [""] `, i)
		fmt.Fprintf(tw.textView, "%s%s\n\n", rg, strings.Replace(text, "\n", "\n"+rg, -1))
	}
	tw.textView.Highlight("tweet_0")
}

func (tw *tweets) cursorUp() {
	idx := tw.getHighlightIndex()
	if idx == -1 {
		return
	}
	if idx--; idx < 0 {
		idx = len(*tw.contents) - 1
	}
	tw.textView.Highlight(fmt.Sprintf("tweet_%d", idx))
}

func (tw *tweets) cursorDown() {
	idx := tw.getHighlightIndex()
	if idx == -1 {
		return
	}
	idx = (idx + 1) % len(*tw.contents)
	tw.textView.Highlight(fmt.Sprintf("tweet_%d", idx))
}

func (tw *tweets) getHighlightIndex() int {
	if tw.contents == nil || len(*tw.contents) <= 0 {
		return -1
	}
	return getHighlightID(tw.textView)
}
