package view

import (
	"github.com/arrow2nd/tmeow/api"
	"github.com/arrow2nd/tmeow/config"
	"github.com/rivo/tview"
)

// Config 設定情報
type Config struct {
	App     *tview.Application
	Twitter *api.TwitterAPI
	Cfg     *config.Config
}

var sc = Config{
	App:     nil,
	Twitter: nil,
	Cfg:     nil,
}

// SetSharedConfig 共有設定登録
func SetSharedConfig(app *tview.Application, api *api.TwitterAPI, cfg *config.Config) {
	sc.App = app
	sc.Twitter = api
	sc.Cfg = cfg
}
