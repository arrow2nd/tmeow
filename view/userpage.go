package view

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// userPage ユーザータイムライン
type userPage struct {
	*tweets
	frame          *tview.Frame
	userInfo       *tview.TextView
	tweetsCount    *tview.TextView
	followingCount *tview.TextView
	followersCount *tview.TextView
}

func newUserPage() *userPage {
	up := &userPage{
		tweets:         newtweets(),
		frame:          &tview.Frame{},
		userInfo:       &tview.TextView{},
		tweetsCount:    &tview.TextView{},
		followingCount: &tview.TextView{},
		followersCount: &tview.TextView{},
	}

	up.userInfo = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetRegions(true).
		SetWrap(false)

	up.tweetsCount = up.createCountsView(0xe06c75)
	up.followingCount = up.createCountsView(0xc678dd)
	up.followersCount = up.createCountsView(0x56b6c2)

	userCountsView := tview.NewFlex().
		AddItem(up.tweetsCount, 0, 1, false).
		AddItem(up.followingCount, 0, 1, false).
		AddItem(up.followersCount, 0, 1, false)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(up.userInfo, 4, 1, false).
		AddItem(userCountsView, 1, 1, false).
		AddItem(nil, 1, 1, false).
		AddItem(up.textView, 0, 1, true)

	up.frame = tview.NewFrame(flex).
		SetBorders(0, 0, 0, 0, 1, 1)

	return up
}

func (up *userPage) init() {
	up.tweetsDraw()
	up.setUserInfo("ユーザー名", "screen_name", "[blue]フォローされています", "BIO", "place")
	up.setUserCounts("10000", "10000", "10000")
}

func (up *userPage) createCountsView(color int32) *tview.TextView {
	txtView := tview.NewTextView().
		SetDynamicColors(true).
		SetTextColor(tcell.ColorBlack).
		SetTextAlign(tview.AlignCenter)
	txtView.SetBackgroundColor(tcell.NewHexColor(color))
	return txtView
}

func (up *userPage) setUserInfo(username, screenname, relation, bio, place string) {
	up.userInfo.Clear()
	fmt.Fprintf(up.userInfo, "[#ffffff]%s[white]\n", username)
	fmt.Fprintf(up.userInfo, "[#9c9c9c]@%s %s[white]\n", screenname, relation)
	fmt.Fprintf(up.userInfo, "[#9c9c9c]📄 : %s[white]\n", bio)
	fmt.Fprintf(up.userInfo, "[#9c9c9c]📍 : %s[white]", place)
}

func (up *userPage) setUserCounts(tweets, following, followers string) {
	up.tweetsCount.SetText(fmt.Sprintf("%s Tweets", tweets))
	up.followingCount.SetText(fmt.Sprintf("%s Following", following))
	up.followersCount.SetText(fmt.Sprintf("%s Followers", followers))
}
