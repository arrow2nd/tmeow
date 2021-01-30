package view

import (
	"github.com/rivo/tview"
)

// listPage リストタイムライン
type listPage struct {
	tweets *tweets
	frame  *tview.Frame
	drop   *tview.DropDown
}

func newListPage() *listPage {
	lp := &listPage{
		tweets: newtweets(),
		frame:  &tview.Frame{},
		drop:   &tview.DropDown{},
	}

	lp.drop = tview.NewDropDown().SetLabel("📑  ")

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(lp.drop, 1, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(lp.tweets.textView, 0, 1, true)

	lp.frame = tview.NewFrame(flex).SetBorders(0, 0, 0, 0, 1, 1)

	return lp
}

func (lp *listPage) init() {
	lp.tweets.draw()
	lp.setListName(sc.Twitter.ListNames)
}

func (lp *listPage) setListName(listname []string) {
	lp.drop.SetOptions(listname, nil).SetCurrentOption(0)
}
