package api

import (
	"fmt"
	"log"

	"github.com/ChimeraCoder/anaconda"
)

const logo = `
888                                                  
888                                                  
888                                                  
888888 88888b.d88b.   .d88b.   .d88b.  888  888  888 
888    888 "888 "88b d8P  Y8b d88""88b 888  888  888 
888    888  888  888 88888888 888  888 888  888  888 
Y88b.  888  888  888 Y8b.     Y88..88P Y88b 888 d88P 
 "Y888 888  888  888  "Y8888   "Y88P"   "Y8888888P"  
`

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

// NewTwitterAPI API構造体を作成
func NewTwitterAPI() *TwitterAPI {
	tw := &TwitterAPI{
		API:       nil,
		OwnUser:   &anaconda.User{},
		ListNames: []string{},
		ListIDs:   []int64{},
	}
	return tw
}

// Init 初期化
func (tw *TwitterAPI) Init(token, secret string) error {
	var err error

	tw.API = anaconda.NewTwitterApi(token, secret)

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

// Auth 認証
func (tw *TwitterAPI) Auth() (string, string) {

	authAPI := anaconda.NewTwitterApi("", "")
	uri, cred, err := authAPI.AuthorizationURL("oob")
	if err != nil {
		fmt.Println("認証URLの発行に失敗しました")
		log.Fatal(err)
	}

	fmt.Printf("%s\n\n", logo)
	fmt.Println("🐈 以下のURLにアクセスしてアプリケーションを認証し、表示されるPINを入力してください。")
	fmt.Printf("[ %s ]\n\n", uri)

	// PIN入力
	pin := ""
	fmt.Print("PIN : ")
	fmt.Scanf("%s", &pin)

	// トークン発行
	cred, _, err = authAPI.GetCredentials(cred, pin)
	if err != nil {
		fmt.Println("アクセストークンが取得できませんでした")
		log.Fatal(err)
	}

	tw.Init(cred.Token, cred.Secret)

	return cred.Token, cred.Secret
}
