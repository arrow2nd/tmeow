package view

import "github.com/rivo/tview"

// mentionPage メンションタイムライン
type mentionPage struct {
	*tweets
	frame *tview.Frame
}

func newMentionPage() *mentionPage {
	mp := &mentionPage{
		tweets: newtweets(),
		frame:  &tview.Frame{},
	}

	mp.frame = tview.NewFrame(mp.textView).
		SetBorders(0, 0, 0, 0, 1, 1)

	return mp
}

func (mp *mentionPage) init() {
	mp.tweetsDraw()
}
