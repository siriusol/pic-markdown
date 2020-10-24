package tests

import (
	"testing"

	"github.com/siriusol/pic-markdown/utils"
)

func TestGenerateUploadToken(t *testing.T) {
	// 每次输出都不同
	t.Log(utils.GenerateUploadToken())
}

func TestUploadFile(t *testing.T) {
	cloudUrl := utils.UploadFile("E:\\测试图片\\Redis比较.jpg")
	// cloudUrl := utils.UploadFile("C:\\Users\\Ther\\AppData\\Roaming\\Typora\\typora-user-images\\image-20200301223735099.png")
	t.Log(cloudUrl)
}
