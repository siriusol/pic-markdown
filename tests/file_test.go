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
	files, _ := utils.ReadDirFiles("E:/test")
	for _, file := range files {
		if !file.IsDir() {
			t.Log("File:", file.Name(), file.Size(), file.Mode(), file.ModTime(), file.IsDir(), file.Sys())
		} else {
			t.Log("Dir:", file.Name(), file.Size(), file.Mode(), file.ModTime(), file.IsDir(), file.Sys())
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

func TestGetAllUrlsFromFile(t *testing.T) {
	urls := utils.GetAllUrlsFromFile("E:/test/Redis in action-1.md")
	t.Log(len(urls))
	t.Log(urls)
}
