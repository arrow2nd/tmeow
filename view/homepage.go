package view

import (
	"github.com/rivo/tview"
)

// homePage ホームタイムライン
type homePage struct {
	tweets *tweets
	frame  *tview.Frame
}

func newHomePage() *homePage {
	home := homePage{
		tweets: newtweets(),
		frame:  &tview.Frame{},
	}

	home.frame = tview.NewFrame(home.tweets.textView).
		SetBorders(0, 0, 0, 0, 1, 1)

	return &home
}

func (hp *homePage) init() {
	hp.tweets.draw()
}
