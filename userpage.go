package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// UserPage ユーザータイムライン
type UserPage struct {
	*TweetsView
	frame         *tview.Frame
	infoView      *tview.TextView
	tweetsView    *tview.TextView
	followingView *tview.TextView
	followersView *tview.TextView
}

func newUserPage() *UserPage {
	up := &UserPage{
		TweetsView:    newTweetsView(),
		frame:         &tview.Frame{},
		infoView:      &tview.TextView{},
		tweetsView:    &tview.TextView{},
		followingView: &tview.TextView{},
		followersView: &tview.TextView{},
	}

	up.infoView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetRegions(true).
		SetWrap(false)

	up.tweetsView = up.createCountsView(0xe06c75)
	up.followingView = up.createCountsView(0xc678dd)
	up.followersView = up.createCountsView(0x56b6c2)

	userCountsView := tview.NewFlex().
		AddItem(up.tweetsView, 0, 1, false).
		AddItem(up.followingView, 0, 1, false).
		AddItem(up.followersView, 0, 1, false)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(up.infoView, 4, 1, false).
		AddItem(userCountsView, 1, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(up.textView, 0, 1, true)

	up.frame = tview.NewFrame(flex).
		SetBorders(0, 0, 0, 0, 1, 1)

	return up
}

func (up *UserPage) init() {
	up.tweetsDraw()
	up.setUserInfo("ユーザー名", "screen_name", "[blue]フォローされています", "BIO", "place")
	up.setUserCounts("10000", "10000", "10000")
}

func (up *UserPage) createCountsView(color int32) *tview.TextView {
	txtView := tview.NewTextView().
		SetDynamicColors(true).
		SetTextColor(tcell.ColorBlack).
		SetTextAlign(tview.AlignCenter)
	txtView.SetBackgroundColor(tcell.NewHexColor(color))
	return txtView
}

func (up *UserPage) setUserInfo(username, screenname, relation, bio, place string) {
	up.infoView.Clear()
	fmt.Fprintf(up.infoView, "[#ffffff]%s[white]\n", username)
	fmt.Fprintf(up.infoView, "[#9c9c9c]@%s %s[white]\n", screenname, relation)
	fmt.Fprintf(up.infoView, "[#9c9c9c]📄 : %s[white]\n", bio)
	fmt.Fprintf(up.infoView, "[#9c9c9c]📍 : %s[white]", place)
}

func (up *UserPage) setUserCounts(tweets, following, followers string) {
	up.tweetsView.SetText(fmt.Sprintf("%s Tweets", tweets))
	up.followingView.SetText(fmt.Sprintf("%s Following", following))
	up.followersView.SetText(fmt.Sprintf("%s Followers", followers))
}
