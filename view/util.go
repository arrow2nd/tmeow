package view

import (
	"fmt"
	"html"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
	"golang.org/x/crypto/ssh/terminal"
)

// getHighlightID ハイライトされているIDを取得
func getHighlightID(p *tview.TextView) int {
	hl := p.GetHighlights()
	if hl == nil {
		return -1
	}
	reg := regexp.MustCompile(".+_(\\d+)")
	find := reg.FindAllStringSubmatch(hl[0], -1)
	idx, _ := strconv.Atoi(find[0][1])
	return idx
}

// getWindowWidth 表示領域の幅を取得
func getWindowWidth() int {
	fd := int(os.Stdout.Fd())
	w, _, err := terminal.GetSize(fd)
	if err != nil {
		log.Fatal(err)
	}
	return w - 4
}

// IsSameDate 今日の日付かどうか
func isSameDate(a time.Time) bool {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	t2 := time.Date(a.Year(), a.Month(), a.Day(), 0, 0, 0, 0, t.Location())
	return t1.Equal(t2)
}

// createUserString ユーザー情報文字列
func createUserString(u *anaconda.User) string {
	// ユーザー名
	username := removeRuneAmbiguousWidth(u.Name)
	// ユーザータイプ
	badge := ""
	if u.Verified {
		badge += fmt.Sprintf("[#%x] verified", sc.Cfg.Color.Verified)
	}
	if u.Protected {
		badge += fmt.Sprintf("[#%x] protected", sc.Cfg.Color.Protected)
	}
	// 結合
	text := fmt.Sprintf("[#%x]%s [#%x]@%s%s", sc.Cfg.Color.Username, username, sc.Cfg.Color.Screenname, u.ScreenName, badge)
	return text
}

// createTimeString 投稿時刻
func createTimeString(t time.Time) string {
	postTime := ""
	if isSameDate(t) {
		postTime = t.Local().Format(sc.Cfg.Option.TimeFormat)
	} else {
		format := fmt.Sprintf("%s %s", sc.Cfg.Option.DateFormat, sc.Cfg.Option.TimeFormat)
		postTime = t.Local().Format(format)
	}
	return fmt.Sprintf("[#%x]%s", sc.Cfg.Color.Accent1, postTime)
}

// createReaction リアクション数文字列
func createReaction(text *string, count int, flg bool, unit string, hex int32) {
	// カウント数チェック
	if count <= 0 {
		return
	} else if count > 1 {
		unit += "s"
	}
	// 文字列作成
	*text += " "
	if flg {
		*text += fmt.Sprintf("[#%x:-:r] %d%s [-:-:-]", hex, count, unit)
	} else {
		*text += fmt.Sprintf("[#%x]%d%s", hex, count, unit)
	}
}

// removeRuneAmbiguousWidth 曖昧な幅の文字を除去する
func removeRuneAmbiguousWidth(text string) string {
	result := ""
	text = html.UnescapeString(text)
	for _, r := range text {
		if runewidth.IsAmbiguousWidth(r) {
			continue
		}
		result += string(r)
	}
	return result
}
