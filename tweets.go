package main

import (
	"fmt"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/rivo/tview"
)

// TweetsView ツイート
type TweetsView struct {
	textView *tview.TextView
	tweets   *[]anaconda.Tweet
}

func newTweetsView() *TweetsView {
	tw := new(TweetsView)

	tw.textView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetRegions(true).
		SetWrap(false)

	tw.textView.SetHighlightedFunc(func(added []string, removed []string, remaining []string) {
		tw.textView.ScrollToHighlight()
	})

	return tw
}

func (tw *TweetsView) tweetsDraw() {
	for i := 0; i < 15; i++ {
		text := fmt.Sprintf("[#ffffff]ユーザー名 [#9c9c9c](@screen_name)[#ce99de] 2021/01/01 00:00:00 [#e887b9]1fav\n")
		text += fmt.Sprintf("[default]@hogehoge ツイート文ここ (%d)", i)
		rg := fmt.Sprintf(`[white]["tweet_%d"] [""] `, i)
		fmt.Fprintf(tw.textView, "%s%s\n\n", rg, strings.Replace(text, "\n", "\n"+rg, -1))
	}
	tw.textView.Highlight("tweet_0")
}

func (tw *TweetsView) cursorUp() {
	idx := getHighlightID(tw.textView)
	if idx == -1 || tw.tweets == nil {
		return
	}

	if idx--; idx < 0 {
		idx = len(*tw.tweets) - 1
	}
	tw.textView.Highlight(fmt.Sprintf("tweet_%d", idx))
}

func (tw *TweetsView) cursorDown() {
	idx := getHighlightID(tw.textView)
	if idx == -1 || tw.tweets == nil {
		return
	}

	idx = (idx + 1) % len(*tw.tweets)
	tw.textView.Highlight(fmt.Sprintf("tweet_%d", idx))
}
