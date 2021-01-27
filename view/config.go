package view

import "github.com/rivo/tview"

// Config 設定情報
type Config struct {
	App *tview.Application
}

// SharedConfig 共通設定
var SharedConfig = Config{
	App: nil,
}
