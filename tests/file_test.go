package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/siriusol/pic-markdown/utils"
)

func TestCheckFileExist(t *testing.T) {
	t.Log(utils.CheckFileExist("E:/测试/Redis in action.md"))
	t.Log(utils.CheckFileExist("E:/测试/Redis in action-1.md"))
	t.Log(utils.CheckFileExist("E:/测试/test.txt"))
	t.Log(utils.CheckFileExist("https://baidu.com"))
}

func TestReadDir(t *testing.T) {
	files1, _ := utils.ReadDirFilesExcludeSubDir("E:/test")
	for _, file := range files1 {
		if !file.IsDir() {
			t.Log(file.Name(), file.Size(), file.Mode(), file.ModTime(), file.IsDir(), file.Sys())
		}
	}
}

func TestFilePath(t *testing.T) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dir)
	t.Log(os.Args[0])
}

func TestReplaceFromOneLine(t *testing.T) {
	utils.ReplaceOneLine("E:/test/Redis in action-1.md")
}
