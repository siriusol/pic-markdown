package tests

import (
	"os"
	"testing"

	"github.com/siriusol/pic-markdown/utils"
)

func TestQiNiuConfig(t *testing.T) {
	nowPath, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(nowPath)
	configPath := nowPath + "\\..\\conf\\config.txt"
	config := utils.GetQiNiuConfig(configPath)
	t.Logf("%+v", config)
}

func TestGenerateUploadToken(t *testing.T) {
	nowPath, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(nowPath)
	configPath := nowPath + "\\..\\conf\\config.txt"
	// 每次输出都不同
	config := utils.GetQiNiuConfig(configPath)
	t.Log(utils.GenerateUploadToken(config))
}

func TestUploadFile(t *testing.T) {
	filePath := "E:\\test\\Redis in action-1.md"
	nowPath, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(nowPath)
	configPath := nowPath + "\\..\\conf\\config.txt"
	config := utils.GetQiNiuConfig(configPath)
	cloudUrl, err := utils.UploadFile(filePath, config)
	if err != nil {
		t.Log(err)
	}
	t.Log(cloudUrl)
}
