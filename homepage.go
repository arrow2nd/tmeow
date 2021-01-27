package main

import (
	"github.com/rivo/tview"
)

// HomePage ホームタイムライン
type HomePage struct {
	*TweetsView
	frame *tview.Frame
}

func newHomePage() *HomePage {
	home := HomePage{
		TweetsView: newTweetsView(),
		frame:      &tview.Frame{},
	}

	home.frame = tview.NewFrame(home.textView).
		SetBorders(0, 0, 0, 0, 1, 1)

	return &home
}

func (hp *HomePage) init() {
	hp.tweetsDraw()
}
