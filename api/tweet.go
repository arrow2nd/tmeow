package api

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"path"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

// PostTweet ツイートを投稿
func (tw *TwitterAPI) PostTweet(status, replyToID string, files []string) error {
	// 画像をアップロード
	val, err := tw.uploadImage(files)
	if err != nil {
		return err
	}

	// リプライ先設定
	if replyToID != "" {
		val.Add("in_reply_to_status_id", replyToID)
		val.Add("auto_populate_reply_metadata", "true")
	}

	// ツイート
	_, err = tw.API.PostTweet(status, val)
	if err != nil {
		return errors.New(parseAPIError(err))
	}

	return nil
}

// DeleteTweet ツイートを削除
func (tw *TwitterAPI) DeleteTweet(tweetID string) error {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	_, err := tw.API.DeleteTweet(id, true)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}

// uploadImage 画像をアップロード
func (tw *TwitterAPI) uploadImage(files []string) (url.Values, error) {
	var (
		val      = url.Values{}
		fileNum  = len(files)
		mediaIds = make([]string, fileNum)
	)

	// ファイル数チェック
	if fileNum == 0 {
		return val, nil
	} else if fileNum > 4 {
		return nil, errors.New("添付できる画像は4枚までです")
	}

	eg, ctx := errgroup.WithContext(context.Background())

	for idx, filename := range files {
		// 未対応の拡張子ならスキップ
		if ext := strings.ToLower(path.Ext(filename)); ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			return nil, fmt.Errorf("サポートされていない拡張子です(%s)", ext)
		}

		// アップロード処理
		idx, filename := idx, filename
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				// ファイルを読み込む
				data, err := ioutil.ReadFile(filename)
				if err != nil {
					return fmt.Errorf("ファイルの読み込みに失敗しました(%s)", filename)
				}

				// base64にエンコードしてアップロード
				enc := base64.StdEncoding.EncodeToString(data)
				media, err := tw.API.UploadMedia(enc)
				if err != nil {
					return fmt.Errorf("アップロードに失敗しました(%s)", filename)
				}

				mediaIds[idx] = media.MediaIDString
				return nil
			}
		})
	}

	// 待機
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// カンマで連結
	val.Add("media_ids", strings.Join(mediaIds, ","))
	return val, nil
}
