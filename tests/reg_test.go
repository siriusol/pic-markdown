package tests

import (
	"testing"

	"github.com/siriusol/pic-markdown/utils"
)

func TestReg(t *testing.T) {
	pic := "This is a picture ![百度](https://baidu.com). And here is another. ![腾讯QQ](https://qq.com)"
	urls := utils.GetUrlFromOneLine(pic)
	for _, value := range urls {
		t.Log(value)
	}
}

func TestReplaceOneLine(t *testing.T) {
	text := "This is a picture ![百度](C:\\Users\\Ther\\AppData\\Roaming\\Typora\\typora-user-images\\image-20200301223735099.png). And here is another. ![腾讯QQ](https://qq.com)"
	reText := utils.ReplaceUrlFromOneLine(text, []string{"First", "Second"})
	t.Log(reText)
}
