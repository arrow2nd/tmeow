package view

import (
	"github.com/rivo/tview"
)

// homePage ホームタイムライン
type homePage struct {
	*tweets
	frame *tview.Frame
}

func newHomePage() *homePage {
	home := homePage{
		tweets: newtweets(),
		frame:  &tview.Frame{},
	}

	home.frame = tview.NewFrame(home.textView).
		SetBorders(0, 0, 0, 0, 1, 1)

	return &home
}

func (hp *homePage) init() {
	hp.tweetsDraw()
}
