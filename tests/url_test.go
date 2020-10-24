package tests

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	decodeUrl := "qip9mf1yt.hb-bkt.clouddn.com/Redis比较.jpg?time=2020-10-24#17:34:57"
	fmt.Println(decodeUrl)
	encodeUrl := url.QueryEscape(decodeUrl)
	fmt.Println(encodeUrl)

}
