package tests

import (
	"github.com/siriusol/pic-markdown/utils"
	"testing"
)

func TestReg(t *testing.T) {
	pic := "This is a picture ![百度](https://baidu.com). And here is another. ![腾讯QQ](https://qq.com)"
	urls := utils.GetUrlFromOneLine(pic)
	for _, value := range urls {
		t.Log(value)
	}
}
