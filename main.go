package main

import (
	"github.com/arrow2nd/tmeow/api"
	"github.com/arrow2nd/tmeow/config"
	"github.com/arrow2nd/tmeow/view"
	"github.com/rivo/tview"
)

func main() {
	// API・設定構造体作成
	api := api.NewTwitterAPI()
	cfg := config.NewConfig()

	// 設定読込
	if !cfg.Load() {
		cfg.Cred.Token, cfg.Cred.Secret = api.Auth()
		cfg.Save()
	} else {
		api.Init(cfg.Cred.Token, cfg.Cred.Secret)
	}

	// アプリケーション作成
	app := tview.NewApplication()

	// 画面作成
	view.SetSharedConfig(app, api, cfg)
	view := view.NewView()
	view.Init()

	// 実行
	if err := app.Run(); err != nil {
		panic(err)
	}
}
