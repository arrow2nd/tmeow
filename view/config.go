package view

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/tmeow/config"
	"github.com/rivo/tview"
)

// Config 設定情報
type Config struct {
	App *tview.Application
	API *anaconda.TwitterApi
	Cfg *config.Config
}

// SharedConfig 共通設定
var SharedConfig = Config{
	App: nil,
	API: nil,
	Cfg: nil,
}
