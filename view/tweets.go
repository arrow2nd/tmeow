package view

import (
	"fmt"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
)

// tweets ツイート
type tweets struct {
	textView *tview.TextView
	contents *[]anaconda.Tweet
	count    int
}

func newtweets() *tweets {
	tw := &tweets{
		textView: tview.NewTextView(),
		contents: &[]anaconda.Tweet{},
		count:    0,
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
	width := getWindowWidth()
	for i, tweet := range *tw.contents {
		tw.drawTweet(i, &tweet, width)
	}
	tw.textView.Highlight("tweet_0")
}

func (tw *tweets) drawTweet(i int, t *anaconda.Tweet, w int) {
	text := ""

	// リツイート元のツイートに置換
	if t.RetweetedStatus != nil {
		text += fmt.Sprintf("[#%x]RT by @%s\n", sc.Cfg.Color.Retweet, t.User.ScreenName)
		t = t.RetweetedStatus
	}

	// リプライ先表示
	if t.InReplyToScreenName != "" {
		text += fmt.Sprintf("[#%x]Reply to @%s\n", sc.Cfg.Color.Reply, t.InReplyToScreenName)
	}

	// ヘッダー
	userInfo := createUserString(&t.User)
	pt, _ := t.CreatedAtTime()
	postTime := createTimeString(pt)
	reaction := ""
	createReaction(&reaction, t.FavoriteCount, t.Favorited, "Fav", sc.Cfg.Color.Favorite)
	createReaction(&reaction, t.RetweetCount, t.Retweeted, "RT", sc.Cfg.Color.Retweet)
	text += fmt.Sprintf("%s %s%s\n", userInfo, postTime, reaction)

	// テキスト
	text += runewidth.Wrap(removeRuneAmbiguousWidth(t.FullText), w)

	// カーソル
	rg := fmt.Sprintf(`[white]["tweet_%d"] [""][default] `, i)
	fmt.Fprintf(tw.textView, "%s%s\n\n", rg, strings.Replace(text, "\n", "\n"+rg, -1))
}

func (tw *tweets) add(addContents *[]anaconda.Tweet) {
	// end := 200 - len(*addContents)
	// if end <= 0 {
	// 	tw.contents = addContents
	// 	return
	// }
	// tmp := append(*addContents, (*tw.contents)[:end]...)
	tmp := append(*addContents, *tw.contents...)
	tw.count = len(tmp)
	tw.contents = &tmp
}

func (tw *tweets) register(newContents *[]anaconda.Tweet) {
	tw.count = len(*newContents)
	tw.contents = newContents
}

func (tw *tweets) cursorUp() {
	idx := tw.getHighlightIndex()
	if idx == -1 {
		return
	}
	if idx--; idx < 0 {
		idx = tw.count - 1
	}
	tw.textView.Highlight(fmt.Sprintf("tweet_%d", idx))
}

func (tw *tweets) cursorDown() {
	idx := tw.getHighlightIndex()
	if idx == -1 {
		return
	}
	idx = (idx + 1) % tw.count
	tw.textView.Highlight(fmt.Sprintf("tweet_%d", idx))
}

func (tw *tweets) getHighlightIndex() int {
	if tw.contents == nil || tw.count <= 0 {
		return -1
	}
	return getHighlightID(tw.textView)
}
