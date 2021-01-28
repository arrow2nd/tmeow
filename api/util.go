package api

import (
	"fmt"
	"net/url"
	"regexp"
)

func parseAPIError(err error) string {
	bytes := []byte(err.Error())
	errMsg := regexp.MustCompile("\"(message|error)\":\\s*\"(.+)\"").FindSubmatch(bytes)
	return fmt.Sprint(errMsg[2])
}

func createURLValues(cnt string) url.Values {
	v := url.Values{}
	v.Add("tweet_mode", "extended")
	v.Add("count", cnt)
	return v
}
