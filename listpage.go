package main

import (
	"github.com/rivo/tview"
)

// ListPage リストタイムライン
type ListPage struct {
	*TweetsView
	frame *tview.Frame
	drop  *tview.DropDown
}

func newListPage() *ListPage {
	lp := &ListPage{
		TweetsView: newTweetsView(),
		frame:      &tview.Frame{},
		drop:       &tview.DropDown{},
	}

	lp.drop = tview.NewDropDown().SetLabel("📑  ")

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(lp.drop, 1, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(lp.textView, 0, 1, true)

	lp.frame = tview.NewFrame(flex).SetBorders(0, 0, 0, 0, 1, 1)

	return lp
}

func (lp *ListPage) init() {
	lp.tweetsDraw()
	lp.setListName([]string{"LIST 1", "LIST 2"})
}

func (lp *ListPage) setListName(listname []string) {
	lp.drop.SetOptions(listname, nil).SetCurrentOption(0)
}
