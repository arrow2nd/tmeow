package view

import (
	"github.com/arrow2nd/tmeow/api"
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
	hp.load()
	hp.tweets.draw()
}

func (hp *homePage) load() {
	v := api.CreateURLValues(sc.Cfg.Option.Counts)
	tweets, err := sc.Twitter.GetTimeline("home", v)
	if err != nil {
		return
	}
	hp.tweets.add(tweets)
}
