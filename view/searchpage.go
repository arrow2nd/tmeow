package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// searchPage 検索ページ
type searchPage struct {
	*tweets
	frame *tview.Frame
	input *tview.InputField
}

func newSearchPage() *searchPage {
	sp := &searchPage{
		tweets: newtweets(),
		frame:  &tview.Frame{},
		input:  &tview.InputField{},
	}

	sp.input = tview.NewInputField().
		SetLabel("🔍 ").
		SetPlaceholder("キーワード検索（ユーザー検索はIDの先頭に＠をつけてください）").
		SetPlaceholderTextColor(tcell.ColorGrey)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(sp.input, 1, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(sp.textView, 0, 1, true)

	sp.frame = tview.NewFrame(flex).
		SetBorders(0, 0, 0, 0, 1, 1)

	return sp
}

func (sp *searchPage) init() {
	sp.tweetsDraw()
}
