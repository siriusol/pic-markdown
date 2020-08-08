package main

import (
	"github.com/siriusol/pic-markdown/utils"
	"log"
)

func main() {
	// 输入一个 markdown 目录, 一个 picture 目录
	textDirPath := `E:\测试`
	picDirPath := `E:\测试图片`
	execute(textDirPath, picDirPath)
}

func execute(textDirPath, picDirPath string) {
	texts, textErr := utils.ReadDirFilesExcludeSubDir(textDirPath)
	if textErr != nil {
		log.Fatal(textErr)
	}
	_, picErr := utils.ReadDirFilesExcludeSubDir(picDirPath)
	if picErr != nil {
		log.Fatal(picErr)
		return
	}
	for _, text := range texts {
		if !text.IsDir() {
			// TODO  Replace separator.
			utils.ReplaceOneLine(textDirPath+"\\"+text.Name(), picDirPath)
		}
	}
}