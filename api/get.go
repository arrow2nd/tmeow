package api

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

// GetFriendships ユーザーとの関係を取得
func (tw *TwitterAPI) GetFriendships(u *anaconda.User) (string, error) {
	v := url.Values{"user_id": {u.IdStr}}
	friendships, err := tw.API.GetFriendshipsLookup(v)
	if err != nil {
		return "", errors.New(parseAPIError(err))
	}

	followedBy, following, blocking, muting := false, false, false, false
	for _, v := range friendships[0].Connections {
		switch v {
		case "followed_by":
			followedBy = true
		case "following":
			following = true
		case "blocking":
			blocking = true
		case "muting":
			muting = true
		}
	}

	status := ""
	if followedBy && following {
		status += fmt.Sprintf("[blue]相互フォロー[white] ")
	} else if followedBy {
		status += fmt.Sprintf("[blue]フォローされています[while] ")
	} else if following {
		status += fmt.Sprintf("[blue]フォロー中[while] ")
	}
	if blocking {
		status += fmt.Sprintf("[red]ブロック中[white] ")
	}
	if muting {
		status += fmt.Sprintf(" [yellow]ミュート中[white]")
	}

	return status, nil
}

// GetTimeline タイムラインを取得
func (tw *TwitterAPI) GetTimeline(mode string, v url.Values) (*[]anaconda.Tweet, error) {
	var (
		timeline []anaconda.Tweet
		err      error
	)

	switch mode {
	case "home":
		timeline, err = tw.API.GetHomeTimeline(v)
	case "mention":
		timeline, err = tw.API.GetMentionsTimeline(v)
	case "user":
		timeline, err = tw.API.GetUserTimeline(v)
	}

	if err != nil {
		return nil, errors.New(parseAPIError(err))
	}

	return &timeline, nil
}

// GetListTimeline リストタイムラインを取得
func (tw *TwitterAPI) GetListTimeline(listID int64, count string) (*[]anaconda.Tweet, error) {
	v := createURLValues(count)

	timeline, err := tw.API.GetListTweets(listID, true, v)
	if err != nil {
		return nil, errors.New(parseAPIError(err))
	}

	return &timeline, nil
}

// GetSearchResult 検索結果を取得
func (tw *TwitterAPI) GetSearchResult(query, count string) (*[]anaconda.Tweet, error) {
	v := createURLValues(count)
	query += " -filter:retweets"

	result, err := tw.API.GetSearch(query, v)
	if err != nil {
		return nil, errors.New(parseAPIError(err))
	}

	return &result.Statuses, nil
}

// getSelf 自分のユーザー情報を取得
func (tw *TwitterAPI) getSelf() (*anaconda.User, error) {
	user, err := tw.API.GetSelf(nil)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// getLists リストの一覧を取得
func (tw *TwitterAPI) getLists() ([]string, []int64, error) {
	lists, err := tw.API.GetListsOwnedBy(tw.OwnUser.Id, nil)
	if err != nil {
		return nil, nil, err
	}

	// リスト名とIDのスライスを作成
	id := make([]int64, len(lists))
	name := make([]string, len(lists))
	for i, l := range lists {
		name[i] = l.Name
		id[i] = l.Id
	}

	return name, id, nil
}
