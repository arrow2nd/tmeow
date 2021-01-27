package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// View 画面
type View struct {
	pages       *tview.Pages
	homePage    *homePage
	mentionPage *mentionPage
	listPage    *listPage
	searchPage  *searchPage
	userPage    *userPage
	tabBar      *tabBar
	statusBar   *statusBar
	inputBar    *inputBar
}

// NewView 画面を作成
func NewView() *View {
	// 配色設定
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault
	tview.Styles.ContrastBackgroundColor = tcell.ColorDefault

	view := &View{
		pages:       tview.NewPages(),
		homePage:    newHomePage(),
		mentionPage: newMentionPage(),
		listPage:    newListPage(),
		searchPage:  newSearchPage(),
		userPage:    newUserPage(),
		tabBar:      &tabBar{},
		statusBar:   newStatusBar(),
		inputBar:    newInputBar(),
	}

	// ページ
	view.pages.AddPage("page_0", view.homePage.frame, true, true).
		AddPage("page_1", view.mentionPage.frame, true, false).
		AddPage("page_2", view.listPage.frame, true, false).
		AddPage("page_3", view.searchPage.frame, true, false).
		AddPage("page_4", view.userPage.frame, true, false)

	// タブバー
	view.tabBar = newTabBar(view.pages)
	view.tabBar.setTab([]string{"HOME", "MENTION", "LIST", "SEARCH", "USER"})

	// レイアウト
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(view.tabBar.textView, 2, 1, false).
		AddItem(view.pages, 0, 1, true).
		AddItem(view.statusBar.textView, 1, 1, false).
		AddItem(view.inputBar.inputField, 1, 1, false)

	// ルートプリミティブに設定
	SharedConfig.App.SetRoot(layout, true)

	// オプション設定
	SharedConfig.App.EnableMouse(true).
		SetBeforeDrawFunc(func(screen tcell.Screen) bool {
			screen.Clear()
			return false
		})

	return view
}

// Init 初期化
func (view *View) Init() {
	// キーイベントハンドラ設定
	view.setInputCapture()

	// 各ページ初期化
	view.homePage.init()
	view.mentionPage.init()
	view.listPage.init()
	view.searchPage.init()
	view.userPage.init()
}

func (view *View) setInputCapture() {
	SharedConfig.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyLeft:
			// 左のタブへ
			view.tabBar.cursorLeft(view.pages)
			return nil

		case tcell.KeyRight:
			// 右のタブへ
			view.tabBar.cursorRight(view.pages)
			return nil

		case tcell.KeyCtrlI:
			// フォーカスを入力欄へ
			SharedConfig.App.SetFocus(view.inputBar.inputField)
			return nil

		case tcell.KeyEscape:
			// フォーカスをページへ
			SharedConfig.App.SetFocus(view.pages)
			return nil
		}
		return event
	})
}
