package main

import "github.com/rivo/tview"

// MentionPage メンションタイムライン
type MentionPage struct {
	*TweetsView
	frame *tview.Frame
}

func newMentionPage() *MentionPage {
	mp := &MentionPage{
		TweetsView: newTweetsView(),
		frame:  &tview.Frame{},
	}

	mp.frame = tview.NewFrame(mp.textView).
		SetBorders(0, 0, 0, 0, 1, 1)

	return mp
}

func (mp *MentionPage) init() {
	mp.tweetsDraw()
}
