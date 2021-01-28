package api

import (
	"github.com/ChimeraCoder/anaconda"
)

const (
	consumerKey    = "MTSe5vV5KjtyCKAEKwgvBuxUV"
	consumerSecret = "GRUBtJhmdSqOpcP4pKN9ZDUMwgxU22NyB5CvU7T63TViUvCFdG"
)

// TwitterAPI API構造体
type TwitterAPI struct {
	API       *anaconda.TwitterApi
	OwnUser   *anaconda.User
	ListNames []string
	ListIDs   []int64
}

func init() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
}

// NewTwitterAPI TwitterAPI構造体を作成
func NewTwitterAPI(token, secret string) *TwitterAPI {
	tw := &TwitterAPI{
		API:       anaconda.NewTwitterApi(token, secret),
		OwnUser:   &anaconda.User{},
		ListNames: []string{},
		ListIDs:   []int64{},
	}
	return tw
}

// Init 初期化
func (tw *TwitterAPI) Init() error {
	var err error

	// ユーザー情報を取得
	tw.OwnUser, err = tw.getSelf()
	if err != nil {
		return err
	}

	// リスト情報を取得
	tw.ListNames, tw.ListIDs, err = tw.getLists()
	if err != nil {
		return err
	}

	return nil
}
