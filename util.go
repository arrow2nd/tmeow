package main

import (
	"regexp"
	"strconv"

	"github.com/rivo/tview"
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
